package plan

import (
	"fmt"
	"path/filepath"
	"testing"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
	"github.com/dagger/daggerlsp/loader"
	"github.com/stretchr/testify/assert"
)

const TestSourceDir = "./testdata"

// TestNew verify that plans are correctly loaded
// /!\ It doesn't test much about methods, simply about loading
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
			file: "main.cue",
			out: &Plan{
				rootPath:     TestSourceDir,
				RootFilePath: "main.cue",
				Kind:         File,
				imports:      map[string]*loader.Instance{},
			},
		},
		{
			name: "directory with multi-files",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			out: &Plan{
				rootPath:     TestSourceDir,
				RootFilePath: filepath.Join("dir-multi-files", "multi.cue"),
				Kind:         Directory,
				imports:      map[string]*loader.Instance{},
			},
		},
		{
			name: "file with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			out: &Plan{
				rootPath:     filepath.Join(TestSourceDir, "with-cue-mod"),
				RootFilePath: "main.cue",
				Kind:         File,
				imports: map[string]*loader.Instance{
					"test": nil,
				},
			},
		},
		{
			name: "dir with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "path.cue"),
			out: &Plan{
				rootPath:     filepath.Join(TestSourceDir, "with-cue-mod"),
				RootFilePath: filepath.Join("dir", "path.cue"),
				Kind:         Directory,
				imports: map[string]*loader.Instance{
					"test": nil,
				},
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

			// Check that String do not crash
			assert.NotEmpty(t, p.String())
		})
	}
}

func TestPlan_GetDefinition(t *testing.T) {
	type Def struct {
		path string
		line int
		char int
	}

	type TestCase struct {
		name string
		root string
		file string
		defs []Def
	}

	testsCases := []TestCase{
		{
			name: "single definition",
			root: TestSourceDir,
			file: "./main.cue",
			defs: []Def{
				{
					path: "#Num",
					line: 3,
					char: 1,
				},
			},
		},
		{
			name: "directory with multi-files with multi values",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			defs: []Def{
				{
					path: "#Plan",
					line: 3,
					char: 1,
				},
				{
					path: "#Action",
					line: 7,
					char: 3,
				},
				{
					path: "#Action",
					line: 11,
					char: 3,
				},
			},
		},
		{
			name: "directory with multi-files - found in range",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			defs: []Def{
				{
					path: "#Action",
					line: 11,
					char: 3,
				},
				{
					path: "#Action",
					line: 11,
					char: 4,
				},
				{
					path: "#Action",
					line: 11,
					char: 5,
				},
				{
					path: "#Action",
					line: 11,
					char: 6,
				},
				{
					path: "#Action",
					line: 11,
					char: 10,
				},
			},
		},
		{
			name: "file with cue.mod - private definitions and imported",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			defs: []Def{
				{
					path: "_#TestName",
					line: 7,
					char: 1,
				},
				{
					path: "#Test",
					line: 9,
					char: 9,
				},
				{
					path: "_#TestName",
					line: 15,
					char: 12,
				},
				{
					path: "#Test",
					line: 14,
					char: 15,
				},
			},
		},
		{
			name: "dir with cue.mod and merged values",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "path.cue"),
			defs: []Def{
				{
					path: "#Path",
					line: 7,
					char: 6,
				},
				{
					path: "#Test",
					line: 9,
					char: 9,
				},
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := New(tt.root, tt.file)
			assert.Nil(t, err)

			// Get definition
			for _, def := range tt.defs {
				d, err := p.GetDefinition(tt.file, def.line, def.char)
				assert.Nil(t, err)
				assert.Equal(t, true, d.IsDefinition())
				assert.Equal(t, def.path, d.Path().String())
			}
		})
	}
}

func TestPlan_GetDefinition_NotFound(t *testing.T) {
	type Def struct {
		path string
		line int
		char int
	}

	type TestCase struct {
		name string
		root string
		file string
		def  Def
	}

	testsCases := []TestCase{
		{
			name: "file not found",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "unknown",
			},
		},
		{
			name: "line not found",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "./main.cue",
				line: 1,
			},
		},
		{
			name: "char not found",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "./main.cue",
				line: 3,
				char: 9,
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := New(tt.root, tt.file)
			assert.Nil(t, err)

			// Get definition
			d, err := p.GetDefinition(tt.def.path, tt.def.line, tt.def.char)
			assert.NotNil(t, err)
			assert.Nil(t, d)
		})
	}
}

