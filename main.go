package main

import (
	"github.com/dagger/dlsp/server"

	_ "github.com/tliron/kutil/logging/simple"
)

func main() {
	s := server.New()

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
