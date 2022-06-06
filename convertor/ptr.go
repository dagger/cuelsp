package util

// BoolPtr convert a boolean to a pointer to boolean
func BoolPtr(v bool) *bool {
	b := v
	return &b
}
