package garbiter_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/dorrit/garbiter"
	"github.com/dorrit/garbiter/model"
)

func ExampleConnectTLS() {
	client, err := garbiter.ConnectTLS(
		"router.example.com:8729",
		"admin",
		"",
		&tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: "router.example.com",
		},
		garbiter.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	identity, err := client.System().PrintIdentity()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(identity.Name)
}

func ExampleIPAPI_AddAddress() {
	client, err := garbiter.NewClient(exampleTransport{})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.IP().AddAddress(model.IPAddressSet{
		Address:   "192.168.88.1/24",
		Interface: "bridge",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleFirewallAPI_AddAddressList() {
	client, err := garbiter.NewClient(exampleTransport{})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Firewall().AddAddressList(model.AddressListSet{
		List:    "blocked",
		Address: "203.0.113.10",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_Run() {
	client, err := garbiter.NewClient(exampleTransport{})
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Run("/system/identity/print")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res["name"])
}

type exampleTransport struct{}

func (exampleTransport) Connect(string, string, string) error { return nil }
func (exampleTransport) ConnectContext(context.Context, string, string, string) error {
	return nil
}
func (exampleTransport) ConnectTLS(string, string, string, *tls.Config) error { return nil }
func (exampleTransport) ConnectTLSContext(context.Context, string, string, string, *tls.Config) error {
	return nil
}
func (exampleTransport) Close() error { return nil }
func (exampleTransport) Ping() error  { return nil }
func (exampleTransport) Run(string, ...string) (map[string]string, error) {
	return map[string]string{"name": "router"}, nil
}
func (exampleTransport) RunContext(context.Context, string, ...string) (map[string]string, error) {
	return map[string]string{"name": "router"}, nil
}
func (exampleTransport) RunList(string, ...string) ([]map[string]string, error) {
	return []map[string]string{}, nil
}
func (exampleTransport) RunListContext(context.Context, string, ...string) ([]map[string]string, error) {
	return []map[string]string{}, nil
}
