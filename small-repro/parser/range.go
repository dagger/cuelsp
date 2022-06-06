package parser

import (
	"fmt"

	"cuelang.org/go/cue/token"
)

// Range position of definitions
type Range struct {
	start token.Pos
	end   token.Pos
	name  string
}

func (r Range) String() string {
	start := r.start.Position()
	end := r.end.Position()
	return fmt.Sprintf("%s: s[%d:%d]|e[%d:%d]", r.name, start.Line, start.Column, end.Line, end.Column)
}
