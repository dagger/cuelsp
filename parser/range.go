package parser

import (
	"fmt"

	"cuelang.org/go/cue/token"
)

// Range position of a definition
type Range struct {
	start token.Pos
	end   token.Pos
	name  string
}

func (r Range) Start() token.Pos {
	return r.start
}

func (r Range) End() token.Pos {
	return r.end
}

func (r Range) Name() string {
	return r.name
}

func (r Range) String() string {
	start := r.start.Position()
	end := r.end.Position()
	return fmt.Sprintf("%s: s[%d:%d]|e[%d:%d]", r.name, start.Line, start.Column, end.Line, end.Column)
}

