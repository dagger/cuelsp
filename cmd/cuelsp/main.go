package main

import (
	"github.com/dagger/cuelsp/server"

	_ "github.com/tliron/kutil/logging/simple"
)

func main() {
	s, err := server.New(
		server.WithMode(server.ModeDev),
		server.WithDebug(true),
	)
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
