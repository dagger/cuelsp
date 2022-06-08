package workspace

import (
	"path/filepath"

	"github.com/dagger/dlsp/plan"
	"github.com/tliron/kutil/logging"
)

// Workspace is a representation of working directory
type Workspace struct {
	path  string
	log   logging.Logger
	plans []*plan.Plan
}

// New initialize a new workspace in the language server
func New(path string, logger logging.Logger) *Workspace {
	return &Workspace{
		path: path,
		log:  logging.NewScopeLogger(logger, "workspace"),
	}
}

// AddPlan load a new plan into the workspace
// It does not load a plan if it already exists
func (wk *Workspace) AddPlan(file string) error {
	p := wk.GetPlan(file)
	if p != nil {
		wk.log.Debugf("plan %s is already loaded", file)

		err := p.AddFile(wk.TrimRootPath(file))
		wk.log.Debugf("Plan: %s", p)
		return err
	}

	wk.log.Debugf("Add new plan: %s", file)

	file = wk.TrimRootPath(file)
	wk.log.Debugf("Source: %s; File: %s", wk.path, file)

	newPlan, err := plan.New(wk.path, file)
	if err != nil {
		return err
	}

	wk.log.Debugf("NewPlan: %s", newPlan)
	wk.plans = append(wk.plans, newPlan)
	return nil
}

// GetPlan return the plan selected at the root
// Return nil if not found
func (wk *Workspace) GetPlan(file string) *plan.Plan {
	wk.log.Debugf("Call to Get plan: %s", file)
	file = wk.TrimRootPath(file)

	for _, p := range wk.plans {
		wk.log.Debugf("compare %s %s with %s", p.Kind(), p.RootFilePath(), file)
		if p.Kind() == plan.File && p.RootFilePath() == file {
			return p
		}
		if p.Kind() == plan.Directory && filepath.Dir(p.RootFilePath()) == filepath.Dir(file) {
			return p
		}
	}

	wk.log.Debugf("Plan not found %s", file)
	return nil
}
