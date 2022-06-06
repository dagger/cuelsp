package plan

import (
	"fmt"

	loader2 "github.com/dagger/dlsp/loader"
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

func (p *Plan) Value() *loader2.Value {
	return p.v
}

func (p *Plan) Instance() *loader2.Instance {
	return p.instance
}

func (p *Plan) String() string {
	var imports string
	for _, i := range p.imports {
		imports += fmt.Sprintf("\n- %s", i)
	}

	return fmt.Sprintf("Root: %s, File: %s, Type: %s, Value: %s\n%s\n Imports: %s\n", p.root, p.file, p.kind, p.v, p.instance, imports)
}
