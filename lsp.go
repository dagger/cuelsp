package main

import (
	"fmt"

	util "github.com/dagger/dlsp/convertor"
	"github.com/dagger/dlsp/workspace"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
	"github.com/tliron/kutil/logging"
	"go.lsp.dev/uri"
)

type LSPServer struct {
	workspace *workspace.Workspace

	handler      protocol.Handler
	capabilities protocol.ServerCapabilities

	server *server.Server

	log logging.Logger
}

func NewLSPServer(lsName string, log logging.Logger) *LSPServer {
	var handler protocol.Handler

	s := LSPServer{
		log: log,
	}

	handler = protocol.Handler{
		Initialize:             s.initialize,
		Initialized:            s.initialized,
		Shutdown:               s.shutdown,
		SetTrace:               s.setTrace,
		TextDocumentDidSave:    s.documentDidSave,
		TextDocumentDidOpen:    s.documentDidOpen,
		TextDocumentDefinition: s.documentDefinition,
	}

	// we complete the fields once they are correctly filled
	s.handler = handler
	s.server = server.NewServer(&handler, lsName, false)

	return &s
}

func (s *LSPServer) Run() error {
	s.log.Info("Run server Stdio")
	err := s.server.RunStdio()
	if err != nil {
		return err
	}
	return nil
}

func (s *LSPServer) initialize(_ *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	change := protocol.TextDocumentSyncKindFull

	capabilities := s.handler.CreateServerCapabilities()
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
		s.log.Errorf("No workspace folder found")
	case 1:
		_uri, err := uri.Parse(params.WorkspaceFolders[0].URI)
		if err != nil {
			return nil, err
		}
		s.workspace = workspace.New(_uri.Filename(), s.log)
	default:
		s.log.Errorf("Multiple workspace not suported")
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func (s *LSPServer) initialized(_ *glsp.Context, params *protocol.InitializedParams) error {
	s.log.Debugf("Initialized")
	s.log.Debugf("params: %#v", params)

	return nil
}

func (s *LSPServer) shutdown(_ *glsp.Context) error {
	s.log.Debugf("Shutdown")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func (s *LSPServer) setTrace(_ *glsp.Context, params *protocol.SetTraceParams) error {
	s.log.Debugf("Set strace")
	protocol.SetTraceValue(params.Value)
	return nil
}

func (s *LSPServer) documentDidSave(_ *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	s.log.Debugf("Document saved")
	s.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return err
	}

	p := s.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return fmt.Errorf("plan not found")
	}

	if err := p.Reload(); err != nil {
		return err
	}

	return nil
}

func (s *LSPServer) documentDidOpen(_ *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	s.log.Debugf("Document opened: %s", params.TextDocument.URI)
	s.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return err
	}

	if err := s.workspace.AddPlan(_uri.Filename()); err != nil {
		return err
	}

	return nil
}

func (s *LSPServer) documentDefinition(_ *glsp.Context, params *protocol.DefinitionParams) (interface{}, error) {
	s.log.Debugf("Jump to def: %s", params.TextDocument.URI)
	s.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}

	p := s.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return nil, fmt.Errorf("plan not found")
	}

	s.log.Debugf("Pos {%x, %x}", params.Position.Line, params.Position.Character)
	s.log.Debugf("Find plan of %s", _uri.Filename())
	location, err := p.GetDefinition(
		s.workspace.TrimRootPath(_uri.Filename()),
		int(params.Position.Line)+1,
		int(params.Position.Character)+1,
	)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("Position: %#v", location.Pos().Position())

	res := protocol.Location{
		URI: fmt.Sprintf("%s", uri.File(location.Pos().Filename())),
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

	s.log.Debugf("Res: %#v", res)

	return res, nil
}
