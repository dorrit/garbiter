package garbiter

import (
	"context"
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

// WithCommandTimeout configures the timeout applied to commands that do not
// receive an explicit context. A non-positive value disables the default.
func WithCommandTimeout(timeout time.Duration) Option {
	return service.WithCommandTimeout(timeout)
}

// New creates a client with the configured transport options without connecting.
func New(opts ...service.Option) *Client {
	return &Client{service: service.NewRouterOSService(opts...)}
}

// Connect dials a RouterOS API endpoint and returns a typed client.
func Connect(addr, user, pass string, opts ...service.Option) (*Client, error) {
	return ConnectContext(context.Background(), addr, user, pass, opts...)
}

// ConnectContext dials a RouterOS API endpoint with context cancellation.
func ConnectContext(ctx context.Context, addr, user, pass string, opts ...service.Option) (*Client, error) {
	c := New(opts...)
	if err := c.service.ConnectContext(ctx, addr, user, pass); err != nil {
		return nil, err
	}
	return c, nil
}

// ConnectTLS dials a RouterOS API endpoint over TLS and returns a typed client.
func ConnectTLS(addr, user, pass string, tlsCfg *tls.Config, opts ...service.Option) (*Client, error) {
	return ConnectTLSContext(context.Background(), addr, user, pass, tlsCfg, opts...)
}

// ConnectTLSContext dials a RouterOS TLS endpoint with context cancellation.
func ConnectTLSContext(ctx context.Context, addr, user, pass string, tlsCfg *tls.Config, opts ...service.Option) (*Client, error) {
	if tlsCfg == nil {
		return nil, service.ErrInvalidTLSConfig
	}
	c := New(opts...)
	if err := c.service.ConnectTLSContext(ctx, addr, user, pass, tlsCfg); err != nil {
		return nil, err
	}
	return c, nil
}
