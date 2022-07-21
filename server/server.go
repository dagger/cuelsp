// Package server is a simple abstraction of glsp server that includes in
// custom handler that supports CUE language.
// For more information about glsp: https://github.com/tliron/glsp
package server

import (
	"github.com/dagger/daggerlsp/server/handler"
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"
)

// LSP stands for Language Server Protocol, this type represents the server
type LSP struct {
	handler *handler.Handler

	server *server.Server

	log logging.Logger
}

const (
	Name    = "daggerlsp"
	Version = "0.0.1"
)

type Mode handler.Mode

const (
	DEV Mode = iota
	PROD
)

// New initializes a new language protocol server that contains his logger
// and his handler
func New(mode Mode) *LSP {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/daggerlsp.log"
	// logging.Configure(2, &logTo)
	if mode == DEV {
		logging.Configure(2, nil)
	} else {
		logging.Configure(0, nil)
	}
	log := logging.GetLogger(Name)

	h := handler.New(Name, Version, log, handler.Mode(mode))
	return &LSP{
		log:     log,
		handler: h,
		server:  server.NewServer(h.Handler(), Name, false),
	}
}

// Run will start the server
func (s *LSP) Run() error {
	s.log.Info("Run server Stdio")

	return s.server.RunStdio()
}
