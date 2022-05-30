package plan

import (
	"cuelang.org/go/cue"
	"github.com/dagger/dlsp/small-repro/loader"
)

// Plan is a representation of a cue value in a workspace
type Plan struct {
	// Root path of the plan
	root string

	// File loaded
	file string

	// Plan's kind
	kind Kind

	// Cue Value
	v cue.Value
}

// New load a new cue value
func New(root, file string) (*Plan, error) {
	v, err := loader.Dir(root, file)
	k := Directory
	if err != nil {
		v2, err2 := loader.File(root, file)
		if err2 != nil {
			return nil, err
		}

		k = File
		v = v2
	}

	// Load cue value
	return &Plan{
		root: root,
		file: file,
		kind: k,
		v:    v,
	}, nil
}

// Reload will rebuild the cue value
func (p *Plan) Reload() error {
	var (
		v   cue.Value
		err error
	)

	switch p.kind {
	case File:
		v, err = loader.File(p.root, p.file)
	case Directory:
		v, err = loader.Dir(p.root, p.file)
	}

	if err != nil {
		return err
	}

	p.v = v
	return nil
}
