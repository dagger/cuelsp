package file

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestSourceDir = "./testdata"

func TestNew(t *testing.T) {
	type TestCase struct {
		name       string
		success    bool
		input      string
		outputDefs []string
	}

	testsCases := []TestCase{
		{
			name:       "simple",
			success:    true,
			input:      filepath.Join(TestSourceDir, "simple.cue"),
			outputDefs: []string{"#House", "#House"},
		},
		{
			name:       "other file in same directory",
			success:    true,
			input:      filepath.Join(TestSourceDir, "other.cue"),
			outputDefs: []string{"#Human", "#Human", "#Human", "#Unused"},
		},
		{
			name:       "file in nested directory",
			success:    true,
			input:      filepath.Join(TestSourceDir, "nested/nested.cue"),
			outputDefs: []string{"#Continent", "#Continent", "#Continent", "#Tree", "#Tree", "#Tree"},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := New(tt.input)
			if !tt.success {
				assert.NotNil(t, err)
			}

			assert.Nil(t, err)
			var defs []string
			for _, d := range *(f.Defs()) {
				for _, r := range d {
					defs = append(defs, r.Name())
				}
			}

			assert.ElementsMatch(t, tt.outputDefs, defs)
		})
	}

}
