package garbiter

import (
	"crypto/tls"

	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

func New(opts ...service.Option) *model.Garbiter {
	return &model.Garbiter{
		Service: service.NewRouterOSService(opts...),
	}
}

func Connect(addr, user, pass string) (*model.Garbiter, error) {
	g := New()
	err := g.Service.Connect(addr, user, pass, nil)
	return g, err
}

func ConnectTLS(addr, user, pass string, tlsCfg *tls.Config) (*model.Garbiter, error) {
	g := New()
	err := g.Service.Connect(addr, user, pass, tlsCfg)
	return g, err
}
