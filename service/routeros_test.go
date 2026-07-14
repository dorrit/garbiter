package service

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	routeros "github.com/go-routeros/routeros/v3"
	"github.com/go-routeros/routeros/v3/proto"
)

type fakeRouterOSClient struct {
	run func(context.Context, ...string) (*routeros.Reply, error)
}

func (f *fakeRouterOSClient) RunContext(ctx context.Context, sentences ...string) (*routeros.Reply, error) {
	return f.run(ctx, sentences...)
}

func (f *fakeRouterOSClient) Close() error {
	return nil
}

func TestRunPrefersDataRowOverDoneMetadata(t *testing.T) {
	svc := NewRouterOSService()
	svc.client = &fakeRouterOSClient{run: func(context.Context, ...string) (*routeros.Reply, error) {
		return &routeros.Reply{
			Re:   []*proto.Sentence{{Map: map[string]string{"name": "router"}}},
			Done: &proto.Sentence{Map: map[string]string{"ret": "*1"}},
		}, nil
	}}

	row, err := svc.Run("/system/identity/print")
	if err != nil {
		t.Fatalf("Run error = %v", err)
	}
	if row["name"] != "router" {
		t.Fatalf("Run row = %#v, want data row", row)
	}
}

func TestRunListDoesNotPromoteDoneMetadata(t *testing.T) {
	svc := NewRouterOSService()
	svc.client = &fakeRouterOSClient{run: func(context.Context, ...string) (*routeros.Reply, error) {
		return &routeros.Reply{Done: &proto.Sentence{Map: map[string]string{"ret": "*1"}}}, nil
	}}

	rows, err := svc.RunList("/interface/print")
	if err != nil {
		t.Fatalf("RunList error = %v", err)
	}
	if len(rows) != 0 {
		t.Fatalf("RunList rows = %#v, want empty", rows)
	}
}

func TestRunAppliesDefaultCommandTimeout(t *testing.T) {
	svc := NewRouterOSService(WithCommandTimeout(time.Millisecond))
	svc.client = &fakeRouterOSClient{run: func(ctx context.Context, _ ...string) (*routeros.Reply, error) {
		<-ctx.Done()
		return nil, ctx.Err()
	}}

	_, err := svc.Run("/system/identity/print")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Run error = %v, want deadline exceeded", err)
	}
}

func TestRunSerializesConcurrentCommands(t *testing.T) {
	svc := NewRouterOSService()
	var stateMu sync.Mutex
	active := 0
	maxActive := 0
	svc.client = &fakeRouterOSClient{run: func(context.Context, ...string) (*routeros.Reply, error) {
		stateMu.Lock()
		active++
		if active > maxActive {
			maxActive = active
		}
		stateMu.Unlock()

		time.Sleep(time.Millisecond)

		stateMu.Lock()
		active--
		stateMu.Unlock()
		return &routeros.Reply{}, nil
	}}

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := svc.Run("/system/identity/print"); err != nil {
				t.Errorf("Run error = %v", err)
			}
		}()
	}
	wg.Wait()

	if maxActive != 1 {
		t.Fatalf("maximum concurrent commands = %d, want 1", maxActive)
	}
}
