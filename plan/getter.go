package plan

import (
	"fmt"
)

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

	return fmt.Sprintf("Root: %s/%s Files: %s\nType: %s, Value: %s\n%s\n Imports: %s\n", p.rootPath, p.RootFilePath, files, p.Kind, p.v, p.instance, imports)
}
