package service

import "errors"

var ErrNotConnected = errors.New("routeros: not connected")

var ErrInvalidID = errors.New("routeros: invalid empty id")

var ErrInvalidCommand = errors.New("routeros: invalid empty command")

var ErrInvalidTLSConfig = errors.New("routeros: TLS config is required")
