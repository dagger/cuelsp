package parser

import (
	"fmt"
	"regexp"
	"testing"

	cueparser "cuelang.org/go/cue/parser"
	"github.com/stretchr/testify/require"
)

var testCases = map[string]struct {
	input  string
	output []string
}{
	"simple": {
		input: `package main

		import (
		"dagger.io/dagger"
		)

		#Foo: {
			bar: string
		}`,
		output: []string{
			`#Foo: s[7:1]|e[7:5]`,
		},
	},
	"with definition": {
		input: `package main

		#Plan: {
			jo: string
		}

		ok: #Plan & {
			jo: "no"
		}

		dagger.#Plan`,
		output: []string{
			`#Plan: s[3:1]|e[3:6]`,
			`#Plan: s[7:5]|e[7:10]`,
			`dagger.#Plan: s[11:1]|e[11:13]`,
		},
	},
	"many refs": {
		input: `package main

		import (
		   "dagger.io/dagger"
		)
		
		#Bar: {
		  foo: string
		}
		
		#Test: {
		  foo: string
		  too: #Bar
		}
		
		#Plan: {
		  jo: string
		}
		
		ok: #Plan & {
		  jo: "no"
		}
		
		ok: #Jo: #Bar: #Jo: #Bar
		
		deux: #Jo
		
		dagger.#Plan & {
		  actions: test: #Test & {
			"foo": "a"
			#Bar: {
			  "foo": "jo"
			}
			too: #Bar & {
			  "foo": "ok"
			}
		  }
		}`,
		output: []string{
			`#Bar: s[7:1]|e[7:5]`,
			`#Bar: s[13:8]|e[13:12]`,
			`#Plan: s[20:5]|e[20:10]`,
			`#Jo: s[24:5]|e[24:8]`,
			`#Bar: s[24:10]|e[24:14]`,
			`#Jo: s[24:16]|e[24:19]`,
			`#Bar: s[24:21]|e[24:25]`,
			`#Jo: s[26:7]|e[26:10]`,
			`#Bar: s[34:7]|e[34:11]`,
			`#Test: s[11:1]|e[11:6]`,
			`#Plan: s[16:1]|e[16:6]`,
			`dagger.#Plan: s[28:1]|e[28:13]`,
			`#Test: s[29:18]|e[29:23]`,
			`#Bar: s[31:2]|e[31:6]`,
		},
	},
}

func TestDefinitionParsing(t *testing.T) {
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// trim multiline tabs
			re := regexp.MustCompile(`		`)
			strippedStr := re.ReplaceAllString(tc.input, "")

			f, err := cueparser.ParseFile("test", strippedStr)
			if err != nil {
				t.Fatal(err)
			}

			defs := Definitions{}
			ParseDefs(&defs, f)
			output := defs.String()
			for _, o := range tc.output {
				require.Contains(t, output, o)
			}
		})
	}
}

type Input struct {
	line int
	col  int
}

type Output struct {
	name string
	err  error
}

var testFindCases = map[string]struct {
	input  Input
	output Output
}{
	"7:1": {
		input:  Input{line: 7, col: 1},
		output: Output{name: `#Bar`, err: nil},
	},
	"7:4": {
		input:  Input{line: 7, col: 4},
		output: Output{name: `#Bar`, err: nil},
	},
	"7:5": {
		input:  Input{line: 7, col: 5},
		output: Output{name: `#Bar`, err: nil},
	},
	"7:6": {
		input:  Input{line: 7, col: 6},
		output: Output{name: ``, err: fmt.Errorf("definition not found")},
	},
}

func TestFindDefinition(t *testing.T) {
	// Rely on one of above test case (the most complex)
	cueFile, ok := testCases["many refs"]
	if !ok {
		t.Fatal("bad key in testCases")
	}

	// trim multiline tabs
	re := regexp.MustCompile(`		`)
	strippedStr := re.ReplaceAllString(cueFile.input, "")

	for name, tc := range testFindCases {
		t.Run(name, func(t *testing.T) {

			// Use ParseFile wrapper to load URI files instead of this one
			f, err := cueparser.ParseFile("test", strippedStr)
			if err != nil {
				t.Fatal(err)
			}

			defs := Definitions{}
			ParseDefs(&defs, f)

			name, err := defs.Find(tc.input.line, tc.input.col)
			if err != nil {
				require.Equal(t, tc.output.err, err)
			} else {
				require.Equal(t, tc.output.name, name)
			}
		})
	}
}
