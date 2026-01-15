# garbiter

Typed, minimal RouterOS client built on top of `go-routeros`.

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
- `Client` (root package) exposes typed modules (starting with `System`) that return structs instead of raw maps.
- Options are passed on construction (`garbiter.Connect`/`ConnectTLS`) to configure the transport (e.g., timeouts).
