package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

// Client is the typed entrypoint for RouterOS interactions.
type Client struct {
	service  model.Transport
	system   *SystemAPI
	iface    *InterfaceAPI
	ip       *IPAPI
	dhcp     *DHCPAPI
	firewall *FirewallAPI
	queue    *QueueAPI
}

// System provides typed system commands.
func (c *Client) System() *SystemAPI {
	if c == nil {
		return &SystemAPI{}
	}
	if c.system == nil {
		c.system = &SystemAPI{svc: c.service}
	}
	return c.system
}

// Interface provides typed interface commands.
func (c *Client) Interface() *InterfaceAPI {
	if c == nil {
		return &InterfaceAPI{}
	}
	if c.iface == nil {
		c.iface = &InterfaceAPI{svc: c.service}
	}
	return c.iface
}

// IP provides typed IP commands.
func (c *Client) IP() *IPAPI {
	if c == nil {
		return &IPAPI{}
	}
	if c.ip == nil {
		c.ip = &IPAPI{svc: c.service}
	}
	return c.ip
}

// DHCP provides typed DHCP commands.
func (c *Client) DHCP() *DHCPAPI {
	if c == nil {
		return &DHCPAPI{}
	}
	if c.dhcp == nil {
		c.dhcp = &DHCPAPI{svc: c.service}
	}
	return c.dhcp
}

// Firewall provides typed firewall commands.
func (c *Client) Firewall() *FirewallAPI {
	if c == nil {
		return &FirewallAPI{}
	}
	if c.firewall == nil {
		c.firewall = &FirewallAPI{svc: c.service}
	}
	return c.firewall
}

// Queue provides typed queue commands.
func (c *Client) Queue() *QueueAPI {
	if c == nil {
		return &QueueAPI{}
	}
	if c.queue == nil {
		c.queue = &QueueAPI{svc: c.service}
	}
	return c.queue
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
	return c.service.Run(cmd, args...)
}

// Service exposes the underlying transport for advanced use-cases.
func (c *Client) Service() model.Transport {
	if c == nil {
		return nil
	}
	return c.service
}
