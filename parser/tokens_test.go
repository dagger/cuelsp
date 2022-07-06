package parser

import (
	"regexp"
	"testing"

	cueparser "cuelang.org/go/cue/parser"
	"github.com/stretchr/testify/require"
)

var tokensTest = map[string]struct {
	input  string
	output []string
}{
	"simple": {
		input: `package main

		import (
		"dagger.io/dagger"
		)

		jo: #"""
		\(ok) is good
		"""#

		#Foo: {
			bar: string
		}`,
		output: []string{
			`#Foo: s[7:1]|e[7:5]`,
		},
	},
}

func TestTokensParsing(t *testing.T) {
	for name, tc := range tokensTest {
		t.Run(name, func(t *testing.T) {
			// trim multiline tabs
			re := regexp.MustCompile(`		`)
			strippedStr := re.ReplaceAllString(tc.input, "")

			f, err := cueparser.ParseFile("test", strippedStr)
			if err != nil {
				t.Fatal(err)
			}

			scan := Scanner{}
			ParseSemanticTokens(&scan, f)
			// output := defs.String()
			// for _, o := range tc.output {
			// require.Contains(t, output, o)
			// }
			require.Fail(t, "not implemented")
		})
	}
}
