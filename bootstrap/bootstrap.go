package bootstrap

import (
	"context"
	"github.com/snehal1112/gateway/config"
	"github.com/snehal1112/gateway/registry"
)

type Bootstrap interface {
	Config() *config.Config
	Manager() registry.ManagerInterface
}

type Config struct {
	URIBasePath string
	Listen      string
	BackendURL  string
}

type bootstrap struct {
	uriBasePath string
	backendURI  string

	cfg             *config.Config
	RegistryManager registry.ManagerInterface
}

func (b *bootstrap) Config() *config.Config {
	return b.cfg
}

func (b *bootstrap) Manager() registry.ManagerInterface {
	return b.RegistryManager
}

func Boot(ctx context.Context, cfg *Config, serverConf *config.Config) (Bootstrap, error) {
	bs := &bootstrap{
		cfg: serverConf,
	}

	if err := bs.setup(ctx, cfg); err != nil {
		return nil, err
	}

	return bs, nil
}

func (b *bootstrap) setup(ctx context.Context, cfg *Config) error {
	b.cfg.ListenAddr = cfg.Listen
	b.uriBasePath = cfg.URIBasePath
	b.backendURI = cfg.BackendURL

	b.RegistryManager = registry.NewManager(
		registry.WithCtx(ctx),
		registry.WithDB(ctx, b.backendURI),
	)

	if err := b.RegistryManager.Load(ctx); err != nil {
		return err
	}

	return nil
}