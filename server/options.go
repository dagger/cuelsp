package server

import "github.com/tliron/kutil/logging"

type Options = func(*LSP)

func WithMode(mode Mode) Options {
	return func(lsp *LSP) {
		lsp.mode = mode

		switch mode {
		case ModeDev:
			logging.Configure(2, nil)
		case ModeProd:
			logging.Configure(0, nil)
		default:
			logging.Configure(1, nil)
		}
	}
}

func WithDebug(debug bool) Options {
	return func(lsp *LSP) {
		lsp.debug = debug
	}
}
