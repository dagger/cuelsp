package parser

import (
	"fmt"

	"cuelang.org/go/cue/ast"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

// Definitions is a map with the line number of a CUE file as key
// and a range of definition as value.
// Those Range are sorted in ascending order by start.Column()
// Definitions can hold 10 definition (so 10 Range) in a line.
// type Tokens map[int][]Range
type Scanner struct {
	// startLine uint32
	// startCol  uint32
	tokens []protocol.UInteger
	logger logging.Logger
}

func (s *Scanner) String() string {
	// start := r.start.Position()
	// end := r.end.Position()
	// return fmt.Sprintf("%s: s[%d:%d]|e[%d:%d]", r.name, start.Line, start.Column, end.Line, end.Column)
	return fmt.Sprintf("deb:[%v]", s.tokens)
}

// AppendRange add a new definition to the line.
// Due to AST structure, it will always be sorted in ascending order by
// start.Column()
// func (s *Scanner) append(diffLine uint32, diffCol uint32, node *sitter.Node, patternIndex uint16) {
// 	s.tokens = append(
// 		s.tokens,
// 		diffLine,
// 		diffCol,
// 		node.EndPoint().Column-node.StartPoint().Column,
// 		semantic.TokenTypeIndex(patternIndex),
// 		semantic.TokenModifierIndex(patternIndex),
// 	)
// }

func ParseSemanticTokens(tokens *Scanner, f *ast.File) {
	ast.Walk(f, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.Comment:
			comment := Range{
				start: n.Pos(),
				end:   n.End(),
				name:  n.Text,
			}
			fmt.Println("comment", comment, "pos", n.Pos(), n.End())

		case *ast.CommentGroup:
			// fmt.Println("attribute", attribute, n.At, "pos", n.Pos(), n.End())
		// 	for _, c := range n.List {
		// 		walk(v, c)
		// 	}

		case *ast.Attribute:
			attribute := Range{
				start: n.Pos(),
				end:   n.End(),
				name:  n.Text,
			}
			fmt.Println("attribute", attribute, n.At, "pos", n.Pos(), n.End())

		case *ast.Field: // not important
			// field := Range{
			// 	start: n.Pos(),
			// 	end:   n.End(),
			// }

			fmt.Println("Field", n.Optional, n.Token, "pos", n.Pos(), n.End())
		// walk(v, n.Label)
		// if n.Value != nil {
		// 	walk(v, n.Value)
		// }
		// for _, a := range n.Attrs {
		// 	walk(v, a)
		// }

		case *ast.StructLit:
			fmt.Println("structlit", n.Lbrace, n.Rbrace, "pos", n.Pos(), n.End())
			// walkDeclList(v, n.Elts)

			// Expressions
		case *ast.BottomLit:
			fmt.Println("bottomLit", "pos", n.Pos(), n.End(), n.Bottom)
		case *ast.BadExpr:
			fmt.Println("BadExpr", "pos", n.Pos(), n.End(), n.From, n.To)
		case *ast.Ident:
			fmt.Println("ident", "pos", n.Pos(), n.End(), n.Name, n.NamePos, n.String())

		case *ast.BasicLit:
			fmt.Println("BasicLit", "pos", n.Pos(), n.End(), n.Kind, n.Value, n.ValuePos)
			// nothing to do

		case *ast.Interpolation:
			fmt.Println("Interpolation", "pos", n.Pos(), n.End())
		// for _, e := range n.Elts {
		// 	walk(v, e)
		// }

		case *ast.ListLit:
			fmt.Println("ListLit", "[", n.Lbrack, "]", n.Rbrack, "pos", n.Pos(), n.End())
			// walkExprList(v, n.Elts)

		case *ast.Ellipsis:
			fmt.Println("Ellipsis", "ellipsis", n.Ellipsis, "pos:", "pos", n.Pos(), n.End())
		// if n.Type != nil {
		// 	walk(v, n.Type)
		// }

		case *ast.ParenExpr:
			fmt.Println("ParenExpr", "(", n.Lparen, "):", n.Rparen, "pos:", n.Pos(), n.End())
			// walk(v, n.X)

		case *ast.SelectorExpr:
			fmt.Println("SelectorExpr", "pos:", n.Pos(), n.End())
		// 	n.
		// walk(v, n.X)
		// walk(v, n.Sel)

		case *ast.IndexExpr:
			fmt.Println("IndexExpr", "[", n.Lbrack, "]", n.Rbrack, "pos:", n.Pos(), n.End())
			// walk(v, n.X)
			// walk(v, n.Index)

		case *ast.SliceExpr:
			fmt.Println("SliceExpr", "[", n.Lbrack, "]", n.Rbrack, "pos:", n.Pos(), n.End())
			// walk(v, n.X)
			// if n.Low != nil {
		// 	walk(v, n.Low)
		// }
		// if n.High != nil {
		// 	walk(v, n.High)
		// }

		case *ast.CallExpr:
			fmt.Println("CallExpr", "[", n.Lparen, "]", n.Rparen, "pos:", n.Pos(), n.End())
			// walk(v, n.Fun)
			// walkExprList(v, n.Args)

		case *ast.UnaryExpr:
			fmt.Println("UnaryExpr", "type", n.Op, "]", "pos:", n.Pos(), n.End())
			// walk(v, n.X)

		case *ast.BinaryExpr:
			fmt.Println("BinaryExpr")
			// 	walk(v, n.X)
			// walk(v, n.Y)

			// Declarations
		case *ast.ImportSpec:
			// n.Name
			fmt.Println("ImportSpec", "pos:", n.Pos(), n.End())
			// if n.Name != nil {
		// }
		// if n.Name != nil {
		// 	walk(v, n.Name)
		// }
		// walk(v, n.Path)

		case *ast.BadDecl:
			fmt.Println("BadDecl", "from", n.From, "to", n.To, "pos:", n.Pos(), n.End())
			// nothing to do

		case *ast.ImportDecl:
			fmt.Println("ImportDecl", "import", n.Import, "Lparen", n.Lparen, "Rparen", n.Rparen, "pos", n.Pos(), n.End())
		// for _, s := range n.Specs {
		// 	walk(v, s)
		// }

		case *ast.EmbedDecl:
			fmt.Println("EmbedDecl", "pos", n.Pos(), n.End())
			// walk(v, n.Expr)

		case *ast.LetClause:
			fmt.Println("LetClause", "let", n.Let, "equal", n.Equal, "pos", n.Pos(), n.End())
			// walk(v, n.Ident)
			// walk(v, n.Expr)

		case *ast.Alias:
			fmt.Println("EmbedDecl", "equal", n.Equal, "pos", n.Pos(), n.End())
			// walk(v, n.Ident)
			// walk(v, n.Expr)

		case *ast.Comprehension:
			fmt.Println("Comprehension", "pos", n.Pos(), n.End())
		// for _, c := range n.Clauses {
		// 	walk(v, c)
		// }
		// walk(v, n.Value)

		// Files and packages
		case *ast.File:
		// 	fmt.Println("File", "pos", n.Pos(), n.End())
		// Not important !!!!!
		// walkDeclList(v, n.Decls)

		case *ast.Package:
		// 	fmt.Println("Package", "pos", n.Pos(), n.End()) // [!!in probation!!]
		// walk(v, n.Name)

		case *ast.ForClause:
			fmt.Println("ForClause", "for", n.For, "in", n.In, "colon", n.Colon, "pos", n.Pos(), n.End())
		// if n.Key != nil {
		// 	walk(v, n.Key)
		// }
		// walk(v, n.Value)
		// walk(v, n.Source)

		case *ast.IfClause:
			fmt.Println("IfClause", "if", n.If, "pos", n.Pos(), n.End())
			// walk(v, n.Condition)

		default:
			panic(fmt.Sprintf("Walk: unexpected node type %T", n))
		}

		// case: #Def
		// case *ast.Ident:
		// 	if IsDefinition(v.Name) {
		// 		defs.AppendRange(v.Name, v.Pos(), v.End())
		// 	}

		// // case: pkg.#Def
		// case *ast.SelectorExpr:
		// 	labelName, _, _ := ast.LabelName(v.Sel)
		// 	if IsDefinition(labelName) {
		// 		pkg, ok := v.X.(*ast.Ident)
		// 		if !ok {
		// 			return false
		// 		}
		// 		definitionName := fmt.Sprintf("%s.%s", pkg.Name, labelName)
		// 		defs.AppendRange(definitionName, pkg.Pos(), v.Sel.End())
		// 		return false
		// 	}
		// }
		return true
	}, nil)
}
