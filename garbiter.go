package garbiter

import (
	"crypto/tls"
	"time"

	"github.com/dorrit/garbiter/service"
)

// Option configures the underlying transport.
type Option = service.Option

// WithTimeout wraps service.WithTimeout for convenience.
func WithTimeout(timeout time.Duration) Option {
	return service.WithTimeout(timeout)
}

func New(opts ...service.Option) *Client {
	return &Client{service: service.NewRouterOSService(opts...)}
}

func Connect(addr, user, pass string, opts ...service.Option) (*Client, error) {
	c := New(opts...)
	err := c.service.Connect(addr, user, pass, nil)
	return c, err
}

func ConnectTLS(addr, user, pass string, tlsCfg *tls.Config, opts ...service.Option) (*Client, error) {
	c := New(opts...)
	err := c.service.Connect(addr, user, pass, tlsCfg)
	return c, err
}
