package handler

import (
	"fmt"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) documentDidSave(_ *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	h.log.Debugf("Document saved")
	h.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return err
	}

	p := h.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return fmt.Errorf("plan not found")
	}

	if err := p.Reload(); err != nil {
		return err
	}

	return nil
}
