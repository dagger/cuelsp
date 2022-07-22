// Package server is a simple abstraction of glsp server that includes in
// custom handler that supports CUE language.
// For more information about glsp: https://github.com/tliron/glsp
package server

import (
	"errors"

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

// New initializes a new language protocol server that contains his logger
// and his handler
func New(mode Mode) (*LSP, error) {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/daggerlsp.log"
	// logging.Configure(2, &logTo)
	switch mode {
	case ModeDev:
		logging.Configure(2, nil)
	case ModeProd:
		logging.Configure(0, nil)
	default:
		return nil, errors.New("unknown logging mode")
	}

	baseLog := logging.GetLogger(Name)
	log := handler.Logger{
		Logger:     logging.NewScopeLogger(baseLog, "workspace"),
		ServerMode: mode,
	}

	h := handler.New(Name, Version, log, mode)
	return &LSP{
		log:     log,
		handler: h,
		server:  server.NewServer(h.Handler(), Name, false),
	}, nil
}

// Run will start the server
func (s *LSP) Run() error {
	s.log.Info("Run server Stdio")

	return s.server.RunStdio()
}
