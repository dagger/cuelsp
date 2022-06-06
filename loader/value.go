package loader

import (
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

	return strings.HasPrefix(p, "#") && !strings.Contains(p, ".")
}

// ListDefinitions recursively walk through cue Value to retrieve all
// definition defines in it.
func (v *Value) ListDefinitions() ([]*Value, error) {
	var defs []*Value

	opts := []cue.Option{
		cue.Definitions(true),
	}

	customWalk(v, opts, func(v *Value) bool {
		if v.IsDefinition() {
			defs = append(defs, v)
		}

		return true
	}, nil)

	return defs, nil
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
