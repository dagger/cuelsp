package util

import "strings"

// UriToPath convert a URI into a usable path
// Example
// file:///foo/bar -> /foo/bar
// If there is no URI prefix, the string is returned unchanged
func UriToPath(uri string) string {
	return strings.TrimPrefix(uri, "file://")
}
