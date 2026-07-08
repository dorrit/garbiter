package garbiter

import (
	"reflect"
	"testing"

	"github.com/dorrit/garbiter/model"
)

func TestSystemClockPackagesAndRouterboard(t *testing.T) {
	transport := &fakeTransport{runResult: map[string]string{"time": "12:00:00", "date": "jul/07/2026", "time-zone-autodetect": "yes"}}
	system := (&Client{service: transport}).System()

	clock, err := system.PrintClock()
	if err != nil {
		t.Fatalf("PrintClock error = %v", err)
	}
	if transport.runCmd != "/system/clock/print" || !clock.TimeZoneAuto {
		t.Fatalf("Clock = %#v command = %q", clock, transport.runCmd)
	}

	auto := false
	if err := system.SetClock(model.ClockSet{TimeZoneName: "Asia/Jakarta", TimeZoneAuto: &auto}); err != nil {
		t.Fatalf("SetClock error = %v", err)
	}
	wantClockArgs := []string{"=time-zone-name=Asia/Jakarta", "=time-zone-autodetect=no"}
	if transport.runCmd != "/system/clock/set" || !reflect.DeepEqual(transport.runArgs, wantClockArgs) {
		t.Fatalf("Run = %q %#v, want clock/set %#v", transport.runCmd, transport.runArgs, wantClockArgs)
	}

	transport.listResult = []map[string]string{{".id": "*1", "name": "routeros", "version": "7.15", "disabled": "false"}}
	packages, err := system.Packages()
	if err != nil {
		t.Fatalf("Packages error = %v", err)
	}
	if transport.listCmd != "/system/package/print" || len(packages) != 1 || packages[0].Name != "routeros" {
		t.Fatalf("Packages = %#v command = %q", packages, transport.listCmd)
	}

	transport.runResult = map[string]string{"routerboard": "true", "model": "RB5009", "firmware-upgrade-needed": "yes"}
	routerboard, err := system.PrintRouterboard()
	if err != nil {
		t.Fatalf("PrintRouterboard error = %v", err)
	}
	if transport.runCmd != "/system/routerboard/print" || !routerboard.Routerboard || !routerboard.FirmwareUpgradeNeeded {
		t.Fatalf("Routerboard = %#v command = %q", routerboard, transport.runCmd)
	}
}

func TestIPServicesARPAndPools(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*1", "name": "www", "port": "80", "disabled": "yes"}}}
	api := (&Client{service: transport}).IP()

	services, err := api.Services()
	if err != nil {
		t.Fatalf("Services error = %v", err)
	}
	if transport.listCmd != "/ip/service/print" || len(services) != 1 || !services[0].Disabled {
		t.Fatalf("Services = %#v command = %q", services, transport.listCmd)
	}

	disabled := false
	if err := api.SetService("*1", model.IPServiceSet{Port: "8080", Disabled: &disabled}); err != nil {
		t.Fatalf("SetService error = %v", err)
	}
	wantServiceArgs := []string{"=.id=*1", "=port=8080", "=disabled=no"}
	if transport.runCmd != "/ip/service/set" || !reflect.DeepEqual(transport.runArgs, wantServiceArgs) {
		t.Fatalf("Run = %q %#v, want service/set %#v", transport.runCmd, transport.runArgs, wantServiceArgs)
	}

	_, err = api.AddARP(model.ARPSet{Address: "192.168.1.10", MACAddress: "AA:BB", Interface: "bridge"})
	if err != nil {
		t.Fatalf("AddARP error = %v", err)
	}
	wantARPArgs := []string{"=address=192.168.1.10", "=mac-address=AA:BB", "=interface=bridge"}
	if transport.runCmd != "/ip/arp/add" || !reflect.DeepEqual(transport.runArgs, wantARPArgs) {
		t.Fatalf("Run = %q %#v, want arp/add %#v", transport.runCmd, transport.runArgs, wantARPArgs)
	}

	_, err = api.AddPool(model.PoolSet{Name: "lan", Ranges: "192.168.1.10-192.168.1.100"})
	if err != nil {
		t.Fatalf("AddPool error = %v", err)
	}
	wantPoolArgs := []string{"=name=lan", "=ranges=192.168.1.10-192.168.1.100"}
	if transport.runCmd != "/ip/pool/add" || !reflect.DeepEqual(transport.runArgs, wantPoolArgs) {
		t.Fatalf("Run = %q %#v, want pool/add %#v", transport.runCmd, transport.runArgs, wantPoolArgs)
	}
}

func TestLogUserAndTool(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*1", "time": "12:00:00", "topics": "system", "message": "started"}}}
	c := &Client{service: transport}

	logs, err := c.Log().Print()
	if err != nil {
		t.Fatalf("Log Print error = %v", err)
	}
	if transport.listCmd != "/log/print" || len(logs) != 1 || logs[0].Message != "started" {
		t.Fatalf("Logs = %#v command = %q", logs, transport.listCmd)
	}

	_, err = c.User().Add(model.UserSet{Name: "ops", Password: "secret", Group: "full"})
	if err != nil {
		t.Fatalf("User Add error = %v", err)
	}
	wantUserArgs := []string{"=name=ops", "=password=secret", "=group=full"}
	if transport.runCmd != "/user/add" || !reflect.DeepEqual(transport.runArgs, wantUserArgs) {
		t.Fatalf("Run = %q %#v, want user/add %#v", transport.runCmd, transport.runArgs, wantUserArgs)
	}

	transport.listResult = []map[string]string{{"host": "1.1.1.1", "seq": "0", "time": "10ms"}}
	results, err := c.Tool().Ping(model.PingRequest{Address: "1.1.1.1", Count: "1"})
	if err != nil {
		t.Fatalf("Tool Ping error = %v", err)
	}
	wantPingArgs := []string{"=address=1.1.1.1", "=count=1"}
	if transport.listCmd != "/ping" || !reflect.DeepEqual(transport.listArgs, wantPingArgs) || len(results) != 1 || results[0].Host != "1.1.1.1" {
		t.Fatalf("Ping = %#v command = %q args = %#v", results, transport.listCmd, transport.listArgs)
	}
}
