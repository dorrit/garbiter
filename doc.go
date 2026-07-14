// Package garbiter provides a typed RouterOS client built on top of go-routeros.
//
// The package exposes typed modules for common RouterOS areas such as System,
// Interface, IP, DHCP, Firewall, Queue, PPP, Hotspot, and administration tools.
// Raw RouterOS commands remain available through Client.Run for commands that do
// not yet have a typed wrapper.
//
// Prefer ConnectTLS with certificate verification. Connect uses plaintext
// RouterOS API transport and should be limited to trusted, protected networks.
package garbiter
