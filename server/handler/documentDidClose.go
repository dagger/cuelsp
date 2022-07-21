package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) documentDidClose(_ *glsp.Context, _ *protocol.DidCloseTextDocumentParams) error {
	return nil
}
