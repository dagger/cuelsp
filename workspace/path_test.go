package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tliron/kutil/logging"
)

func TestWorkspace_TrimRootPath(t *testing.T) {
	type TestCase struct {
		name     string
		rootPath string
		input    string
		output   string
	}

	testsCases := []TestCase{
		{
			name:     "absolute path",
			rootPath: "/foo/bar",
			input:    "/foo/bar/test",
			output:   "./test",
		},
		{
			name:     "relative path",
			rootPath: "/foo/bar",
			input:    "/foo/bar/./test",
			output:   "./test",
		},
		{
			name:     "empty path",
			rootPath: "/foo/bar",
			input:    "",
			output:   "./",
		},
		{
			name:     "relative nested path",
			rootPath: "./testdata",
			input:    "./testdata/main.cue",
			output:   "./main.cue",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			ws := New(tt.rootPath, logging.GetLogger("test"))

			p := ws.TrimRootPath(tt.input)
			assert.Equal(t, tt.output, p)
		})
	}
}