func TestPlan_GetInstance(t *testing.T) {
	type Def struct {
		path string
		line int
		char int
	}

	type TestCase struct {
		name string
		root string
		file string
		defs []Def
	}

	testsCases := []TestCase{
		{
			name: "single definition",
			root: TestSourceDir,
			file: "./main.cue",
			defs: []Def{
				{
					path: "#Num",
					line: 3,
					char: 1,
				},
			},
		},
		{
			name: "directory with multi-files with multi values",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			defs: []Def{
				{
					path: "#Plan",
					line: 3,
					char: 1,
				},
				{
					path: "#Action",
					line: 7,
					char: 3,
				},
				{
					path: "#Action",
					line: 11,
					char: 3,
				},
			},
		},
		{
			name: "directory with multi-files - found in range",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			defs: []Def{
				{
					path: "#Action",
					line: 11,
					char: 3,
				},
				{
					path: "#Action",
					line: 11,
					char: 4,
				},
				{
					path: "#Action",
					line: 11,
					char: 5,
				},
				{
					path: "#Action",
					line: 11,
					char: 6,
				},
				{
					path: "#Action",
					line: 11,
					char: 10,
				},
			},
		},
		{
			name: "file with cue.mod - private definitions and imported",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			defs: []Def{
				{
					path: "_#TestName",
					line: 7,
					char: 1,
				},
				{
					path: "#Test",
					line: 9,
					char: 9,
				},
				{
					path: "_#TestName",
					line: 15,
					char: 12,
				},
				{
					path: "#Test",
					line: 14,
					char: 15,
				},
			},
		},
		{
			name: "dir with cue.mod and merged values",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "path.cue"),
			defs: []Def{
				{
					path: "#Path",
					line: 7,
					char: 6,
				},
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := New(tt.root, tt.file)
			assert.Nil(t, err)

			// Get definition
			for _, def := range tt.defs {
				i, err := p.GetInstance(tt.file, def.line, def.char)
				assert.Nil(t, err)

				for _, file := range i.Files {
					found := false

					for _, node := range file.Decls {
						switch n := node.(type) {
						case *ast.Field:
							d, err := format.Node(n)
							assert.Nil(t, err)

							fmt.Println(string(d))
							fmt.Println(n.Label, n.Value)
							label := fmt.Sprintf("%s", n.Label)
							if label == def.path {
								found = true
								return
							}
						}
					}
					assert.True(t, found)
				}
			}
		})
	}
}

func TestPlan_AddFile(t *testing.T) {
	type TestCase struct {
		name      string
		root      string
		file      string
		fileToAdd string
		success   bool
	}

	testCases := []TestCase{
		{
			name:      "add file in dir",
			root:      filepath.Join(TestSourceDir, "with-cue-mod"),
			file:      filepath.Join("dir", "path.cue"),
			fileToAdd: filepath.Join("dir", "dir.cue"),
			success:   true,
		},
		{
			name:      "wrong file in dir",
			root:      filepath.Join(TestSourceDir, "with-cue-mod"),
			file:      filepath.Join("dir", "path.cue"),
			fileToAdd: filepath.Join("dir", "unknown.cue"),
			success:   false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := New(tt.root, tt.file)
			assert.Nil(t, err)

			err = p.AddFile(tt.fileToAdd)
			if !tt.success {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)

			_, ok := p.files[tt.file]
			assert.Equal(t, true, ok)

			_, ok = p.files[tt.fileToAdd]
			assert.Equal(t, true, ok)
		})
	}
}

// TestPlan_Reload verify that plans are correctly reloaded
// /!\ It doesn't test much about methods, simply about reloaded
// FIXME(TomChv): Find a way to test reload
func TestPlan_Reload(t *testing.T) {
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
			file: "main.cue",
			out: &Plan{
				rootPath:     TestSourceDir,
				RootFilePath: "main.cue",
				Kind:         File,
				imports:      map[string]*loader.Instance{},
			},
		},
		{
			name: "directory with multi-files",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			out: &Plan{
				rootPath:     TestSourceDir,
				RootFilePath: filepath.Join("dir-multi-files", "multi.cue"),
				Kind:         Directory,
				imports:      map[string]*loader.Instance{},
			},
		},
		{
			name: "file with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			out: &Plan{
				rootPath:     filepath.Join(TestSourceDir, "with-cue-mod"),
				RootFilePath: "main.cue",
				Kind:         File,
				imports: map[string]*loader.Instance{
					"test": nil,
				},
			},
		},
		{
			name: "dir with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "path.cue"),
			out: &Plan{
				rootPath:     filepath.Join(TestSourceDir, "with-cue-mod"),
				RootFilePath: filepath.Join("dir", "path.cue"),
				Kind:         Directory,
				imports: map[string]*loader.Instance{
					"test": nil,
				},
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

			err = p.Reload()
			assert.Nil(t, err)

			// Verify simple data
			assert.Equal(t, tt.out.rootPath, p.rootPath)
			assert.Equal(t, tt.out.RootFilePath, p.RootFilePath)
			assert.Equal(t, tt.out.Kind, p.Kind)
			assert.Equal(t, len(tt.out.imports), len(p.imports))
		})
	}
}
