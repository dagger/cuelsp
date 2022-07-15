package loader

import (
	"fmt"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/cuecontext"
)

// Instance is a wrapper around Cue instance to implement
// additional methods
// It also caches various data to improve performances and data fetching
type Instance struct {
	*build.Instance
	defs map[string]*Value
}

// NewInstance is a constructor around Instance to correctly initialize
// cache fields
func NewInstance(i *build.Instance) *Instance {
	return &Instance{
		Instance: i,
		defs:     map[string]*Value{},
	}
}

// GetValue return a built instance value
func (i *Instance) GetValue() (*Value, error) {
	cuectx := cuecontext.New()

	v := cuectx.BuildInstance(i.Instance)
	if err := v.Err(); err != nil {
		return nil, err
	}

	return &Value{v}, nil
}

// GetValidatedValue return a value if it's validated
func (i *Instance) GetValidatedValue() (*Value, error) {
	v, err := i.GetValue()
	if err != nil {
		return nil, err
	}

	if err := v.Validate(); err != nil {
		return nil, err
	}

	return v, nil
}

// Validate verify that instance value is correct
// It return error if there is an error
func (i *Instance) Validate() error {
	if _, err := i.GetValidatedValue(); err != nil {
		return err
	}

	return nil
}

// LoadDefinitions list and store all definitions of an instance
func (i *Instance) LoadDefinitions() error {
	v, err := i.GetValue()
	if err != nil {
		return err
	}

	defs, err := v.ListDefinitions()
	if err != nil {
		return err
	}

	for _, d := range defs {
		i.defs[d.Path().String()] = d
	}

	return nil
}

// GetDefinition return a definition if found or an error
func (i *Instance) GetDefinition(path string) (*Value, error) {
	v, found := i.defs[path]
	if !found {
		return nil, fmt.Errorf("definition %s not found", path)
	}

	return v, nil
}

func (i *Instance) GetNode(def string) (ast.Node, error) {
	for _, file := range i.Files {
		for _, node := range file.Decls {
			switch n := node.(type) {
			case *ast.Field:
				if fmt.Sprintf("%s", n.Label) == def {
					return node, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("node %s not found", def)
}

func (i *Instance) String() string {
	var defs string
	for k := range i.defs {
		defs += fmt.Sprintf("\n\t-\t%s", k)
	}

	return fmt.Sprintf("Instance: %s\nDefs:%s\n", i.PkgName, defs)
}
