package garbiter

import (
	"crypto/tls"
	"time"

	"github.com/dorrit/garbiter/service"
)

// Option configures the underlying transport.
type Option = service.Option

// WithTimeout configures the dial timeout. A non-positive timeout disables the
// timeout-specific dial path and uses the dependency default behavior.
func WithTimeout(timeout time.Duration) Option {
	return service.WithTimeout(timeout)
}

// New creates a client with the configured transport options without connecting.
func New(opts ...service.Option) *Client {
	return &Client{service: service.NewRouterOSService(opts...)}
}

// Connect dials a RouterOS API endpoint and returns a typed client.
func Connect(addr, user, pass string, opts ...service.Option) (*Client, error) {
	c := New(opts...)
	err := c.service.Connect(addr, user, pass, nil)
	return c, err
}

// ConnectTLS dials a RouterOS API endpoint over TLS and returns a typed client.
func ConnectTLS(addr, user, pass string, tlsCfg *tls.Config, opts ...service.Option) (*Client, error) {
	c := New(opts...)
	err := c.service.Connect(addr, user, pass, tlsCfg)
	return c, err
}
