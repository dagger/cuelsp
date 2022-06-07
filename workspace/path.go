package workspace

import "strings"

// TrimRootPath will remove the root path of the file
func (wk *Workspace) TrimRootPath(file string) string {
	return "." + strings.TrimPrefix(file, wk.path)
}
