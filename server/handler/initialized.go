package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// initialized is called after language server client has been successfully
// initialized.
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialized
func (h *Handler) initialized(_ *glsp.Context, params *protocol.InitializedParams) error {
	h.log.Debugf("Initialized")
	h.log.Debugf("params: %#v", params)

	return nil
}
