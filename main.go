package main

import (
	// Must include a backend implementation. See kutil's logging/ for other options.
	"github.com/tliron/kutil/logging"
	_ "github.com/tliron/kutil/logging/simple"
)

const lsName = "dlsp"

var (
	version string = "0.0.1"
)

func main() {
	// This increases logging verbosity (optional)
	// logTo := "/tmp/dlsp.log"
	// logging.Configure(2, &logTo)
	logging.Configure(2, nil)

	log := logging.GetLogger(lsName)

	s := NewLSPServer(lsName, log)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
