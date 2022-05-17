package main

import (
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

func main() {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/dlsp.log"
	// logging.Configure(2, &logTo)
	logging.Configure(2, nil)

	handler = protocol.Handler{
		Initialize:  initialize,
		Initialized: initialized,
		Shutdown:    shutdown,
		SetTrace:    setTrace,
	}

	server := server.NewServer(&handler, lsName, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (interface{}, error) {
	capabilities := handler.CreateServerCapabilities()

	if params.Trace != nil {
		protocol.SetTraceValue(*params.Trace)
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
