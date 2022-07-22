package handler

// wrapError is a utility function to manage any error returned by a handler
// depending on a context
// For now, it ignores error on production mode to not annoy users
func (h *Handler) wrapError(err error) error {
	if h.log.ServerMode.IsProd() {
		return nil
	}

	return err
}
