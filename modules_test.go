package garbiter

import (
	"reflect"
	"testing"

	"github.com/dorrit/garbiter/model"
)

func TestInterfacePrintAndSet(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*1", "name": "ether1", "running": "true", "disabled": "false"}}}
	api := (&Client{service: transport}).Interface()

	items, err := api.Print()
	if err != nil {
		t.Fatalf("Print error = %v", err)
	}
	if transport.listCmd != "/interface/print" {
		t.Fatalf("RunList command = %q", transport.listCmd)
	}
	if len(items) != 1 || items[0].Name != "ether1" || !items[0].Running || items[0].Disabled {
		t.Fatalf("Interface items = %#v", items)
	}

	disabled := true
	if err := api.Set("*1", model.InterfaceSet{Name: "wan", Disabled: &disabled}); err != nil {
		t.Fatalf("Set error = %v", err)
	}
	wantArgs := []string{"=.id=*1", "=name=wan", "=disabled=yes"}
	if transport.runCmd != "/interface/set" || !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run = %q %#v, want /interface/set %#v", transport.runCmd, transport.runArgs, wantArgs)
	}
}

func TestIPAddressAddAndDNSSet(t *testing.T) {
	transport := &fakeTransport{}
	api := (&Client{service: transport}).IP()

	_, err := api.AddAddress(model.IPAddressSet{Address: "192.168.1.1/24", Interface: "bridge"})
	if err != nil {
		t.Fatalf("AddAddress error = %v", err)
	}
	wantAddressArgs := []string{"=address=192.168.1.1/24", "=interface=bridge"}
	if transport.runCmd != "/ip/address/add" || !reflect.DeepEqual(transport.runArgs, wantAddressArgs) {
		t.Fatalf("Run = %q %#v, want /ip/address/add %#v", transport.runCmd, transport.runArgs, wantAddressArgs)
	}

	allow := true
	if err := api.SetDNS(model.DNSSet{Servers: "1.1.1.1,8.8.8.8", AllowRemoteRequests: &allow}); err != nil {
		t.Fatalf("SetDNS error = %v", err)
	}
	wantDNSArgs := []string{"=servers=1.1.1.1,8.8.8.8", "=allow-remote-requests=yes"}
	if transport.runCmd != "/ip/dns/set" || !reflect.DeepEqual(transport.runArgs, wantDNSArgs) {
		t.Fatalf("Run = %q %#v, want /ip/dns/set %#v", transport.runCmd, transport.runArgs, wantDNSArgs)
	}
}

func TestDHCPLeasesAndMakeStatic(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*2", "address": "192.168.1.10", "mac-address": "AA:BB", "dynamic": "yes"}}}
	api := (&Client{service: transport}).DHCP()

	leases, err := api.Leases()
	if err != nil {
		t.Fatalf("Leases error = %v", err)
	}
	if transport.listCmd != "/ip/dhcp-server/lease/print" {
		t.Fatalf("RunList command = %q", transport.listCmd)
	}
	if len(leases) != 1 || leases[0].Address != "192.168.1.10" || !leases[0].Dynamic {
		t.Fatalf("Leases = %#v", leases)
	}

	if err := api.MakeLeaseStatic("*2"); err != nil {
		t.Fatalf("MakeLeaseStatic error = %v", err)
	}
	wantArgs := []string{"=.id=*2"}
	if transport.runCmd != "/ip/dhcp-server/lease/make-static" || !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run = %q %#v, want make-static %#v", transport.runCmd, transport.runArgs, wantArgs)
	}
}

func TestFirewallAddFilterAndAddressList(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*3", "chain": "input", "action": "accept", "disabled": "no"}}}
	api := (&Client{service: transport}).Firewall()

	rules, err := api.FilterRules()
	if err != nil {
		t.Fatalf("FilterRules error = %v", err)
	}
	if transport.listCmd != "/ip/firewall/filter/print" {
		t.Fatalf("RunList command = %q", transport.listCmd)
	}
	if len(rules) != 1 || rules[0].Chain != "input" || rules[0].Disabled {
		t.Fatalf("Rules = %#v", rules)
	}

	_, err = api.AddAddressList(model.AddressListSet{List: "blocked", Address: "10.0.0.1"})
	if err != nil {
		t.Fatalf("AddAddressList error = %v", err)
	}
	wantArgs := []string{"=list=blocked", "=address=10.0.0.1"}
	if transport.runCmd != "/ip/firewall/address-list/add" || !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run = %q %#v, want address-list/add %#v", transport.runCmd, transport.runArgs, wantArgs)
	}
}

func TestQueueSimple(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*4", "name": "client", "target": "192.168.1.10/32", "disabled": "false"}}}
	api := (&Client{service: transport}).Queue()

	queues, err := api.Simple()
	if err != nil {
		t.Fatalf("Simple error = %v", err)
	}
	if transport.listCmd != "/queue/simple/print" {
		t.Fatalf("RunList command = %q", transport.listCmd)
	}
	if len(queues) != 1 || queues[0].Name != "client" || queues[0].Disabled {
		t.Fatalf("Queues = %#v", queues)
	}

	_, err = api.AddSimple(model.SimpleQueueSet{Name: "client", Target: "192.168.1.10/32", MaxLimit: "10M/10M"})
	if err != nil {
		t.Fatalf("AddSimple error = %v", err)
	}
	wantArgs := []string{"=name=client", "=target=192.168.1.10/32", "=max-limit=10M/10M"}
	if transport.runCmd != "/queue/simple/add" || !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run = %q %#v, want queue/simple/add %#v", transport.runCmd, transport.runArgs, wantArgs)
	}
}
