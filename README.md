# garbiter

Typed RouterOS client built on top of `go-routeros`.

## Quick start

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dorrit/garbiter"
)

func main() {
	client, err := garbiter.Connect(
		"192.168.31.1:8728",
		"admin",
		"",
		garbiter.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer client.Close()

	if err := client.Ping(); err != nil {
		log.Fatalf("ping: %v", err)
	}

	id, err := client.System().PrintIdentity()
	if err != nil {
		log.Fatalf("identity: %v", err)
	}

	fmt.Println("Router identity:", id.Name)
}
```

## Design
- `service` wraps the `go-routeros` client and handles transport concerns.
- `Client` (root package) exposes typed modules that return structs instead of raw maps.
- Options are passed on construction (`garbiter.Connect`/`ConnectTLS`) to configure the transport (e.g., timeouts).

## Typed Modules
- `System`: identity, resource, health, health settings.
- `Interface`: interface print/set/enable/disable, bridge, VLAN.
- `IP`: address, route, DNS.
- `DHCP`: DHCP client, DHCP server, leases.
- `Firewall`: filter rules, NAT rules, address-list.
- `Queue`: simple queues.
- `Log`: log entries.
- `User`: users and local access accounts.
- `Tool`: ping.
- `PPP`: profiles, secrets, active sessions.
- `Hotspot`: servers, users, active sessions.
- `Certificate`: certificates.
- `SNMP`: SNMP settings.
- `Schedule`: system scheduler entries.
- `Script`: system scripts and script execution.

Raw commands remain available through `Client.Run` and list-style commands through the transport `RunList` method.

## Examples

List interfaces:

```go
interfaces, err := client.Interface().Print()
if err != nil {
	log.Fatal(err)
}
for _, iface := range interfaces {
	fmt.Println(iface.Name, iface.Running)
}
```

Add an IP address:

```go
_, err := client.IP().AddAddress(model.IPAddressSet{
	Address:   "192.168.88.1/24",
	Interface: "bridge",
})
```

Add a firewall address-list entry:

```go
_, err := client.Firewall().AddAddressList(model.AddressListSet{
	List:    "blocked",
	Address: "203.0.113.10",
})
```

Run a raw command:

```go
res, err := client.Run("/system/identity/print")
```

## Versioning
- The library is currently suitable for `v0.x` use while typed APIs are still being expanded.
- Prefer pinning a tagged version once releases are published.
- Raw RouterOS behavior and unsupported commands can still be accessed with `Client.Run`.

## Errors
- Operations that need a RouterOS transport return `service.ErrNotConnected` when the client is nil or not connected.
- Operations that require a RouterOS item id return `service.ErrInvalidID` when the id is empty.
- Transport and RouterOS command errors are returned unchanged by typed APIs.
