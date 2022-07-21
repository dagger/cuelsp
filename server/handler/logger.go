package handler

import "github.com/tliron/kutil/logging"

type ServerMode interface {
	IsProd() bool
}

type Logger struct {
	logging.Logger
	ServerMode ServerMode
}
