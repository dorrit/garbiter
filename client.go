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
	log      *LogAPI
	user     *UserAPI
	tool     *ToolAPI
	ppp      *PPPAPI
	hotspot  *HotspotAPI
	cert     *CertificateAPI
	snmp     *SNMPAPI
	schedule *ScheduleAPI
	script   *ScriptAPI
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

// Log provides typed log commands.
func (c *Client) Log() *LogAPI {
	if c == nil {
		return &LogAPI{}
	}
	if c.log == nil {
		c.log = &LogAPI{svc: c.service}
	}
	return c.log
}

// User provides typed user commands.
func (c *Client) User() *UserAPI {
	if c == nil {
		return &UserAPI{}
	}
	if c.user == nil {
		c.user = &UserAPI{svc: c.service}
	}
	return c.user
}

// Tool provides typed tool commands.
func (c *Client) Tool() *ToolAPI {
	if c == nil {
		return &ToolAPI{}
	}
	if c.tool == nil {
		c.tool = &ToolAPI{svc: c.service}
	}
	return c.tool
}

// PPP provides typed PPP commands.
func (c *Client) PPP() *PPPAPI {
	if c == nil {
		return &PPPAPI{}
	}
	if c.ppp == nil {
		c.ppp = &PPPAPI{svc: c.service}
	}
	return c.ppp
}

// Hotspot provides typed hotspot commands.
func (c *Client) Hotspot() *HotspotAPI {
	if c == nil {
		return &HotspotAPI{}
	}
	if c.hotspot == nil {
		c.hotspot = &HotspotAPI{svc: c.service}
	}
	return c.hotspot
}

// Certificate provides typed certificate commands.
func (c *Client) Certificate() *CertificateAPI {
	if c == nil {
		return &CertificateAPI{}
	}
	if c.cert == nil {
		c.cert = &CertificateAPI{svc: c.service}
	}
	return c.cert
}

// SNMP provides typed SNMP commands.
func (c *Client) SNMP() *SNMPAPI {
	if c == nil {
		return &SNMPAPI{}
	}
	if c.snmp == nil {
		c.snmp = &SNMPAPI{svc: c.service}
	}
	return c.snmp
}

// Schedule provides typed scheduler commands.
func (c *Client) Schedule() *ScheduleAPI {
	if c == nil {
		return &ScheduleAPI{}
	}
	if c.schedule == nil {
		c.schedule = &ScheduleAPI{svc: c.service}
	}
	return c.schedule
}

// Script provides typed script commands.
func (c *Client) Script() *ScriptAPI {
	if c == nil {
		return &ScriptAPI{}
	}
	if c.script == nil {
		c.script = &ScriptAPI{svc: c.service}
	}
	return c.script
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
