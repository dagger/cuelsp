package lsp

func boolPtr(v bool) *bool {
	b := v
	return &b
}

func isTrue(v *bool) bool {
	return v != nil && *v == true
}
