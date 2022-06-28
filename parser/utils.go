package parser

import (
	"io/ioutil"
	"strings"

	"cuelang.org/go/cue/ast"
	cueparser "cuelang.org/go/cue/parser"
)

// IsDefinition returns true if the current name is a CUE definition
// Pattern detected are:
// - #Foo
// - _#Foo
// It returns false if it's not a definition
func IsDefinition(name string) bool {
	return strings.HasPrefix(name, "#") || strings.HasPrefix(name, "_#")
}

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

// parse URI to cueparse.AST
func ParseFile(URI string) (*ast.File, error) {
	// 	// https://pkg.go.dev/cuelang.org/go@v0.4.3/cue/parser#ParseFile
	// options := []parser.Option{parser.AllErrors, parser.ParseComments, parser.FromVersion(parser.Latest)}
	options := []cueparser.Option{cueparser.AllErrors, cueparser.FromVersion(cueparser.Latest)}
	src := ReadFile(URI)
	f, err := cueparser.ParseFile("", src, options...)
	if err != nil {
		return nil, err
	}
	return f, nil
}
