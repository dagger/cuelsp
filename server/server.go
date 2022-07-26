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

	mode Mode

	debug bool
}

const (
	Name    = "daggerlsp"
	Version = "0.3.3"
)

// New initializes a new language protocol server that contains his logger
// and his handler
func New(opts ...Options) (*LSP, error) {
	lsp := &LSP{}

	for _, opt := range opts {
		opt(lsp)
	}

	baseLog := logging.GetLogger(Name)
	log := handler.Logger{
		Logger:     logging.NewScopeLogger(baseLog, "workspace"),
		ServerMode: lsp.mode,
	}

	lsp.log = log

	lsp.handler = handler.New(
		handler.WithName(Name),
		handler.WithVersion(Version),
		handler.WithLogger(log),
	)

	lsp.server = server.NewServer(lsp.handler.Handler(), Name, lsp.debug)

	return lsp, nil
}

// Run will start the server
func (s *LSP) Run() error {
	s.log.Info("Run server Stdio")

	return s.server.RunStdio()
}
