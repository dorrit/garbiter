package service

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	routeros "github.com/go-routeros/routeros/v3"
)

type RouterOSService struct {
	mu             sync.Mutex
	timeout        time.Duration
	commandTimeout time.Duration
	client         routerOSClient
}

type routerOSClient interface {
	RunContext(ctx context.Context, sentences ...string) (*routeros.Reply, error)
	Close() error
}

type Option func(*RouterOSService)

func NewRouterOSService(opts ...Option) *RouterOSService {
	svc := &RouterOSService{
		timeout:        5 * time.Second,
		commandTimeout: 30 * time.Second,
	}

	for _, opt := range opts {
		if opt != nil {
			opt(svc)
		}
	}

	return svc
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *RouterOSService) {
		s.timeout = timeout
	}
}

func WithCommandTimeout(timeout time.Duration) Option {
	return func(s *RouterOSService) {
		s.commandTimeout = timeout
	}
}

func (s *RouterOSService) Connect(address, username, password string) error {
	return s.ConnectContext(context.Background(), address, username, password)
}

func (s *RouterOSService) ConnectContext(ctx context.Context, address, username, password string) error {
	ctx, cancel := s.withDialTimeout(ctx)
	defer cancel()

	client, err := routeros.DialContext(ctx, address, username, password)
	if err != nil {
		return err
	}

	return s.replaceClient(client)
}

func (s *RouterOSService) ConnectTLS(address, username, password string, tlsConfig *tls.Config) error {
	return s.ConnectTLSContext(context.Background(), address, username, password, tlsConfig)
}

func (s *RouterOSService) ConnectTLSContext(ctx context.Context, address, username, password string, tlsConfig *tls.Config) error {
	if tlsConfig == nil {
		return ErrInvalidTLSConfig
	}

	ctx, cancel := s.withDialTimeout(ctx)
	defer cancel()

	client, err := routeros.DialTLSContext(ctx, address, username, password, tlsConfig)
	if err != nil {
		return err
	}

	return s.replaceClient(client)
}

func (s *RouterOSService) replaceClient(client routerOSClient) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.client != nil {
		if err := s.client.Close(); err != nil {
			client.Close()
			return err
		}
	}
	s.client = client
	return nil
}

func (s *RouterOSService) withDialTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if s.timeout <= 0 {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, s.timeout)
}

func (s *RouterOSService) withCommandTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if s.commandTimeout <= 0 {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, s.commandTimeout)
}

func (s *RouterOSService) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.client == nil {
		return nil
	}

	err := s.client.Close()
	s.client = nil
	return err
}

func (s *RouterOSService) Ping() error {
	_, err := s.Run("/system/identity/print")
	return err
}

func (s *RouterOSService) Run(cmd string, args ...string) (map[string]string, error) {
	ctx, cancel := s.withCommandTimeout(context.Background())
	defer cancel()
	return s.RunContext(ctx, cmd, args...)
}

func (s *RouterOSService) RunContext(ctx context.Context, cmd string, args ...string) (map[string]string, error) {
	rows, done, err := s.runContext(ctx, cmd, args...)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	}

	if len(done) > 0 {
		return done, nil
	}

	return map[string]string{}, nil
}

func (s *RouterOSService) RunList(cmd string, args ...string) ([]map[string]string, error) {
	ctx, cancel := s.withCommandTimeout(context.Background())
	defer cancel()
	return s.RunListContext(ctx, cmd, args...)
}

func (s *RouterOSService) RunListContext(ctx context.Context, cmd string, args ...string) ([]map[string]string, error) {
	rows, _, err := s.runContext(ctx, cmd, args...)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows, nil
	}

	return []map[string]string{}, nil
}

func (s *RouterOSService) runContext(ctx context.Context, cmd string, args ...string) ([]map[string]string, map[string]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.client == nil {
		return nil, nil, ErrNotConnected
	}

	sentences := append([]string{cmd}, args...)
	reply, err := s.client.RunContext(ctx, sentences...)
	if err != nil {
		return nil, nil, err
	}

	rows := make([]map[string]string, 0, len(reply.Re))
	for _, sentence := range reply.Re {
		if len(sentence.Map) > 0 {
			rows = append(rows, sentence.Map)
		}
	}

	done := map[string]string(nil)
	if reply.Done != nil && len(reply.Done.Map) > 0 {
		done = reply.Done.Map
	}

	return rows, done, nil
}
