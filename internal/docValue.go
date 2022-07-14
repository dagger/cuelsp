package internal

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
	"github.com/dagger/daggerlsp/loader"
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

	cleanDisplay := func(str string) string {
		if depth > 0 && !strings.Contains(str, "\t") {
			str = fmt.Sprintf("\t%s", str)
		}

		cleanTab := strings.ReplaceAll(str, "\t", "  ")
		cleanReturn := strings.ReplaceAll(cleanTab, "\n", "\n  ")
		return cleanReturn
	}

	formatNode := func(n ast.Node) string {
		display, err := format.Node(n, format.Simplify())
		if err == nil {
			return fmt.Sprintf("%s\n\n", cleanDisplay(string(display)))
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
		default:
			log.Println(reflect.TypeOf(v))
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
