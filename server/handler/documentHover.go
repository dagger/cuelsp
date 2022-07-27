package handler

import (
	"fmt"

	"github.com/dagger/cuelsp/server/utils"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

// documentHover will return documentation for the hovered definition
// FIXME(TomChv): Handle keys instead of only definition
// FIXME(TomChv): Refactor function to avoid code duplication on "Get Definition logic)
func (h *Handler) documentHover(_ *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	h.log.Debugf("Hover from: %s", params.TextDocument.URI)
	h.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, h.wrapError(err)
	}

	p := h.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return nil, h.wrapError(fmt.Errorf("plan not found"))
	}

	h.log.Debugf("Pos {%x, %x}", params.Position.Line, params.Position.Character)
	h.log.Debugf("Find plan of %s", _uri.Filename())
	doc, err := p.GetDocDefinition(
		h.workspace.TrimRootPath(_uri.Filename()),
		utils.UIntToInt(params.Position.Line),
		utils.UIntToInt(params.Position.Character),
	)
	if err != nil {
		return nil, h.wrapError(err)
	}

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: doc.MarkdownString(),
		},
	}, nil
}
