package app

import "errors"

var (
	errServerShutdown = errors.New("server shutdown error")
	errStartupServer  = errors.New("server startup error")
)
