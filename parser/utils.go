package parser

import (
	"io/ioutil"
	"strings"

	"cuelang.org/go/cue/ast"
	cueparser "cuelang.org/go/cue/parser"
)

// append to slice sorted
// func appendSorted(s []Range, r Range) []Range {
// 	i := sort.Search(len(s), func(i int) bool { return s[i].start.Column() >= r.start.Column() })
// 	return append(s[:i], append([]Range{r}, s[i:]...)...)
// }

// Append Range to Definition, sorted
// func (def Definitions) AppendRange(name string, start token.Pos, end token.Pos) {
// 	def[start.Line()] = appendSorted(def[start.Line()], Range{start, end, name})
// }

func IsDefinition(name string) bool {
	return strings.HasPrefix(name, "#")
}

func readFile(filename string) []byte {
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
	src := readFile(URI)
	f, err := cueparser.ParseFile("", src, options...)
	if err != nil {
		return nil, err
	}
	return f, nil
}
