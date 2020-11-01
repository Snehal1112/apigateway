package server

import (
	"github.com/snehal1112/gateway/config"
	"github.com/snehal1112/gateway/registry"
)

// Config defines a Server's configuration settings.
type Config struct {
	Config   *config.Config
	Services []*registry.Service
	BasePath string
}

func (c *Config) getServices() []*registry.Service {
	return c.Services
}
