package handler

import (
	"fmt"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

// documentDidSave reload the plan after it has been saved to populate the
// latest changes in workspace.
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didSave
func (h *Handler) documentDidSave(context *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	h.log.Debugf("Document saved")
	h.log.Debugf("params: %#v", params)

	go context.Call(protocol.MethodWorkspaceSemanticTokensRefresh, nil, nil)

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
