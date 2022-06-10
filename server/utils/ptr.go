// Package utils is a simple package that group any pure functions useful
// in LSP
// /!\: functions in package utils should be the simplest possible
package utils

// BoolPtr convert a boolean to a pointer to boolean
func BoolPtr(v bool) *bool {
	b := v
	return &b
}
