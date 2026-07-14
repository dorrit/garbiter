package garbiter

import (
	"context"

	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

// Client is the typed entrypoint for RouterOS interactions.
type Client struct {
	service model.Transport
}

// System provides typed system commands.
func (c *Client) System() *SystemAPI {
	if c == nil {
		return &SystemAPI{}
	}
	return &SystemAPI{svc: c.service}
}

// Interface provides typed interface commands.
func (c *Client) Interface() *InterfaceAPI {
	if c == nil {
		return &InterfaceAPI{}
	}
	return &InterfaceAPI{svc: c.service}
}

// IP provides typed IP commands.
func (c *Client) IP() *IPAPI {
	if c == nil {
		return &IPAPI{}
	}
	return &IPAPI{svc: c.service}
}

// DHCP provides typed DHCP commands.
func (c *Client) DHCP() *DHCPAPI {
	if c == nil {
		return &DHCPAPI{}
	}
	return &DHCPAPI{svc: c.service}
}

// Firewall provides typed firewall commands.
func (c *Client) Firewall() *FirewallAPI {
	if c == nil {
		return &FirewallAPI{}
	}
	return &FirewallAPI{svc: c.service}
}

// Queue provides typed queue commands.
func (c *Client) Queue() *QueueAPI {
	if c == nil {
		return &QueueAPI{}
	}
	return &QueueAPI{svc: c.service}
}

// Log provides typed log commands.
func (c *Client) Log() *LogAPI {
	if c == nil {
		return &LogAPI{}
	}
	return &LogAPI{svc: c.service}
}

// User provides typed user commands.
func (c *Client) User() *UserAPI {
	if c == nil {
		return &UserAPI{}
	}
	return &UserAPI{svc: c.service}
}

// Tool provides typed tool commands.
func (c *Client) Tool() *ToolAPI {
	if c == nil {
		return &ToolAPI{}
	}
	return &ToolAPI{svc: c.service}
}

// PPP provides typed PPP commands.
func (c *Client) PPP() *PPPAPI {
	if c == nil {
		return &PPPAPI{}
	}
	return &PPPAPI{svc: c.service}
}

// Hotspot provides typed hotspot commands.
func (c *Client) Hotspot() *HotspotAPI {
	if c == nil {
		return &HotspotAPI{}
	}
	return &HotspotAPI{svc: c.service}
}

// Certificate provides typed certificate commands.
func (c *Client) Certificate() *CertificateAPI {
	if c == nil {
		return &CertificateAPI{}
	}
	return &CertificateAPI{svc: c.service}
}

// SNMP provides typed SNMP commands.
func (c *Client) SNMP() *SNMPAPI {
	if c == nil {
		return &SNMPAPI{}
	}
	return &SNMPAPI{svc: c.service}
}

// Schedule provides typed scheduler commands.
func (c *Client) Schedule() *ScheduleAPI {
	if c == nil {
		return &ScheduleAPI{}
	}
	return &ScheduleAPI{svc: c.service}
}

// Script provides typed script commands.
func (c *Client) Script() *ScriptAPI {
	if c == nil {
		return &ScriptAPI{}
	}
	return &ScriptAPI{svc: c.service}
}

// Close closes the underlying RouterOS connection.
func (c *Client) Close() error {
	if c == nil || c.service == nil {
		return service.ErrNotConnected
	}
	return c.service.Close()
}

// Ping performs a lightweight connectivity check.
func (c *Client) Ping() error {
	if c == nil || c.service == nil {
		return service.ErrNotConnected
	}
	return c.service.Ping()
}

// Run executes a raw RouterOS command and returns the untyped map response.
func (c *Client) Run(cmd string, args ...string) (map[string]string, error) {
	if c == nil || c.service == nil {
		return nil, service.ErrNotConnected
	}
	if cmd == "" {
		return nil, service.ErrInvalidCommand
	}
	return c.service.Run(cmd, args...)
}

// RunContext executes a raw RouterOS command with context cancellation.
func (c *Client) RunContext(ctx context.Context, cmd string, args ...string) (map[string]string, error) {
	if c == nil || c.service == nil {
		return nil, service.ErrNotConnected
	}
	if cmd == "" {
		return nil, service.ErrInvalidCommand
	}
	return c.service.RunContext(ctx, cmd, args...)
}

// RunList executes a raw RouterOS command and returns all data rows.
func (c *Client) RunList(cmd string, args ...string) ([]map[string]string, error) {
	if c == nil || c.service == nil {
		return nil, service.ErrNotConnected
	}
	if cmd == "" {
		return nil, service.ErrInvalidCommand
	}
	return c.service.RunList(cmd, args...)
}

// RunListContext executes a raw RouterOS command with context cancellation and
// returns all data rows.
func (c *Client) RunListContext(ctx context.Context, cmd string, args ...string) ([]map[string]string, error) {
	if c == nil || c.service == nil {
		return nil, service.ErrNotConnected
	}
	if cmd == "" {
		return nil, service.ErrInvalidCommand
	}
	return c.service.RunListContext(ctx, cmd, args...)
}

// Service exposes the underlying transport for advanced use-cases.
func (c *Client) Service() model.Transport {
	if c == nil {
		return nil
	}
	return c.service
}
