package model

import (
	"context"
	"crypto/tls"
)

// Transport describes the minimal RouterOS transport implementation the client needs.
// It intentionally mirrors the go-routeros client surface we rely on, making it easy to mock.
type Transport interface {
	Connect(address, username, password string) error
	ConnectContext(ctx context.Context, address, username, password string) error
	ConnectTLS(address, username, password string, tlsConfig *tls.Config) error
	ConnectTLSContext(ctx context.Context, address, username, password string, tlsConfig *tls.Config) error
	Close() error
	Ping() error
	Run(cmd string, args ...string) (map[string]string, error)
	RunContext(ctx context.Context, cmd string, args ...string) (map[string]string, error)
	RunList(cmd string, args ...string) ([]map[string]string, error)
	RunListContext(ctx context.Context, cmd string, args ...string) ([]map[string]string, error)
}
