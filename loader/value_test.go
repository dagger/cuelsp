package loader

import (
	"testing"

	"cuelang.org/go/cue/cuecontext"
	"github.com/stretchr/testify/assert"
)

func TestValue_ListDefinitions(t *testing.T) {
	ctx := cuecontext.New()

	type TestCase struct {
		name   string
		value  *Value
		output []string
	}

	testsCases := []TestCase{
		{
			name:   "inline definition",
			value:  &Value{ctx.CompileString(`#Test: string`)},
			output: []string{"#Test"},
		},
		{
			name:   "no definition",
			value:  &Value{ctx.CompileString(`a: 4`)},
			output: []string{},
		},
		{
			name: "multi definition",
			value: &Value{ctx.CompileString(`
				#A: 4
				#B: "foo"
				#C: true
			`)},
			output: []string{"#A", "#B", "#C"},
		},
		{
			name: "private definition",
			value: &Value{ctx.CompileString(`
				#A: 4
				_#B: "foo"
			`)},
			output: []string{"#A", "_#B"},
		},
		{
			name: "Reference definition",
			value: &Value{ctx.CompileString(`
				#A: number
				_#B: [#A, #A]
				#C: #A & >0
				#D: _#B
			`)},
			output: []string{"#A", "_#B", "#C", "#D"},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			defs, err := tt.value.ListDefinitions()
			assert.Nil(t, err)

			var list []string
			for _, d := range defs {
				list = append(list, d.Path().String())
			}

			assert.ElementsMatch(t, tt.output, list)
		})
	}
}
