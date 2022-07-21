package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) documentDidChange(_ *glsp.Context, _ *protocol.DidChangeTextDocumentParams) error {
	return nil
}
