# Security Policy

## Supported Versions

Security fixes are applied to the latest `v0.x` release line. Pre-`v1.0.0`
APIs may change when a security or protocol-correctness fix requires it.

## Reporting a Vulnerability

Please report vulnerabilities through GitHub's private security advisory flow
for this repository. Do not open a public issue for credential exposure, TLS
downgrade, authentication bypass, or remote command execution concerns.

Include the affected garbiter version, RouterOS version, reproduction steps,
and expected impact. Reports will be acknowledged as soon as practical.

## Transport Security

Use `ConnectTLS` with certificate verification on RouterOS API port `8729`.
Plaintext `Connect` should only be used on a trusted, protected network.
