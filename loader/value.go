package loader

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
)

// Value is a wrapper around cue.Value to implement additional utility
// methods
type Value struct {
	cue.Value
}

// IsDefinition return true if the current value is a definition.
// It checks it through path name.
func (v *Value) IsDefinition() bool {
	p := v.Path().String()

	return (strings.HasPrefix(p, "#") || strings.HasPrefix(p, "_#")) && !strings.Contains(p, ".")
}

// ListDefinitions recursively walk through cue Value to retrieve all
// definition defines in it.
func (v *Value) ListDefinitions() ([]*Value, error) {
	var defs []*Value

	opts := []cue.Option{
		cue.Definitions(true),
		cue.Hidden(true),
	}

	customWalk(v, opts, func(v *Value) bool {
		if v.IsDefinition() {
			defs = append(defs, v)
		}

		// Ignore list because there cannot be definition
		// define in it.
		return v.Kind() != cue.ListKind
	}, nil)

	return defs, nil
}

// ListFieldDoc builds a text that contains all fields of a CUE value with
// their associated documentation.
// It also walks through children until find primitive dagger task.
//
// TODO (TomChv): Implement special handling for primitive type like
// `dagger.#FS`
func (v *Value) ListFieldDoc() (string, error) {
	var fieldDoc string

	opts := []cue.Option{
		cue.Optional(true),
	}

	customWalk(v, opts, func(v *Value) bool {
		var fieldType string
		var doc string

		path := strings.Split(v.Path().String(), ".")
		field := path[len(path)-1]
		indent := strings.Repeat("\t", len(path)-1)

		// Ignore original path
		if field == v.Path().String() {
			return true
		}

		// Do not explore dagger task
		if field == "$dagger" {
			return false
		}

		// Special type management
		// TODO(TomChv): Talk with Joel to enhance it with more details
		switch v.IncompleteKind() {
		case cue.StructKind:
			if len(path) == 1 {
				fieldType = "{"
			} else {
				fieldType = v.IncompleteKind().String()
			}
		case cue.ListKind:
			fieldType = v.IncompleteKind().String()
		default:
			fieldType = v.IncompleteKind().String()
		}

		// Aggregate docs in only one variable
		for _, d := range v.Doc() {
			doc += d.Text()
		}

		// Add documentation field and prettify it
		if doc != "" {
			fieldDoc += fmt.Sprintf("%s// %s", indent, strings.Replace(
				doc,
				"\n",
				fmt.Sprintf("\n%s// ", indent), strings.Count(doc, "\n")-1))
		}

		// Concat documentation
		fieldDoc += fmt.Sprintf("%s%s: %s  \n", indent, field, fieldType)

		return true
	}, func(v *Value) {
		path := strings.Split(v.Path().String(), ".")

		// Add close bracket for structure
		if v.IncompleteKind() == cue.StructKind && len(path) == 2 {
			fieldDoc += fmt.Sprintf("%s}  \n", strings.Repeat("\t", len(path)-1))
		}
	})

	return fieldDoc, nil
}

// customWalk is an alternative to cue.Value.Walk that enable options
// configuration to retrieve any kind of data in the cue value.
func customWalk(v *Value, opts []cue.Option, before func(v *Value) bool, after func(v *Value)) {
	// call before and possibly stop recursion
	if before != nil && !before(v) {
		return
	}

	// possibly recurse
	switch v.IncompleteKind() {
	case cue.StructKind:
		s, _ := v.Fields(opts...)

		for s.Next() {
			customWalk(&Value{s.Value()}, opts, before, after)
		}

	case cue.ListKind:
		l, _ := v.List()
		for l.Next() {
			customWalk(&Value{l.Value()}, opts, before, after)
		}
	}

	if after != nil {
		after(v)
	}
}
