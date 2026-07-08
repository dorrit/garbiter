package garbiter

import (
	"errors"
	"testing"

	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

func TestRunRejectsEmptyCommand(t *testing.T) {
	c := &Client{service: &fakeTransport{}}

	if _, err := c.Run(""); !errors.Is(err, service.ErrInvalidCommand) {
		t.Fatalf("Run error = %v, want %v", err, service.ErrInvalidCommand)
	}
}

func TestIDCommandsRejectEmptyID(t *testing.T) {
	c := &Client{service: &fakeTransport{}}

	tests := []struct {
		name string
		call func() error
	}{
		{name: "interface", call: func() error { return c.Interface().Set("", model.InterfaceSet{Name: "wan"}) }},
		{name: "ip", call: func() error { return c.IP().SetAddress("", model.IPAddressSet{Address: "192.168.1.1/24"}) }},
		{name: "dhcp", call: func() error { return c.DHCP().SetClient("", model.DHCPClientSet{Interface: "ether1"}) }},
		{name: "firewall", call: func() error { return c.Firewall().SetFilterRule("", model.FirewallRuleSet{Action: "accept"}) }},
		{name: "queue", call: func() error { return c.Queue().SetSimple("", model.SimpleQueueSet{Name: "client"}) }},
		{name: "user", call: func() error { return c.User().Set("", model.UserSet{Name: "admin"}) }},
		{name: "ppp", call: func() error { return c.PPP().SetSecret("", model.PPPSecretSet{Name: "client"}) }},
		{name: "hotspot", call: func() error { return c.Hotspot().SetUser("", model.HotspotUserSet{Name: "guest"}) }},
		{name: "certificate", call: func() error { return c.Certificate().Set("", model.CertificateSet{Name: "cert"}) }},
		{name: "schedule", call: func() error { return c.Schedule().Set("", model.ScheduleSet{Name: "backup"}) }},
		{name: "script", call: func() error { return c.Script().Run("") }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, service.ErrInvalidID) {
				t.Fatalf("error = %v, want %v", err, service.ErrInvalidID)
			}
		})
	}
}
