package parser

import (
	"strings"
)

func IsDefinition(name string) bool {
	return strings.HasPrefix(name, "#") || strings.HasPrefix(name, "_#")
}
