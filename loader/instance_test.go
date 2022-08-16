package loader

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInstance_GetDefinition load definitions of an instance and verify that
// all definitions are correctly loaded.
// It uses TestSourceDir defined in build_test.go as source dir
func TestInstance_GetDefinition(t *testing.T) {
	type TestCase struct {
		name   string
		src    string
		file   string
		loader func(src, file string) (*Instance, error)
		defs   []string
	}

	testsCases := []TestCase{
		{
			name:   "single def",
			src:    TestSourceDir,
			file:   "main.cue",
			loader: File,
			defs:   []string{"#Num"},
		},
		{
			name:   "single def in other file",
			src:    TestSourceDir,
			file:   "main2.cue",
			loader: File,
			defs:   []string{"#String"},
		},
		{
			name:   "single file with cue.mod",
			src:    filepath.Join(TestSourceDir, "with-cue-mod"),
			file:   "main.cue",
			loader: File,
			defs:   []string{"_#TestName"},
		},
		{
			name:   "multi file",
			src:    TestSourceDir,
			file:   filepath.Join("dir-multi-files", "multi.cue"),
			loader: Dir,
			defs:   []string{"#Action", "#Plan"},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			i, err := tt.loader(tt.src, tt.file)
			assert.Nil(t, err)

			err = i.LoadDefinitions()
			assert.Nil(t, err)

			for _, def := range tt.defs {
				v, err := i.GetDefinition(def)
				assert.Nil(t, err)
				assert.Equal(t, def, v.Path().String())
			}
		})
	}
}

func TestInstance_GetDefinition_NotFound(t *testing.T) {
	type TestCaseStruct struct {
		name string
		src  string
		file string
	}

	testsCases := []TestCaseStruct{
		{
			name: "single def",
			src:  TestSourceDir,
			file: "main.cue",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			i, err := File(tt.src, tt.file)
			assert.Nil(t, err)

			err = i.LoadDefinitions()
			assert.Nil(t, err)

			_, err = i.GetDefinition("Unknown")
			assert.NotNil(t, err)
		})
	}
}
