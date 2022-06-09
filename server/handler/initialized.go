package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) initialized(_ *glsp.Context, params *protocol.InitializedParams) error {
	h.log.Debugf("Initialized")
	h.log.Debugf("params: %#v", params)

	return nil
}
