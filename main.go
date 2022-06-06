package main

import (
	"github.com/dagger/dlsp/convertor"
	"github.com/dagger/dlsp/workspace"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"
	"go.lsp.dev/uri"

	// Must include a backend implementation. See kutil's logging/ for other options.
	_ "github.com/tliron/kutil/logging/simple"
)

const lsName = "dlsp"

var (
	version string = "0.0.1"
	handler protocol.Handler
	wk      *workspace.Workspace
	log     = logging.GetLogger(lsName)
)

func main() {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/dlsp.log"
	// logging.Configure(2, &logTo)
	logging.Configure(2, nil)

	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentDidSave:    documentDidSave,
		TextDocumentDidOpen:    documentDidOpen,
		TextDocumentDefinition: documentDefinition,
	}

	serv := server.NewServer(&handler, lsName, false)

	log.Errorf("Run server Stdio")
	if err := serv.RunStdio(); err != nil {
		panic(err)
	}
}

func initialize(_ *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
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
	capabilities.DefinitionProvider = true

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	switch len(params.WorkspaceFolders) {
	case 0:
		log.Errorf("No workspace folder found")
	case 1:
		_uri, err := uri.Parse(params.WorkspaceFolders[0].URI)
		if err != nil {
			return nil, err
		}
		wk = workspace.New(_uri.Filename())
	default:
		log.Errorf("Multiple workspace not suported")
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(_ *glsp.Context, params *protocol.InitializedParams) error {
	log.Debugf("Initialized")
	log.Debugf("params: %#v", params)

	return nil
}

func shutdown(_ *glsp.Context) error {
	log.Debugf("Shutdown")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(_ *glsp.Context, params *protocol.SetTraceParams) error {
	log.Debugf("Set strace")
	protocol.SetTraceValue(params.Value)
	return nil
}

func documentDidSave(_ *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	log.Debugf("Document saved")
	log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return err
	}

	p, err := wk.GetPlan(_uri.Filename())
	if err != nil {
		return err
	}

	if err := p.Reload(); err != nil {
		return err
	}

	return nil
}

func documentDidOpen(_ *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	log.Debugf("Document opened: %s", params.TextDocument.URI)
	log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return err
	}

	if err := wk.AddPlan(_uri.Filename()); err != nil {
		return err
	}

	return nil
}

func documentDefinition(_ *glsp.Context, params *protocol.DefinitionParams) (interface{}, error) {
	log.Debugf("Jump to def: %s", params.TextDocument.URI)
	log.Debugf("params: %#v", params)

	// Look up for definition position
	// Look up for cue value from definition
	return nil, nil
}
