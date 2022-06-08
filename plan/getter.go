package plan

import (
	"fmt"

	"github.com/dagger/dlsp/file"
	loader2 "github.com/dagger/dlsp/loader"
)

func (p *Plan) RootPath() string {
	return p.rootPath
}

func (p *Plan) RootFilePath() string {
	return p.rootFilePath
}

func (p *Plan) Files() map[string]*file.File {
	return p.files
}

func (p *Plan) Kind() Kind {
	return p.kind
}

func (p *Plan) Value() *loader2.Value {
	return p.v
}

func (p *Plan) Instance() *loader2.Instance {
	return p.instance
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
		files += fmt.Sprintf("\n%s", f)
	}

	return fmt.Sprintf("Root: %s/%s Files: %s\nType: %s, Value: %s\n%s\n Imports: %s\n", p.rootPath, p.RootFilePath, files, p.kind, p.v, p.instance, imports)
}
