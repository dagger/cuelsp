package handler

type Options = func(*Handler)

// WithName set language server name
func WithName(name string) Options {
	return func(h *Handler) {
		h.lsName = name
	}
}

// WithVersion set language server version
func WithVersion(version string) Options {
	return func(h *Handler) {
		h.lsVersion = version
	}
}

// WithLogger set logger in handler
func WithLogger(log Logger) Options {
	return func(h *Handler) {
		h.log = log
	}
}
