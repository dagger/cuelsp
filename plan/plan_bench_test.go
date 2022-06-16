package plan

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkNew(b *testing.B) {
	type TestCase struct {
		name string
		root string
		file string
	}

	testsCases := []TestCase{
		{
			name: "single file main",
			root: TestSourceDir,
			file: "main.cue",
		},
		{
			name: "directory with multi-files",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
		},
		{
			name: "file with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
		},
		{
			name: "dir with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "path.cue"),
		},
	}

	b.ResetTimer()
	for _, tt := range testsCases {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = New(tt.root, tt.file)
			}
		})
	}
}

func BenchmarkPlan_GetDefinition(b *testing.B) {
	type Def struct {
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
			name: "single file main",
			root: TestSourceDir,
			file: "main.cue",
			def: Def{
				line: 3,
				char: 1,
			},
		},
		{
			name: "directory with multi-files",
			root: TestSourceDir,
			file: filepath.Join("dir-multi-files", "multi.cue"),
			def: Def{
				line: 11,
				char: 10,
			},
		},
		{
			name: "file with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: "main.cue",
			def: Def{
				line: 15,
				char: 12,
			},
		},
		{
			name: "dir with cue.mod",
			root: filepath.Join(TestSourceDir, "with-cue-mod"),
			file: filepath.Join("dir", "path.cue"),
			def: Def{
				line: 9,
				char: 9,
			},
		},
	}

	b.ResetTimer()
	for _, tt := range testsCases {
		b.StopTimer()
		p, err := New(tt.root, tt.file)
		assert.Nil(b, err)
		b.StartTimer()

		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = p.GetDefinition(tt.file, tt.def.line, tt.def.char)
			}
		})
	}
}
