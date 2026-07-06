package garbiter

import (
	"crypto/tls"
	"errors"
	"testing"

	"github.com/dorrit/garbiter/service"
)

type fakeTransport struct {
	connectErr error
	closeErr   error
	pingErr    error
	runErr     error
	runResult  map[string]string

	connectCalled bool
	closeCalled   bool
	pingCalled    bool
	runCmd        string
	runArgs       []string
}

func (f *fakeTransport) Connect(address, username, password string, tlsConfig *tls.Config) error {
	f.connectCalled = true
	return f.connectErr
}

func (f *fakeTransport) Close() error {
	f.closeCalled = true
	return f.closeErr
}

func (f *fakeTransport) Ping() error {
	f.pingCalled = true
	return f.pingErr
}

func (f *fakeTransport) Run(cmd string, args ...string) (map[string]string, error) {
	f.runCmd = cmd
	f.runArgs = append([]string(nil), args...)
	return f.runResult, f.runErr
}

func TestClientWithoutServiceReturnsNotConnected(t *testing.T) {
	c := &Client{}

	if err := c.Close(); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("Close error = %v, want %v", err, service.ErrNotConnected)
	}
	if err := c.Ping(); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("Ping error = %v, want %v", err, service.ErrNotConnected)
	}
	if _, err := c.Run("/system/identity/print"); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("Run error = %v, want %v", err, service.ErrNotConnected)
	}
	if _, err := c.System().PrintIdentity(); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("System PrintIdentity error = %v, want %v", err, service.ErrNotConnected)
	}
}

func TestNilClientReturnsNotConnected(t *testing.T) {
	var c *Client

	if err := c.Ping(); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("Ping error = %v, want %v", err, service.ErrNotConnected)
	}
	if _, err := c.Run("/system/identity/print"); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("Run error = %v, want %v", err, service.ErrNotConnected)
	}
	if c.Service() != nil {
		t.Fatal("Service() = non-nil, want nil")
	}
	if _, err := c.System().PrintIdentity(); !errors.Is(err, service.ErrNotConnected) {
		t.Fatalf("System PrintIdentity error = %v, want %v", err, service.ErrNotConnected)
	}
}

func TestClientDelegatesToTransport(t *testing.T) {
	transport := &fakeTransport{runResult: map[string]string{"name": "router"}}
	c := &Client{service: transport}

	if err := c.Ping(); err != nil {
		t.Fatalf("Ping error = %v", err)
	}
	if !transport.pingCalled {
		t.Fatal("Ping did not call transport")
	}

	res, err := c.Run("/system/identity/print", "=.proplist=name")
	if err != nil {
		t.Fatalf("Run error = %v", err)
	}
	if res["name"] != "router" {
		t.Fatalf("Run result name = %q, want router", res["name"])
	}
	if transport.runCmd != "/system/identity/print" {
		t.Fatalf("Run command = %q", transport.runCmd)
	}
	if len(transport.runArgs) != 1 || transport.runArgs[0] != "=.proplist=name" {
		t.Fatalf("Run args = %#v", transport.runArgs)
	}

	if err := c.Close(); err != nil {
		t.Fatalf("Close error = %v", err)
	}
	if !transport.closeCalled {
		t.Fatal("Close did not call transport")
	}
}
