package util

func BoolPtr(v bool) *bool {
	b := v
	return &b
}
