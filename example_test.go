package garbiter_test

import (
	"fmt"
	"log"
	"time"

	"github.com/dorrit/garbiter"
	"github.com/dorrit/garbiter/model"
)

func ExampleConnect() {
	client, err := garbiter.Connect(
		"192.168.88.1:8728",
		"admin",
		"",
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
	client := garbiter.New()

	_, err := client.IP().AddAddress(model.IPAddressSet{
		Address:   "192.168.88.1/24",
		Interface: "bridge",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleFirewallAPI_AddAddressList() {
	client := garbiter.New()

	_, err := client.Firewall().AddAddressList(model.AddressListSet{
		List:    "blocked",
		Address: "203.0.113.10",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleClient_Run() {
	client := garbiter.New()

	res, err := client.Run("/system/identity/print")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res["name"])
}
