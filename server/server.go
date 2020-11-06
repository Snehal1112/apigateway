package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/snehal1112/gateway/registry"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Config *Config

	listenAddr string
	logger     logrus.FieldLogger
	requestLog bool
}

func NewServer(options ...Options) (*Server, error) {
	ser := &Server{}
	for _, option := range options {
		option(ser)
	}
	return ser, nil
}

func (s *Server) AddContext(parent context.Context, next *httputil.ReverseProxy, service *registry.Service) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithCancel(parent)
		dd, _ := json.MarshalIndent(service, "", "    ")
		logrus.Println("service:-", string(dd))
		log.Println("params:-", req.URL.Query().Get("data"))
		next.ServeHTTP(rw, req.WithContext(ctx))

		cancel()
	})
}

func (s *Server) registerService(serverCtx context.Context, router *mux.Router) {
	services := s.Config.getServices()
	for _, service := range services {
		proxy := service.Proxy
		upstreams := proxy.Upstreams

		for _, target := range upstreams.Targets {
			url, err := url.Parse(target.Target)
			if err != nil {
				s.logger.WithFields(logrus.Fields{
					"service_name":   service.Name,
					"service_active": service.Active,
					"target":         url,
				}).Infoln("Unable to pars the service url.")
				continue
			}

			s.logger.WithFields(logrus.Fields{
				"service_name":   service.Name,
				"service_active": service.Active,
				"target":         url,
			}).Infoln("Service registered.")

			log.Println(s.Config.BasePath + proxy.ListenPath)
			if proxy.StripPath {
				router.Handle(
					s.Config.BasePath + proxy.ListenPath,
					http.StripPrefix(
						s.Config.BasePath,
						s.AddContext(
							serverCtx,
							httputil.NewSingleHostReverseProxy(url),
							service,
						),
					),
				)
				//router.Handle(
				//	s.Config.BasePath + proxy.ListenPath,
				//	http.StripPrefix(
				//		s.Config.BasePath,
				//		s.AddContext(
				//			serverCtx,
				//			httputil.NewSingleHostReverseProxy(url),
				//			service,
				//		),
				//	),
				//)
			}
		}
	}
}

func (s *Server) Serve(ctx context.Context) error {
	serverCtx, serveCtxCancel := context.WithCancel(ctx)
	defer serveCtxCancel()

	logger := s.logger
	errCh := make(chan error, 2)
	exitCh := make(chan bool, 1)
	signalCh := make(chan os.Signal)

	mux := mux.NewRouter()
	s.registerService(serverCtx, mux)

	// HTTP listener.
	srv := &http.Server{
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.WithField("listenAddr", s.listenAddr).Infoln("starting http listener")

	// TODO: Also support unix socket.
	listener, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	logger.Infoln("ready to handle requests")

	go func() {
		serveErr := srv.Serve(listener)
		if serveErr != nil {
			errCh <- serveErr
		}

		logger.Debugln("http listener stopped")
		close(exitCh)
	}()

	// Wait for exit or error.
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err = <-errCh:
		// breaks
	case reason := <-signalCh:
		logger.WithField("signal", reason).Warnln("received signal")
		// breaks
	}

	// Shutdown, server will stop to accept new connections, requires Go 1.8+.
	logger.Infoln("clean server shutdown start")

	shutDownCtx, shutDownCtxCancel := context.WithTimeout(ctx, 10*time.Second)
	if shutdownErr := srv.Shutdown(shutDownCtx); shutdownErr != nil {
		logger.WithError(shutdownErr).Warn("clean server shutdown failed")
	}

	// Cancel our own context, wait on managers.
	serveCtxCancel()
	func() {
		for {
			select {
			case <-exitCh:
				return
			default:
				// HTTP listener has not quit yet.
				logger.Info("waiting for http listener to exit")
			}
			select {
			case reason := <-signalCh:
				logger.WithField("signal", reason).Warn("received signal")
				return
			case <-time.After(100 * time.Millisecond):
			}
		}
	}()
	shutDownCtxCancel() // prevent leak.

	return err
}
