package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) setTrace(_ *glsp.Context, params *protocol.SetTraceParams) error {
	h.log.Debugf("Set strace")
	protocol.SetTraceValue(params.Value)
	return nil
}
