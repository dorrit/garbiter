package garbiter

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/dorrit/garbiter/model"
)

func TestSystemPrintIdentity(t *testing.T) {
	transport := &fakeTransport{runResult: map[string]string{"name": "edge-router"}}
	system := (&Client{service: transport}).System()

	identity, err := system.PrintIdentity()
	if err != nil {
		t.Fatalf("PrintIdentity error = %v", err)
	}
	if identity.Name != "edge-router" {
		t.Fatalf("Identity name = %q, want edge-router", identity.Name)
	}
	if transport.runCmd != "/system/identity/print" {
		t.Fatalf("Run command = %q", transport.runCmd)
	}
}

func TestSystemSetIdentity(t *testing.T) {
	transport := &fakeTransport{}
	system := (&Client{service: transport}).System()

	if err := system.SetIdentity("core-router"); err != nil {
		t.Fatalf("SetIdentity error = %v", err)
	}
	if transport.runCmd != "/system/identity/set" {
		t.Fatalf("Run command = %q", transport.runCmd)
	}
	wantArgs := []string{"=name=core-router"}
	if !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run args = %#v, want %#v", transport.runArgs, wantArgs)
	}
}

func TestSystemPrintResource(t *testing.T) {
	transport := &fakeTransport{runResult: map[string]string{
		"uptime":                  "1d2h3m",
		"version":                 "7.15.1",
		"build-time":              "Jun/01/2024 10:00:00",
		"factory-software":        "7.1",
		"free-memory":             "128MiB",
		"total-memory":            "256MiB",
		"cpu":                     "ARM",
		"cpu-count":               "4",
		"cpu-frequency":           "1400MHz",
		"cpu-load":                "12%",
		"free-hdd-space":          "64MiB",
		"total-hdd-space":         "128MiB",
		"write-sect-since-reboot": "123",
		"write-sect-total":        "456",
		"bad-blocks":              "0%",
		"architecture-name":       "arm64",
		"board-name":              "RB5009",
		"platform":                "MikroTik",
	}}
	system := (&Client{service: transport}).System()

	resource, err := system.PrintResource()
	if err != nil {
		t.Fatalf("PrintResource error = %v", err)
	}
	if resource.CPUCount != 4 {
		t.Fatalf("CPUCount = %d, want 4", resource.CPUCount)
	}
	if resource.WriteSectSinceBoot != 123 {
		t.Fatalf("WriteSectSinceBoot = %d, want 123", resource.WriteSectSinceBoot)
	}
	if resource.WriteSectTotal != 456 {
		t.Fatalf("WriteSectTotal = %d, want 456", resource.WriteSectTotal)
	}
	if resource.BoardName != "RB5009" || resource.Platform != "MikroTik" {
		t.Fatalf("Resource board/platform = %q/%q", resource.BoardName, resource.Platform)
	}
	if transport.runCmd != "/system/resource/print" {
		t.Fatalf("Run command = %q", transport.runCmd)
	}
}

func TestSystemPrintHealth(t *testing.T) {
	transport := &fakeTransport{runResult: map[string]string{
		"voltage":           "24.2",
		"temperature":       "44.5",
		"power-consumption": "15.7",
		"cpu-temperature":   "55.1",
		"fan1-speed":        "1200",
		"fan2-speed":        "1300",
		"fan3-speed":        "1400",
		"fan4-speed":        "1500",
		"board-temp1":       "42.1",
		"board-temp2":       "43.2",
		"psu1-voltage":      "24.0",
		"psu2-voltage":      "24.1",
		"psu1-current":      "0.7",
		"psu2-current":      "0.8",
	}}
	system := (&Client{service: transport}).System()

	health, err := system.PrintHealth()
	if err != nil {
		t.Fatalf("PrintHealth error = %v", err)
	}
	if health.Voltage != 24.2 || health.Temperature != 44.5 {
		t.Fatalf("Health voltage/temperature = %v/%v", health.Voltage, health.Temperature)
	}
	if health.Fan1Speed != 1200 || health.Fan4Speed != 1500 {
		t.Fatalf("Health fan speeds = %d/%d", health.Fan1Speed, health.Fan4Speed)
	}
	if health.PSU1Current != 0.7 || health.PSU2Current != 0.8 {
		t.Fatalf("Health PSU current = %v/%v", health.PSU1Current, health.PSU2Current)
	}
	if transport.runCmd != "/system/health/print" {
		t.Fatalf("Run command = %q", transport.runCmd)
	}
}

func TestSystemSetHealth(t *testing.T) {
	transport := &fakeTransport{}
	system := (&Client{service: transport}).System()

	err := system.SetHealth(model.HealthSettings{
		CPUOvertempCheck:        true,
		CPUOvertempThreshold:    75,
		CPUOvertempStartupDelay: 3 * time.Second,
		FanMode:                 "auto",
		FanOnThreshold:          50,
		FanSwitch:               "all",
		UseFan:                  true,
	})
	if err != nil {
		t.Fatalf("SetHealth error = %v", err)
	}

	if transport.runCmd != "/system/health/settings/set" {
		t.Fatalf("Run command = %q", transport.runCmd)
	}
	wantArgs := []string{
		"=cpu-overtemp-check=yes",
		"=cpu-overtemp-threshold=75",
		"=cpu-overtemp-startup-delay=3s",
		"=fan-mode=auto",
		"=fan-on-threshold=50",
		"=fan-switch=all",
		"=use-fan=yes",
	}
	if !reflect.DeepEqual(transport.runArgs, wantArgs) {
		t.Fatalf("Run args = %#v, want %#v", transport.runArgs, wantArgs)
	}
}

func TestSystemPropagatesTransportErrors(t *testing.T) {
	wantErr := errors.New("router failed")
	transport := &fakeTransport{runErr: wantErr}
	system := (&Client{service: transport}).System()

	if _, err := system.PrintIdentity(); !errors.Is(err, wantErr) {
		t.Fatalf("PrintIdentity error = %v, want %v", err, wantErr)
	}
	if err := system.SetIdentity("router"); !errors.Is(err, wantErr) {
		t.Fatalf("SetIdentity error = %v, want %v", err, wantErr)
	}
}
