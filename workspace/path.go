package workspace

import "strings"

// trimRootPath will remove the root path of the file
func (wk *Workspace) trimRootPath(file string) string {
	return "." + strings.TrimPrefix(file, wk.path)
}
