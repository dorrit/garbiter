package garbiter

import (
	"reflect"
	"testing"

	"github.com/dorrit/garbiter/model"
)

func TestPPPProfilesSecretsAndActive(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*1", "name": "default", "local-address": "10.0.0.1"}}}
	api := (&Client{service: transport}).PPP()

	profiles, err := api.Profiles()
	if err != nil {
		t.Fatalf("Profiles error = %v", err)
	}
	if transport.listCmd != "/ppp/profile/print" || len(profiles) != 1 || profiles[0].Name != "default" {
		t.Fatalf("Profiles = %#v command = %q", profiles, transport.listCmd)
	}

	disabled := true
	_, err = api.AddSecret(model.PPPSecretSet{Name: "client", Password: "secret", Service: "pppoe", Profile: "default", Disabled: &disabled})
	if err != nil {
		t.Fatalf("AddSecret error = %v", err)
	}
	wantArgs := []string{"=name=client", "=password=secret", "=service=pppoe", "=profile=default", "=disabled=yes"}
	if transport.runCmd != "/ppp/secret/add" || !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run = %q %#v, want ppp/secret/add %#v", transport.runCmd, transport.runArgs, wantArgs)
	}

	if err := api.RemoveActive("*2"); err != nil {
		t.Fatalf("RemoveActive error = %v", err)
	}
	if transport.runCmd != "/ppp/active/remove" || !reflect.DeepEqual(transport.runArgs, []string{"=.id=*2"}) {
		t.Fatalf("Run = %q %#v, want ppp/active/remove", transport.runCmd, transport.runArgs)
	}
}

func TestHotspotServerUserAndActive(t *testing.T) {
	transport := &fakeTransport{listResult: []map[string]string{{".id": "*1", "name": "hs", "interface": "bridge", "disabled": "no"}}}
	api := (&Client{service: transport}).Hotspot()

	servers, err := api.Servers()
	if err != nil {
		t.Fatalf("Servers error = %v", err)
	}
	if transport.listCmd != "/ip/hotspot/print" || len(servers) != 1 || servers[0].Name != "hs" {
		t.Fatalf("Servers = %#v command = %q", servers, transport.listCmd)
	}

	_, err = api.AddUser(model.HotspotUserSet{Name: "guest", Password: "secret", Profile: "default"})
	if err != nil {
		t.Fatalf("AddUser error = %v", err)
	}
	wantArgs := []string{"=name=guest", "=password=secret", "=profile=default"}
	if transport.runCmd != "/ip/hotspot/user/add" || !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run = %q %#v, want hotspot/user/add %#v", transport.runCmd, transport.runArgs, wantArgs)
	}

	if err := api.DisableUser("*3"); err != nil {
		t.Fatalf("DisableUser error = %v", err)
	}
	if transport.runCmd != "/ip/hotspot/user/disable" || !reflect.DeepEqual(transport.runArgs, []string{"=.id=*3"}) {
		t.Fatalf("Run = %q %#v, want hotspot/user/disable", transport.runCmd, transport.runArgs)
	}
}

func TestCertificateSNMPScheduleAndScript(t *testing.T) {
	c := &Client{service: &fakeTransport{listResult: []map[string]string{{".id": "*1", "name": "cert", "trusted": "true"}}}}
	transport := c.service.(*fakeTransport)

	certs, err := c.Certificate().Print()
	if err != nil {
		t.Fatalf("Certificate Print error = %v", err)
	}
	if transport.listCmd != "/certificate/print" || len(certs) != 1 || !certs[0].Trusted {
		t.Fatalf("Certificates = %#v command = %q", certs, transport.listCmd)
	}

	enabled := true
	if err := c.SNMP().Set(model.SNMPSet{Enabled: &enabled, Contact: "noc"}); err != nil {
		t.Fatalf("SNMP Set error = %v", err)
	}
	wantSNMPArgs := []string{"=enabled=yes", "=contact=noc"}
	if transport.runCmd != "/snmp/set" || !reflect.DeepEqual(transport.runArgs, wantSNMPArgs) {
		t.Fatalf("Run = %q %#v, want snmp/set %#v", transport.runCmd, transport.runArgs, wantSNMPArgs)
	}

	_, err = c.Schedule().Add(model.ScheduleSet{Name: "backup", Interval: "1d", OnEvent: "export"})
	if err != nil {
		t.Fatalf("Schedule Add error = %v", err)
	}
	wantScheduleArgs := []string{"=name=backup", "=interval=1d", "=on-event=export"}
	if transport.runCmd != "/system/scheduler/add" || !reflect.DeepEqual(transport.runArgs, wantScheduleArgs) {
		t.Fatalf("Run = %q %#v, want scheduler/add %#v", transport.runCmd, transport.runArgs, wantScheduleArgs)
	}

	if err := c.Script().Run("*4"); err != nil {
		t.Fatalf("Script Run error = %v", err)
	}
	if transport.runCmd != "/system/script/run" || !reflect.DeepEqual(transport.runArgs, []string{"=.id=*4"}) {
		t.Fatalf("Run = %q %#v, want script/run", transport.runCmd, transport.runArgs)
	}
}
