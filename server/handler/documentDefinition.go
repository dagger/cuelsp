package handler

import (
	"fmt"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) documentDefinition(_ *glsp.Context, params *protocol.DefinitionParams) (interface{}, error) {
	h.log.Debugf("Jump to def: %s", params.TextDocument.URI)
	h.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}

	p := h.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return nil, fmt.Errorf("plan not found")
	}

	h.log.Debugf("Pos {%x, %x}", params.Position.Line, params.Position.Character)
	h.log.Debugf("Find plan of %s", _uri.Filename())
	location, err := p.GetDefinition(
		h.workspace.TrimRootPath(_uri.Filename()),
		int(params.Position.Line)+1,
		int(params.Position.Character)+1,
	)
	if err != nil {
		return nil, err
	}

	h.log.Debugf("Position: %#v", location.Pos().Position())

	res := protocol.Location{
		URI: string(uri.File(location.Pos().Filename())),
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      protocol.UInteger(location.Pos().Line()) - 1,
				Character: protocol.UInteger(location.Pos().Column()) - 1,
			},
			End: protocol.Position{
				Line:      protocol.UInteger(location.Pos().Line()) - 1,
				Character: protocol.UInteger(location.Pos().Column()) - 1,
			},
		},
	}

	h.log.Debugf("Res: %#v", res)

	return res, nil
}
