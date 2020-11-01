package transport

import (
	"context"
	"github.com/snehal1112/transport/client"
)

type DBLayer struct {
	ctx     context.Context
	Connect *client.Connect
}

func NewDBLayer(options ...OptionsTransport) *DBLayer {
	dbLayer := &DBLayer{}
	for _, option := range options {
		option(dbLayer)
	}
	return dbLayer
}

func (m DBLayer) Load(ctx context.Context) error {

	return nil
}
