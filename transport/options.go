package transport

import (
	"context"
	"github.com/snehal1112/transport/client"
)

type OptionsTransport func(db *DBLayer)

func WithDBConnect(url string, ctx context.Context) OptionsTransport {
	return func(t *DBLayer) {
		t.Connect = client.NewConnection(
			client.WithCtx(ctx),
			client.WithURL(url),
			client.WithDatabase("gateway"),//TODO: make it configurable
			client.WithLogLevel("error"),
		)
	}
}
