package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToDef(t *testing.T) {
	type TestCase struct {
		name   string
		input  string
		output *Definition
	}

	testsCases := []TestCase{
		{
			name:   "internal definition",
			input:  "#Test",
			output: &Definition{def: "#Test"},
		},
		{
			name:   "imported definition",
			input:  "pkg.#Test",
			output: &Definition{isImported: true, pkg: "pkg", def: "#Test"},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			def := StringToDef(tt.input)

			assert.Equal(t, tt.output, def)

			// Verify getter
			assert.Equal(t, tt.output.IsImported(), def.IsImported())
			assert.Equal(t, tt.output.Pkg(), def.Pkg())
			assert.Equal(t, tt.output.Def(), def.Def())

		})
	}
}
