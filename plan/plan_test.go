package plan

import (
	"path/filepath"
	"testing"

	"github.com/dagger/dlsp/loader"
	"github.com/stretchr/testify/assert"
)

const TestSourceDir = "./testdata"

func TestNew(t *testing.T) {
	type TestCase struct {
		name string
		root string
		file string
		out  *Plan
	}

	testsCases := []TestCase{
		{
			name: "single file main",
			root: TestSourceDir,
			file: filepath.Join(TestSourceDir, "main.cue"),
			out: &Plan{
				rootPath:     TestSourceDir,
				RootFilePath: "main.cue",
				Kind:         File,
				imports:      map[string]*loader.Instance{},
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := New(tt.root, tt.file)
			assert.Nil(t, err)

			// Verify simple data
			assert.Equal(t, tt.out.rootPath, p.rootPath)
			assert.Equal(t, tt.out.RootFilePath, p.RootFilePath)
			assert.Equal(t, tt.out.Kind, p.Kind)
			assert.Equal(t, len(tt.out.imports), len(p.imports))
		})
	}
}
