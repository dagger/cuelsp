package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// setTrace is a simple method that language server client can call to set
// trace level on server
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#setTrace
func (h *Handler) setTrace(_ *glsp.Context, params *protocol.SetTraceParams) error {
	h.log.Debugf("Set strace")
	protocol.SetTraceValue(params.Value)
	return nil
}
