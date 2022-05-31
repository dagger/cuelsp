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

	// Plan's instance
	instance *loader.Instance

	// Cue Value
	v cue.Value

	// Definitions of a plan and his imported package
	defs map[string]cue.Value
}

// New load a new cue value
func New(root, file string) (*Plan, error) {
	k := Directory
	i, err := loader.Dir(root, file)

	if err != nil {
		i, err = loader.File(root, file)
		if err != nil {
			return nil, err
		}

		k = File
	}

	v, err := i.GetValue()
	if err != nil {
		return nil, err
	}

	// Load cue value
	return &Plan{
		root:     root,
		file:     file,
		kind:     k,
		instance: i,
		v:        v,
	}, nil
}

// LoadDefs will explore plan's value and list all definitions contained
// in current values and imported packages
func (p *Plan) LoadDefs() error {
	return nil
}

// Reload will rebuild the cue value
func (p *Plan) Reload() error {
	var (
		i   *loader.Instance
		v   cue.Value
		err error
	)

	switch p.kind {
	case File:
		i, err = loader.File(p.root, p.file)
	case Directory:
		i, err = loader.Dir(p.root, p.file)
	}

	if err != nil {
		return err
	}

	v, err = i.GetValue()
	if err != nil {
		return err
	}

	p.instance = i
	p.v = v
	return nil
}
