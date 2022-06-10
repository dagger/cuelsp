package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// shutdown is a method to stop the server handle request but not stop itself
// If any other request than exist is called, it should return an error
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#shutdown
func (h *Handler) shutdown(_ *glsp.Context) error {
	h.log.Debugf("Shutdown")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}
