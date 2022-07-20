package plan

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/dagger/daggerlsp/file"
	"github.com/dagger/daggerlsp/internal"
	"github.com/dagger/daggerlsp/loader"
	"github.com/tliron/kutil/logging"
)

// Plan is a representation of a cue value in a workspace
type Plan struct {
	// rootPath is the plan's root path.
	rootPath string

	// RootFilePath is the plan's root file path.
	RootFilePath string

	// muFiles protects the access to the files map.
	muFiles sync.RWMutex

	// files store the loaded files.
	files map[string]*file.File

	// Kind stores Plan's Kind.
	Kind Kind

	// Plan's instance
	instance *loader.Instance

	// v represents the CUE Value
	v *loader.Value

	// Imported packages
	// We use a map because for performance reason
	// See https://boltandnuts.wordpress.com/2017/11/20/go-slice-vs-maps/
	imports map[string]*loader.Instance

	log logging.Logger
}

// New load a new cue value
func New(root, filePath string) (*Plan, error) {
	log := logging.GetLogger(fmt.Sprintf("root: %s, plan: %s", root, filePath))

	k := File
	log.Debugf("Try to load plan as file")
	i, err := loader.File(root, filePath)

	if err != nil {
		log.Debugf("Try to load plan as directory")
		i, err = loader.Dir(root, filePath)
		if err != nil {
			return nil, err
		}

		k = Directory
	}

	log.Debugf("Plan loaded")

	v, err := i.GetValue()
	if err != nil {
		return nil, err
	}

	f, err := file.New(filepath.Join(root, filePath))
	if err != nil {
		return nil, err
	}

	files := map[string]*file.File{}
	files[filePath] = f

	// Load cue value
	p := &Plan{
		rootPath:     root,
		RootFilePath: filePath,
		files:        files,
		Kind:         k,
		instance:     i,
		v:            v,
		log:          log,
		imports:      map[string]*loader.Instance{},
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

func (p *Plan) loadFiles() error {
	for path := range p.files {
		f, err := file.New(filepath.Join(p.rootPath, path))
		if err != nil {
			return err
		}

		p.files[path] = f
	}

	return nil
}

// GetDefinition return a value following a path
// - `.#Foo` = definition in current plan
// - `pkg.#Bar` = definition in package pkg
func (p *Plan) GetDefinition(path string, line, char int) (*loader.Value, error) {
	def, err := p.findDefInFile(path, line, char)
	if err != nil {
		return nil, err
	}

	p.log.Debugf("%#v", def)
	if !def.IsImported() {
		// Look definition in current plan
		return p.instance.GetDefinition(def.Def())
	} else {
		i, found := p.imports[def.Pkg()]
		if !found {
			return nil, fmt.Errorf("imported package %s not registered in plan", def.Def())
		}

		return i.GetDefinition(def.Def())
	}
}

func (p *Plan) GetDocDefinition(path string, line, char int) (*internal.DocValue, error) {
	def, err := p.findDefInFile(path, line, char)
	if err != nil {
		return nil, err
	}

	i, err := p.GetInstance(path, line, char)
	if err != nil {
		return nil, err
	}

	node, err := i.GetNode(def.Def())
	if err != nil {
		return nil, err
	}

	v, err := i.GetDefinition(def.Def())
	if err != nil {
		return nil, err
	}

	return internal.NewDocValue(node, v), nil
}

func (p *Plan) GetInstance(path string, line, char int) (*loader.Instance, error) {
	def, err := p.findDefInFile(path, line, char)
	if err != nil {
		return nil, err
	}

	p.log.Debugf("%#v", def)
	if !def.IsImported() {
		return p.instance, nil
	} else {
		i, found := p.imports[def.Pkg()]
		if !found {
			return nil, fmt.Errorf("imported package %s not registered in plan", def.Def())
		}

		return i, nil
	}
}

func (p *Plan) findDefInFile(path string, line, char int) (*internal.Definition, error) {
	p.log.Debugf("Looking for file: %s", path)

	p.muFiles.RLock()
	defer p.muFiles.RUnlock()

	f, ok := p.files[path]
	if !ok {
		return nil, fmt.Errorf("file not registered")
	}

	p.log.Debugf("Looking for def in %s at {%d, %d}", path, line, char)
	def, err := f.Defs().Find(line, char)
	if err != nil {
		return nil, err
	}

	p.log.Debugf("Searching for %s in value", def)

	return internal.StringToDef(def), nil
}

func (p *Plan) GetDoc(path string, line, char int) (*loader.Value, error) {
	p.log.Debugf("Looking for file: %s", path)

	p.muFiles.RLock()
	defer p.muFiles.RUnlock()

	f, ok := p.files[path]
	if !ok {
		return nil, fmt.Errorf("file not registered")
	}

	p.log.Debugf("Looking for def in %s at {%d, %d}", path, line, char)
	def, err := f.Defs().Find(line, char)
	if err != nil {
		return nil, err
	}

	p.log.Debugf("Searching for %s in value", def)

	_def := internal.StringToDef(def)

	p.log.Debugf("%#v", _def)
	if !_def.IsImported() {
		// Look definition in current plan
		return p.instance.GetDefinition(_def.Def())
	} else {
		i, found := p.imports[_def.Pkg()]
		if !found {
			return nil, fmt.Errorf("imported package %s not registered in plan", _def.Def())
		}

		return i.GetValue()
	}
}

// Reload will rebuild the cue value
func (p *Plan) Reload() error {
	var (
		i   *loader.Instance
		v   *loader.Value
		err error
	)

	switch p.Kind {
	case File:
		i, err = loader.File(p.rootPath, p.RootFilePath)
	case Directory:
		i, err = loader.Dir(p.rootPath, p.RootFilePath)
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

	if err := p.loadFiles(); err != nil {
		return err
	}

	return nil
}

// AddFile load and register a new file in the plan
// This file must be part of the instance
func (p *Plan) AddFile(path string) error {
	p.log.Debugf("Add a new file to plan: %s", path)

	f, err := file.New(filepath.Join(p.rootPath, path))
	if err != nil {
		return err
	}
	p.muFiles.Lock()
	p.files[path] = f
	p.muFiles.Unlock()

	return nil
}
