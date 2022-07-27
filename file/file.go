package file

import (
	"fmt"

	"cuelang.org/go/cue/ast"
	cueparser "cuelang.org/go/cue/parser"
	"github.com/dagger/cuelsp/parser"
)

// File is abstraction of a raw cue file
// It useful to statically analyze a file without validating the CUE value.
type File struct {
	// Path of the file
	path string

	// AST file content
	content *ast.File

	// Definitions of the file
	defs *parser.Definitions
}

// New create a File and analise CUE ast in it.
func New(path string) (*File, error) {
	content, err := cueparser.ParseFile(path, nil)
	if err != nil {
		return nil, err
	}

	defs := parser.Definitions{}
	parser.ParseDefs(&defs, content)

	return &File{
		path:    path,
		content: content,
		defs:    &defs,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("%s,%s", f.path, f.defs)
}
