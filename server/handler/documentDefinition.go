package handler

import (
	"fmt"

	"github.com/dagger/daggerlsp/server/utils"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

// documentDefinition returns the protocol.Location of searched value in the
// CUE plan.
// It currently handles CUE definition
// If nothing is found, it will return nil and an error
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition
func (h *Handler) documentDefinition(_ *glsp.Context, params *protocol.DefinitionParams) (interface{}, error) {
	h.log.Debugf("Jump to def: %s", params.TextDocument.URI)
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
	location, err := p.GetDefinition(
		h.workspace.TrimRootPath(_uri.Filename()),
		utils.UIntToInt(params.Position.Line),
		utils.UIntToInt(params.Position.Character),
	)
	if err != nil {
		return nil, h.wrapError(err)
	}

	h.log.Debugf("Position: %#v", location.Pos().Position())

	res := utils.CueLocationToLSPLocation(location)

	h.log.Debugf("Res: %#v", res)

	return res, nil
}
