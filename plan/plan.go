package plan

import (
	"fmt"

	"github.com/dagger/dlsp/loader"
	"github.com/tliron/kutil/logging"
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
	v *loader.Value

	// Imported packages
	// We use a map because for performance reason
	// See https://boltandnuts.wordpress.com/2017/11/20/go-slice-vs-maps/
	imports map[string]*loader.Instance

	log logging.Logger
}

// New load a new cue value
func New(root, file string) (*Plan, error) {
	k := File
	i, err := loader.File(root, file)

	if err != nil {
		i, err = loader.Dir(root, file)
		if err != nil {
			return nil, err
		}

		k = Directory
	}

	v, err := i.GetValue()
	if err != nil {
		return nil, err
	}

	// Load cue value
	p := &Plan{
		root:     root,
		file:     file,
		kind:     k,
		instance: i,
		v:        v,
		log:      logging.GetLogger(fmt.Sprintf("plan: %s", file)),
		imports:  make(map[string]*loader.Instance),
	}

	if err := p.loadImports(); err != nil {
		return nil, err
	}

	if err := p.instance.LoadDefinitions(); err != nil {
		return nil, err
	}

	return p, nil
}

// loadImports will explore plan's value and list all definitions contained
// in current values and imported packages
func (p *Plan) loadImports() error {
	for _, i := range p.instance.Imports {
		i := loader.NewInstance(i)
		err := i.LoadDefinitions()
		if err != nil {
			return err
		}

		p.imports[i.PkgName] = i
	}

	return nil
}

// GetDefinition return a value following a path
// TODO(TomChv): define path format
// TODO(TomChv): Can be optimized with path, for instance
// - `.#Foo` = definition in current plan
// - `pkg.#Bar` = definition in package pkg
func (p *Plan) GetDefinition(path string) (*loader.Value, error) {
	// Look definition in current plan
	v, _ := p.instance.GetDefinition(path)
	if v != nil {
		return v, nil
	}

	// Look definition in imports
	for _, i := range p.imports {
		v, err := i.GetDefinition(path)
		if err != nil {
			continue
		}

		return v, nil
	}

	return nil, fmt.Errorf("definition %s not found", path)
}

// Reload will rebuild the cue value
func (p *Plan) Reload() error {
	var (
		i   *loader.Instance
		v   *loader.Value
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

	v, err = i.GetValidatedValue()
	if err != nil {
		return err
	}

	p.instance = i
	p.v = v

	if err := p.loadImports(); err != nil {
		return err
	}

	if err := p.instance.LoadDefinitions(); err != nil {
		return err
	}

	return nil
}
