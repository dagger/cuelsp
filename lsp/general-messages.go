package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	// "github.com/tliron/kutil/version"
)

// lsp-sample implementation
var hasConfigurationCapability bool = false
var hasWorkspaceFolderCapability bool = false
var hasDiagnosticRelatedInformationCapability bool = false

var clientCapabilities *protocol.ClientCapabilities

var Version = "dev"

// protocol.InitializeFunc signature
// Returns: InitializeResult | InitializeError
func Initialize(context *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	clientCapabilities = &params.Capabilities

	// Does the client support the `workspace/configuration` request?
	// If not, we fall back using global settings.
	hasConfigurationCapability = isTrue(clientCapabilities.Workspace.Configuration)
	hasWorkspaceFolderCapability = isTrue(clientCapabilities.Workspace.WorkspaceFolders)
	hasDiagnosticRelatedInformationCapability = isTrue(clientCapabilities.TextDocument.PublishDiagnostics.RelatedInformation)

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	serverCapabilities := Handler.CreateServerCapabilities()
	serverCapabilities.TextDocumentSync = protocol.TextDocumentSyncKindIncremental
	serverCapabilities.CompletionProvider = &protocol.CompletionOptions{
		ResolveProvider: boolPtr(true),
	}

	serverCapabilities.ReferencesProvider = &protocol.ReferenceOptions{}

	if hasWorkspaceFolderCapability {
		serverCapabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{
			WorkspaceFolders: &protocol.WorkspaceFoldersServerCapabilities{
				Supported: boolPtr(true),
			},
		}
	}

	return &protocol.InitializeResult{
		Capabilities: serverCapabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    "dagger-lsp", // make it a variable (later)
			Version: &Version,     // version string (to be filled in) ???
		},
	}, nil
}

// protocol.InitializedFunc signature
func Initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	// if hasConfigurationCapability {
	// 	Handler.WorkspaceDidChangeConfiguration = WorkspaceDidChangeConfiguration
	// }
	// if hasWorkspaceFolderCapability {
	// 	Handler.WorkspaceDidChangeWorkspaceFolders = WorkspaceDidChangeWorkspaceFolders
	// }
	return nil
}

// protocol.ShutdownFunc signature
func Shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	// resetDocumentStates()
	return nil
}

// protocol.LogTraceFunc signature
func LogTrace(context *glsp.Context, params *protocol.LogTraceParams) error {
	return nil
}

// protocol.SetTraceFunc signature
func SetTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

// TextDocumentDidChange

// WorkspaceDidChangeWorkspaceFolders
