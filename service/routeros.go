package service

import (
	"crypto/tls"
	"time"

	routeros "github.com/go-routeros/routeros/v3"
)

type RouterOSService struct {
	addr     string
	username string
	password string

	timeout time.Duration
	client  *routeros.Client
}

type Option func(*RouterOSService)

func NewRouterOSService(opts ...Option) *RouterOSService {
	svc := &RouterOSService{
		timeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *RouterOSService) {
		s.timeout = timeout
	}
}

func (s *RouterOSService) Connect(address, username, password string, tlsConfig *tls.Config) error {
	s.addr = address
	s.username = username
	s.password = password

	var (
		client *routeros.Client
		err    error
	)

	switch {
	case tlsConfig != nil && s.timeout > 0:
		client, err = routeros.DialTLSTimeout(address, username, password, tlsConfig, s.timeout)
	case tlsConfig != nil:
		client, err = routeros.DialTLS(address, username, password, tlsConfig)
	case s.timeout > 0:
		client, err = routeros.DialTimeout(address, username, password, s.timeout)
	default:
		client, err = routeros.Dial(address, username, password)
	}

	if err != nil {
		return err
	}

	s.client = client
	return nil
}

func (s *RouterOSService) Close() error {
	if s.client == nil {
		return nil
	}

	err := s.client.Close()
	s.client = nil
	return err
}

func (s *RouterOSService) Ping() error {
	client, err := s.ensureClient()
	if err != nil {
		return err
	}

	_, err = client.Run("/system/identity/print")
	return err
}

func (s *RouterOSService) Run(cmd string, args ...string) (map[string]string, error) {
	client, err := s.ensureClient()
	if err != nil {
		return nil, err
	}

	sentences := append([]string{cmd}, args...)
	reply, err := client.Run(sentences...)
	if err != nil {
		return nil, err
	}

	if reply.Done != nil && len(reply.Done.Map) > 0 {
		return reply.Done.Map, nil
	}

	if len(reply.Re) > 0 {
		return reply.Re[0].Map, nil
	}

	return map[string]string{}, nil
}

func (s *RouterOSService) ensureClient() (*routeros.Client, error) {
	if s.client == nil {
		return nil, ErrNotConnected
	}

	return s.client, nil
}
