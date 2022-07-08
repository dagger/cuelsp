package handler

import (
	"fmt"

	"github.com/dagger/daggerlsp/server/utils"
	"github.com/dagger/daggerlsp/workspace"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

// initialize the language server with his capabilities and the user's workspace.
// Spec: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
// /!\ Only one workspace in currently supported.
func (h *Handler) initialize(_ *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	capabilities := h.capabilities()

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	if err := h.initWorkspace(params.WorkspaceFolders, params.RootURI, params.RootPath); err != nil {
		return nil, err
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    h.lsName,
			Version: &h.lsVersion,
		},
	}, nil
}

// initWorkspace creates a new workspace depending on workspace folders.
// Currently, it does not handle multiple workspace
func (h *Handler) initWorkspace(workspaceFolders []protocol.WorkspaceFolder, rootURI, rootPath *string) error {
	switch len(workspaceFolders) {
	case 0:
		var path string
		switch {
		case rootURI != nil:
			path = *rootURI
		case rootPath != nil:
			path = *rootPath
		default:
			return fmt.Errorf("no workspace folder found")
		}

		_uri, err := uri.Parse(path)
		if err != nil {
			return err
		}
		h.workspace = workspace.New(_uri.Filename(), h.log)

		return nil
	case 1:
		_uri, err := uri.Parse(workspaceFolders[0].URI)
		if err != nil {
			return err
		}
		h.workspace = workspace.New(_uri.Filename(), h.log)

		return nil
	default:
		return fmt.Errorf("multiple workspace not supported")
	}
}

// capabilities return set of Handler server capabilities
func (h *Handler) capabilities() protocol.ServerCapabilities {
	capabilities := h.handler.CreateServerCapabilities()

	// Synchronisation
	change := protocol.TextDocumentSyncKindFull
	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		OpenClose: utils.BoolPtr(true),
		Change:    &change,
		Save:      utils.BoolPtr(true),
	}

	// Workspace configuration
	capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{
		WorkspaceFolders: &protocol.WorkspaceFoldersServerCapabilities{
			Supported:           utils.BoolPtr(true),
			ChangeNotifications: &protocol.BoolOrString{Value: utils.BoolPtr(true)},
		}}

	// Jump to definition
	capabilities.DefinitionProvider = true
	capabilities.HoverProvider = true
	return capabilities
}
