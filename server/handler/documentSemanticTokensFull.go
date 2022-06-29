package handler

import (
	"github.com/dagger/dlsp/semantic"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

// documentDidOpen register a new plan in the workspace
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen
func (h *Handler) documentSemanticTokensFull(context *glsp.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	h.log.Debugf("Semantic Tokens: %s", params.TextDocument.URI)
	h.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}

	// TODO: remonter les erreurs
	//

	data, err := semantic.Tokenize(_uri.Filename(), h.log)
	// if err := h.workspace.AddPlan(_uri.Filename()); err != nil {
	// 	return err
	// }

	// ok := []protocol.UInteger{
	// 	0, 0, 7, 0, 0,
	// 	0, 8, 6, 1, 0,
	// }
	// return &protocol.SemanticTokens{
	// 	Data: ok,
	// }, nil
	// toto := protocol.TextDocumentSemanticTokensRefreshFunc(context)
	return &protocol.SemanticTokens{
		Data: data,
	}, nil
	// return nil, nil
}
