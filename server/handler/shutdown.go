package handler

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (h *Handler) shutdown(_ *glsp.Context) error {
	h.log.Debugf("Shutdown")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}
