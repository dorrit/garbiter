package model

import (
	"crypto/tls"
)

type Garbiter struct {
	Service GarbiterService
}

type GarbiterService interface {
	Connect(address, username, password string, tlsConfig *tls.Config) error
	Close() error
	Ping() error
	Run(cmd string, args ...string) (map[string]string, error)
}
