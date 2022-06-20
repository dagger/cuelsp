package plan

import (
	"fmt"
	"path/filepath"

	"github.com/dagger/dlsp/file"
)

func (p *Plan) Files() map[string]*file.File {
	return p.files
}

func (p *Plan) String() string {
	var imports string
	for _, i := range p.imports {
		imports += fmt.Sprintf("\n- %s", i)
	}

	var files string
	p.muFiles.RLock()
	defer p.muFiles.RUnlock()
	for _, f := range p.files {
		files += fmt.Sprintf("\n- %s", f)
	}

	return fmt.Sprintf("Root: %s\nFiles: %s\nType: %s\nValue: %s\n%s\n Imports: %s\n", filepath.Join(p.rootPath, p.RootFilePath), files, p.Kind, p.v, p.instance, imports)
}
