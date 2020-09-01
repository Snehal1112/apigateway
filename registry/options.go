package registry

import (
	"context"
	"github.com/snehal1112/gateway/transport"
)

type OptionsMgr func(mgr *Manager)

func WithCtx(ctx context.Context) OptionsMgr {
	return func(t *Manager) {
		t.ctx = ctx
	}
}

func WithDB(ctx context.Context, url string) OptionsMgr {
	return func(mgr *Manager) {
		mgr.db = transport.NewDBLayer(transport.WithDBConnect(url, ctx))
	}
}
