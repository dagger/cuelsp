package parser

import (
	"fmt"
	"sort"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
)

// map[int][]Range
// int: line number
// []Range: ascending by start.column() pos
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

// Due to AST structure, it will always be sorted in ascending order by start.Column()
func (def Definitions) AppendRange(name string, start token.Pos, end token.Pos) {
	def[start.Line()] = append(def[start.Line()], Range{start, end, name})
}

// Complexity O(log(n))
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

// Parse definitions of a given file
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
