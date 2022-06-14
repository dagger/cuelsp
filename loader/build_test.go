package loader

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestSourceDir = "./testdata"

// TestFile only expect that there is no error
// during the load.
// All tests of functionalities are in instance_test.go
func TestFile(t *testing.T) {
	type TestCase struct {
		name string
		src  string
		file string
		pkg  string
	}

	testsCases := []TestCase{
		{
			name: "simple main file",
			src:  TestSourceDir,
			file: "main.cue",
			pkg:  "main",
		},
		{
			name: "other simple main file",
			src:  TestSourceDir,
			file: "main2.cue",
			pkg:  "main",
		},
		{
			name: "nested simple main file",
			src:  TestSourceDir,
			file: filepath.Join("nested", "nested.cue"),
			pkg:  "nested",
		},
		{
			name: "nested simple main file with cue.mod",
			src:  filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			pkg:  "main",
		},
		{
			name: "nested deep simple main file with cue.mod",
			src:  filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "dir.cue"),
			pkg:  "dir",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			i, err := File(tt.src, tt.file)
			assert.Nil(t, err)
			assert.Equal(t, tt.pkg, i.PkgName)
		})
	}
}

// TestDir only expect that there is no error
// during the load.
// All tests of functionalities are in instance_test.go
func TestDir(t *testing.T) {
	type TestCase struct {
		name string
		src  string
		file string
		pkg  string
	}

	testsCases := []TestCase{
		{
			name: "root dir",
			src:  TestSourceDir,
			file: "main.cue",
			pkg:  "main",
		},
		{
			name: "root dir single file",
			src:  TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			pkg:  "action",
		},
		{
			name: "root dir multi file",
			src:  TestSourceDir,
			file: filepath.Join("dir-single-file", "simple.cue"),
			pkg:  "simple",
		},
		{
			name: "nested dir with root source",
			src:  TestSourceDir,
			file: filepath.Join("nested", "nested.cue"),
			pkg:  "nested",
		},
		{
			name: "nested dir with nested root source",
			src:  filepath.Join(TestSourceDir, "nested"),
			file: filepath.Join("dir-nested", "nested.cue"),
			pkg:  "dirnested",
		},
		{
			name: "dir with nested source with cue.mod",
			src:  filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			pkg:  "main",
		},
		{
			name: "nested dir with nested source with cue.mod",
			src:  filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "dir.cue"),
			pkg:  "dir",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			i, err := Dir(tt.src, tt.file)
			assert.Nil(t, err)
			assert.Equal(t, tt.pkg, i.PkgName)
		})
	}
}
