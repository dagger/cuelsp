package handler

import (
	"fmt"

	"cuelang.org/go/cue/format"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) documentFormatting(_ *glsp.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	h.log.Debugf("Format: %s", params.TextDocument.URI)
	h.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, h.wrapError(err)
	}

	p := h.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return nil, h.wrapError(fmt.Errorf("plan not found"))
	}

	f := p.Files()[h.workspace.TrimRootPath(_uri.Filename())]
	h.log.Debugf("File %v", f)

	_fmt, err := format.Node(f.Content(), format.UseSpaces(2), format.TabIndent(false)) // TODO: gather from params.Options?
	if err != nil {
		return nil, h.wrapError(err)
	}

	h.log.Debugf("Source formatted: %s", _uri.Filename())
	start := protocol.Position{Line: 0, Character: 0}

	end := protocol.Position{Line: uint32(f.Content().End().Line()), Character: 0}

	edit := protocol.TextEdit{
		Range: protocol.Range{
			Start: start,
			End:   end,
		},
		NewText: string(_fmt),
	}

	return []protocol.TextEdit{edit}, nil
}
