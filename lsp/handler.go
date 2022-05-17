package lsp

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var Handler protocol.Handler

func init() {
	// General Messages
	Handler.Initialize = Initialize
	Handler.Initialized = Initialized
	Handler.Shutdown = Shutdown
	Handler.LogTrace = LogTrace
	Handler.SetTrace = SetTrace

	// Workspace
	// Handler.WorkspaceDidRenameFiles = WorkspaceDidRenameFiles

	// // Text Document Synchronization
	// Handler.TextDocumentDidOpen = TextDocumentDidOpen
	// Handler.TextDocumentDidChange = TextDocumentDidChange
	// Handler.WorkspaceDidChangeConfiguration = WorkspaceDidChangeConfiguration
	// Handler.TextDocumentDidSave = TextDocumentDidSave
	// Handler.TextDocumentDidClose = TextDocumentDidClose

	// // Language Features
	// Handler.TextDocumentCompletion = TextDocumentCompletion
	// Handler.TextDocumentDocumentSymbol = TextDocumentDocumentSymbol
}
