package plan

import (
	"fmt"

	"cuelang.org/go/cue"
)

func (p *Plan) Root() string {
	return p.root
}

func (p *Plan) File() string {
	return p.file
}

func (p *Plan) Kind() Kind {
	return p.kind
}

func (p *Plan) Value() cue.Value {
	return p.v
}

func (p *Plan) String() string {
	return fmt.Sprintf("Root: %s, File: %s, Type: %s, Value: %s", p.root, p.file, p.kind, p.v)
}
