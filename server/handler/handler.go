package handler

import (
	"github.com/dagger/dlsp/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

type Handler struct {
	workspace *workspace.Workspace

	handler *protocol.Handler

	log logging.Logger

	lsName string

	lsVersion string
}

func New(lsName, lsVersion string, log logging.Logger) *Handler {
	h := &Handler{
		lsName:    lsName,
		lsVersion: lsVersion,
		log:       logging.NewScopeLogger(log, "workspace"),
	}

	h.handler = &protocol.Handler{
		Initialize:             h.initialize,
		Initialized:            h.initialized,
		Shutdown:               h.shutdown,
		SetTrace:               h.setTrace,
		TextDocumentDidSave:    h.documentDidSave,
		TextDocumentDidOpen:    h.documentDidOpen,
		TextDocumentDefinition: h.documentDefinition,
	}

	return h
}

func (h *Handler) Handler() *protocol.Handler {
	return h.handler
}
