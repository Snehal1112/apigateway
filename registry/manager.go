package registry

import (
	"context"
	"github.com/snehal1112/gateway/transport"
	"go.mongodb.org/mongo-driver/bson"
)

type ManagerInterface interface {
	Load(ctx context.Context) error
	RegisterService(service *Service) error
	CreateNewService(name string, active bool) *Service
	Services() []*Service
}

type Manager struct {
	ctx      context.Context
	services []*Service
	db       *transport.DBLayer
}

func (m *Manager) Services() []*Service {
	return m.services
}

func (m *Manager) CreateNewService(name string, active bool) *Service {
	return NewService(name, active)
}

func (m *Manager) RegisterService(service *Service) error {
	m.services = append(m.services, service)
	return nil
}

func (m *Manager) Load(ctx context.Context) error {
	services, err := m.db.Connect.Search("services", bson.D{}, 0, 0)
	if err != nil {
		ctx.Deadline()
		return err
	}

	for _, b := range services {
		data, err := bson.Marshal(b)
		if err != nil {
			return err
		}
		service := &Service{}
		bson.Unmarshal(data, service)
		m.services = append(m.services, service)
	}
	return nil
}

func NewManager(options ...OptionsMgr) ManagerInterface {
	mgr := &Manager{}
	for _, option := range options {
		option(mgr)
	}
	return mgr
}

//https://github.com/olivere/balancers
