package loader

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/cuecontext"
)

// Instance is a wrapper around Cue instance to implement
// additional methods
type Instance struct {
	*build.Instance
}

func (i *Instance) GetValue() (cue.Value, error) {
	cuectx := cuecontext.New()

	v := cuectx.BuildInstance(i.Instance)
	if err := v.Err(); err != nil {
		return cue.Value{}, err
	}
	if err := v.Validate(); err != nil {
		return cue.Value{}, err
	}

	return v, nil
}

// Validate verify that instance value is correct
// It return error if there is a problem
func (i *Instance) Validate() error {
	_, err := i.GetValue()
	if err != nil {
		return err
	}

	return nil
}
