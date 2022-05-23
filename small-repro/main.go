package main

import (
	util "github.com/dagger/dlsp/small-repro/convertor"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"

	// Must include a backend implementation. See kutil's logging/ for other options.
	_ "github.com/tliron/kutil/logging/simple"
)

const lsName = "dlsp"

var version string = "0.0.1"
var handler protocol.Handler

var log = logging.GetLogger(lsName)

func main() {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/dlsp.log"
	// logging.Configure(2, &logTo)
	logging.Configure(2, nil)

	handler = protocol.Handler{
		Initialize:            initialize,
		Initialized:           initialized,
		Shutdown:              shutdown,
		SetTrace:              setTrace,
		TextDocumentDidChange: documentDidChange,
		TextDocumentDidOpen:   documentDidOpen,
	}

	server := server.NewServer(&handler, lsName, false)

	log.Errorf("Run server Stdio")
	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	capabilities := handler.CreateServerCapabilities()
	change := protocol.TextDocumentSyncKindFull
	capabilities.TextDocumentSync = protocol.TextDocumentSyncOptions{
		OpenClose: util.BoolPtr(true),
		Change:    &change,
		Save:      util.BoolPtr(true),
	}
	capabilities.Workspace = &protocol.ServerCapabilitiesWorkspace{
		WorkspaceFolders: &protocol.WorkspaceFoldersServerCapabilities{
			Supported:           util.BoolPtr(true),
			ChangeNotifications: &protocol.BoolOrString{Value: util.BoolPtr(true)},
		}}

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	log.Infof("rootPath: %#v", params.WorkspaceFolders)

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	log.Errorf("Initialized")
	log.Errorf("params: %#v", params)

	var result interface{}
	context.Call("workspace/workspaceFolders", "bar", &result)
	log.Infof("Result: %#v", result)
	return nil
}

func shutdown(context *glsp.Context) error {
	log.Errorf("Shutdown")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	log.Errorf("Set strace")
	protocol.SetTraceValue(params.Value)
	return nil
}

func documentDidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	log.Infof("Document changed")
	log.Infof("params: %#v", params)
	return nil
}

func documentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	log.Infof("Document opened")
	log.Infof("params: %#v", params)
	return nil
}
