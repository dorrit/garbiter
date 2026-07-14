package garbiter

import (
	"fmt"

	"github.com/dorrit/garbiter/service"
)

// ErrNotConnected indicates that a command requires an active connection.
var ErrNotConnected = service.ErrNotConnected

// ErrInvalidID indicates that an item command received an empty RouterOS ID.
var ErrInvalidID = service.ErrInvalidID

// ErrInvalidCommand indicates that a raw command path is empty.
var ErrInvalidCommand = service.ErrInvalidCommand

// ErrInvalidTLSConfig indicates that ConnectTLS received a nil configuration.
var ErrInvalidTLSConfig = service.ErrInvalidTLSConfig

// ValidationError describes an invalid public API input field.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("routeros: invalid %s: %s", e.Field, e.Message)
}

func requireField(field, value string) error {
	if value == "" {
		return &ValidationError{Field: field, Message: "is required"}
	}
	return nil
}
