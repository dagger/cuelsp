// Package server is a simple abstraction of glsp server that includes in
// custom handler that supports CUE language.
// For more information about glsp: https://github.com/tliron/glsp
package server

import (
	"github.com/dagger/dlsp/server/handler"
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
	Name    = "dlsp"
	Version = "0.0.1"
)

// New initializes a new language protocol server that contains his logger
// and his handler
func New() *LSP {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/dlsp.log"
	// logging.Configure(2, &logTo)
	logging.Configure(2, nil)
	log := logging.GetLogger(Name)

	h := handler.New(Name, Version, log)
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
