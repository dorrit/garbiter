# Compatibility

## Go

The module declares Go 1.22 and CI covers Go 1.22, 1.24, and 1.25.

## RouterOS

The typed API targets RouterOS v7 while retaining compatibility mappings for
known RouterOS v6 response shapes, including legacy system health output.
RouterOS properties vary by device model and installed package; fields that are
not returned remain at their Go zero value and raw result maps are retained where
available.

Hardware-backed integration coverage is not yet part of CI. Consumers should
validate configuration-changing commands against their RouterOS version and
device family before production rollout.

## API Stability

The project is pre-v1. Public APIs can change between minor versions when needed
to correct security, cancellation, update semantics, or RouterOS protocol
behavior. Pin a released module version and review `CHANGELOG.md` when upgrading.
