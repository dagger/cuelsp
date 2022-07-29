package internal

import (
	"fmt"
	"log"
	"strings"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
	"github.com/dagger/cuelsp/loader"
)

type DocValue struct {
	description string
	structure   string
}

func NewDocValue(node ast.Node, v *loader.Value) *DocValue {
	doc := &DocValue{}

	for _, d := range v.Doc() {
		if d.Text() != "" {
			doc.description += fmt.Sprintf("%s\n", d.Text())
		}
	}

	field, ok := node.(*ast.Field)
	if ok {
		doc.structure = customerFormatNode(field, 0)
	}

	return doc
}

func customerFormatNode(node ast.Node, depth int) string {
	var doc string

	formatNode := func(n ast.Node) string {
		display, err := format.Node(n, format.Simplify())
		if err == nil {
			return string(display)
		}
		return "unknown"
	}

	switch n := node.(type) {
	case *ast.Field:
		switch v := n.Value.(type) {
		case *ast.Ident:
			if depth == 0 {
				return fmt.Sprintf("%s: %s", n.Label, v)
			}
			return formatNode(n)

		case *ast.UnaryExpr, *ast.BinaryExpr:
			return formatNode(n)
		case *ast.StructLit:
			return fmt.Sprintf("%s: %s", n.Label, formatNode(v))
		case *ast.BasicLit:
			return fmt.Sprintf("%s: %s", n.Label, formatNode(v))
		default:
			doc += fmt.Sprintf("%s: {\n%s}", n.Label, customerFormatNode(v, depth+1))
		}
	case *ast.StructLit:
		for _, d := range n.Elts {
			doc += customerFormatNode(d, depth+1)
		}
	}

	return doc
}

func (d *DocValue) String() string {
	var doc string

	if d.description != "" {
		doc = fmt.Sprintf("#### Description\n%s", d.description)
	}

	if d.structure != "" {
		doc += fmt.Sprintf("#### Type\n%s", d.structure)
	}

	return doc
}

func (d *DocValue) MarkdownString() string {
	var doc string

	if d.description != "" {
		doc = fmt.Sprintf("#### Description\n%s", d.description)
	}

	if d.structure != "" {
		doc += "#### Type\n"

		structure := d.structure

		// Insert carrier return if it's a definition
		structure = strings.Replace(structure, ": {", ": {\n\n", 1)

		// Insert tab on each fields in a definition
		if strings.Index(structure, ": {") < strings.Index(structure, "\n") {
			lines := strings.Split(structure, "\n")
			for i := 1; i < len(lines)-1; i++ {
				lines[i] = fmt.Sprintf("\t%s", lines[i])
			}
			structure = strings.Join(lines, "\n")
		} else {
			log.Println("not struct")
		}

		doc += structure
	}

	return doc
}
