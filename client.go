package garbiter

import "github.com/dorrit/garbiter/model"

// Client is the typed entrypoint for RouterOS interactions.
type Client struct {
	service model.Transport
	system  *SystemAPI
}

// System provides typed system commands.
func (c *Client) System() *SystemAPI {
	if c.system == nil {
		c.system = &SystemAPI{svc: c.service}
	}
	return c.system
}

// Close closes the underlying RouterOS connection.
func (c *Client) Close() error {
	if c.service == nil {
		return nil
	}
	return c.service.Close()
}

// Ping performs a lightweight connectivity check.
func (c *Client) Ping() error {
	if c.service == nil {
		return nil
	}
	return c.service.Ping()
}

// Run executes a raw RouterOS command and returns the untyped map response.
func (c *Client) Run(cmd string, args ...string) (map[string]string, error) {
	if c.service == nil {
		return nil, nil
	}
	return c.service.Run(cmd, args...)
}

// Service exposes the underlying transport for advanced use-cases.
func (c *Client) Service() model.Transport {
	return c.service
}
