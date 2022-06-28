package handler

import (
	"github.com/dagger/daggerlsp/semantic"
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

	// TODO: Hightlight errors
	data, _ := semantic.Tokenize(_uri.Filename(), h.log)

	return &protocol.SemanticTokens{
		Data: data,
	}, nil
}
