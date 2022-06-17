package parser

import (
	"fmt"
	"sort"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
)

// Definitions is a map with the line number of a CUE file as key
// and a range of definition as value.
// Those Range are sorted in ascending order by start.Column()
// Definitions can hold 10 definition (so 10 Range) in a line.
type Definitions map[int][]Range

func (def Definitions) String() string {
	str := fmt.Sprintf("%s\n", "")
	for line, r := range def {
		str += fmt.Sprintf("Line %d:", line)
		for _, pos := range r {
			str += fmt.Sprintf("\t%s\n", pos)
		}
	}
	return str
}

// AppendRange add a new definition to the line.
// Due to AST structure, it will always be sorted in ascending order by
// start.Column()
func (def Definitions) AppendRange(name string, start token.Pos, end token.Pos) {
	def[start.Line()] = append(def[start.Line()], Range{start, end, name})
}

// Find will search for a definition in the Definition object following line
// and column
// It will return the definition's name if found, or an error if not found
// Find function has a complexity of O(log(n)) thanks Definitions data
// structure that his a map.
func (def Definitions) Find(line int, column int) (string, error) {
	if r, ok := def[line]; ok {
		rLen := len(r)
		i := sort.Search(rLen, func(i int) bool { return column <= r[i].End().Column() })
		if i < rLen && r[i].Start().Column() <= column {
			return r[i].Name(), nil
		}
	}
	return "", fmt.Errorf("definition not found")
}

// ParseDefs will fill Definitions data structure will all definitions found
// in the given ast.File
// There are two type of definitions declared in the AST
// - Ident: those are definitions from the package itself
// - SelectorExpr: those are definitions from external package
// they will be stored as <pkg>.<def> (E.g., foo.#Bar)
func ParseDefs(defs *Definitions, f *ast.File) {
	ast.Walk(f, func(node ast.Node) bool {
		switch v := node.(type) {
		// case: #Def
		case *ast.Ident:
			if IsDefinition(v.Name) {
				defs.AppendRange(v.Name, v.Pos(), v.End())
			}

		// case: pkg.#Def
		case *ast.SelectorExpr:
			labelName, _, _ := ast.LabelName(v.Sel)
			if IsDefinition(labelName) {
				pkg, ok := v.X.(*ast.Ident)
				if !ok {
					return false
				}
				definitionName := fmt.Sprintf("%s.%s", pkg.Name, labelName)
				defs.AppendRange(definitionName, pkg.Pos(), v.Sel.End())
				return false
			}
		}
		return true
	}, nil)
}
