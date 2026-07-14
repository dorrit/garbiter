# Changelog

This project follows semantic versioning. Before `v1.0.0`, minor releases may
contain API changes required for RouterOS correctness and security.

## Unreleased

### Added

- Context-aware connect and raw command APIs.
- Configurable command timeout and external transport injection.
- Typed request validation and root-level public errors.
- RouterOS v7 health sensor rows with legacy health response support.
- CI coverage for multiple Go versions, vet, and race detection.

### Changed

- TLS connections reject nil configurations instead of falling back to plaintext.
- Transport operations are serialized for safe synchronous protocol use.
- Typed print operations request explicit property lists.
- Health settings use partial-update pointer semantics.

## v0.1.0

- Initial RouterOS client, system module, CI, and release workflow.
