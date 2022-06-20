package workspace

import (
	"fmt"
	"path/filepath"
	"strings"
)

// TrimRootPath will remove the root path of the file
func (wk *Workspace) TrimRootPath(file string) string {
	p, _ := filepath.Rel(wk.path, file)

	if !strings.HasPrefix(p, "./") {
		p = fmt.Sprintf("./%s", p)
	}

	return p
}
