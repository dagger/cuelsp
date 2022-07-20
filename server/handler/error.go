package handler

// wrapError is a utility function to manage any error returned by a handler
// depending on a context
// For now, it ignores error on production mode to do not annoying users
func (h *Handler) wrapError(err error) error {
	if h.mode == PROD {
		return nil
	}

	return err
}
