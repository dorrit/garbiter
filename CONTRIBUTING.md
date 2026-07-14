# Contributing

## Development

Use Go 1.22 or newer. Before submitting a change, run:

```sh
gofmt -w <changed-go-files>
go vet ./...
go test ./...
go test -race ./...
```

## API Changes

- Keep typed APIs small and retain `Client.Run` as the raw escape hatch.
- Preserve RouterOS `!re` data rows separately from `!done` metadata.
- Use pointer fields for partial updates where zero and omitted differ.
- Include `.proplist` on typed print commands.
- Add fixtures or unit tests for RouterOS v6/v7 response differences.
- Avoid new dependencies unless they materially reduce maintenance. Pin every
  added dependency to a fixed module version.

## Commits

Keep commits focused and include tests with behavior changes. Do not include
credentials, device exports, private certificates, or production addresses.
