package handler

import (
	"github.com/dagger/dlsp/server/utils"
	"github.com/dagger/dlsp/workspace"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) initialize(_ *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	change := protocol.TextDocumentSyncKindFull

	capabilities := h.handler.CreateServerCapabilities()
	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		OpenClose: utils.BoolPtr(true),
		Change:    &change,
		Save:      utils.BoolPtr(true),
	}
	capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{
		WorkspaceFolders: &protocol.WorkspaceFoldersServerCapabilities{
			Supported:           utils.BoolPtr(true),
			ChangeNotifications: &protocol.BoolOrString{Value: utils.BoolPtr(true)},
		}}
	capabilities.DefinitionProvider = true

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	switch len(params.WorkspaceFolders) {
	case 0:
		h.log.Errorf("No workspace folder found")
	case 1:
		_uri, err := uri.Parse(params.WorkspaceFolders[0].URI)
		if err != nil {
			return nil, err
		}
		h.workspace = workspace.New(_uri.Filename(), h.log)
	default:
		h.log.Errorf("Multiple workspace not supported")
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    h.lsName,
			Version: &h.lsVersion,
		},
	}, nil
}
