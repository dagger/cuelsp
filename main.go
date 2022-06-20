package main

import (
	"github.com/dagger/dlsp/server"

	_ "github.com/tliron/kutil/logging/simple"
)

func main() {
	s := server.New()

	if err := s.Run(); err != nil {
		panic(err)
	}
}
