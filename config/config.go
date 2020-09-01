package config

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type Config struct {
	ListenAddr string

	Logger        logrus.FieldLogger
	HTTPTransport http.RoundTripper

	TrustedProxyIPs  []*net.IP
	TrustedProxyNets []*net.IPNet
}
