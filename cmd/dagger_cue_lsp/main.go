package main

import (
	"github.com/dagger/daggerlsp/server"

	_ "github.com/tliron/kutil/logging/simple"
)

func main() {
	s, err := server.New(server.ModeDev)
	if err != nil {
		panic(err)
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
