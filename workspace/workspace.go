package workspace

import (
	"fmt"
	"path/filepath"

	plan2 "github.com/dagger/dlsp/plan"
	"github.com/tliron/kutil/logging"
)

// Workspace is a representation of working directory
type Workspace struct {
	path  string
	log   logging.Logger
	plans []*plan2.Plan
}

// New initialize a new workspace in the language server
func New(path string) *Workspace {
	return &Workspace{
		path: path,
		log:  logging.GetLogger("workspace"),
	}
}

// AddPlan load a new plan into the workspace
// It does not load a plan if it already exists
func (wk *Workspace) AddPlan(file string) error {
	if wk.isPlanLoaded(file) {
		wk.log.Debugf("plan %s is already loaded", file)
		return nil
	}

	wk.log.Debugf("Add new plan: %s", file)

	file = wk.trimRootPath(file)
	wk.log.Debugf("Source: %s; File: %s", wk.path, file)

	newPlan, err := plan2.New(wk.path, file)
	if err != nil {
		return err
	}

	wk.log.Debugf("NewPlan: %s", newPlan)
	wk.plans = append(wk.plans, newPlan)
	return nil
}

// isPlanLoaded verify if a plan has been already loaded
// in the workspace
func (wk *Workspace) isPlanLoaded(file string) bool {
	wk.log.Debugf("Looking for plan %s", file)

	p, _ := wk.GetPlan(file)
	return p != nil
}

// GetPlan return the plan selected at the root
func (wk *Workspace) GetPlan(file string) (*plan2.Plan, error) {
	wk.log.Debugf("Call to Get plan: %s", file)
	file = wk.trimRootPath(file)

	for _, p := range wk.plans {
		wk.log.Debugf("compare %s %s with %s", p.Kind(), p.File(), file)
		if p.Kind() == plan2.File && p.File() == file {
			return p, nil
		}
		if p.Kind() == plan2.Directory && filepath.Dir(p.File()) == filepath.Dir(file) {
			return p, nil
		}
	}

	wk.log.Debugf("Plan not found %s", file)
	return nil, fmt.Errorf("plan %s not found", file)
}
