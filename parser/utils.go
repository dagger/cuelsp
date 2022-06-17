package parser

import (
	"strings"
)

// IsDefinition returns true if the current name is a CUE definition
// Pattern detected are:
// - #Foo
// - _#Foo
// It returns false if it's not a definition
func IsDefinition(name string) bool {
	return strings.HasPrefix(name, "#") || strings.HasPrefix(name, "_#")
}
