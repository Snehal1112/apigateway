package server

import "github.com/sirupsen/logrus"

type Options func(s *Server)

func WithConfig(cfg *Config) Options {
	return func(s *Server) {
		s.Config = cfg
	}
}

func WithLogger(logger logrus.FieldLogger) Options {
	return func(s *Server) {
		s.logger = logger
	}
}

func WithListener(listen string) Options {
	return func(s *Server) {
		s.listenAddr = listen
	}
}

//func WithBasePath(basePath string) Options {
//	return func(s *Server) {
//		s.listenAddr = listen
//	}
//}
