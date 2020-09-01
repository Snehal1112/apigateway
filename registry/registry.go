package registry

import (
	"github.com/snehal1112/gateway/proxy"
)

type Service struct {
	Name        string       `bson:"name" json:"name"`
	Active      bool         `bson:"active" json:"active"`
	Proxy       *proxy.Proxy `bson:"proxy" json:"proxy"`
	HealthCheck HealthCheck  `bson:"health_check" json:"health_check"`
}

// About this health check we can think about later
type HealthCheck struct {
	URL     string `bson:"url" json:"url"`
	Timeout int    `bson:"timeout" json:"timeout"`
}

func NewService(name string, active bool, proxy *proxy.Proxy) *Service {
	return &Service{Name: name, Active: active, Proxy: proxy}
}
