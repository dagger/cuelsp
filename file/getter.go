package file

import (
	"cuelang.org/go/cue/ast"
	"github.com/dagger/dlsp/parser"
)

func (f *File) Path() string {
	return f.path
}

func (f *File) Content() *ast.File {
	return f.content
}

func (f *File) Defs() *parser.Definitions {
	return f.defs
}
