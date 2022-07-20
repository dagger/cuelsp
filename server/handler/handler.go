// Package handler is a collection of all methods supported by LSP.
// To implements a new methods, just create a new file with the name of the
// method.
package handler

import (
	"github.com/dagger/daggerlsp/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

type Mode int

const (
	DEV Mode = iota
	PROD
)

// Handler is the storage for any handler of the server.LSP.
// It also handles a single workspace for now which represent a VSCode project
type Handler struct {
	workspace *workspace.Workspace

	handler *protocol.Handler

	log logging.Logger

	lsName string

	lsVersion string

	mode Mode
}

// New creates a Handler instance that contains all methods supported by
// the LSP
func New(lsName, lsVersion string, log logging.Logger, mode Mode) *Handler {
	h := &Handler{
		lsName:    lsName,
		lsVersion: lsVersion,
		log:       logging.NewScopeLogger(log, "workspace"),
		mode:      mode,
	}

	h.handler = &protocol.Handler{
		Initialize:                     h.initialize,
		Initialized:                    h.initialized,
		Shutdown:                       h.shutdown,
		SetTrace:                       h.setTrace,
		TextDocumentDidSave:            h.documentDidSave,
		TextDocumentDidOpen:            h.documentDidOpen,
		TextDocumentDefinition:         h.documentDefinition,
		TextDocumentHover:              h.documentHover,
		TextDocumentSemanticTokensFull: h.documentSemanticTokensFull,
	}

	return h
}

func (h *Handler) Handler() *protocol.Handler {
	return h.handler
}
