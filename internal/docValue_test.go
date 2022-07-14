package internal_test

import (
	"testing"

	"github.com/dagger/daggerlsp/plan"
	"github.com/stretchr/testify/assert"
)

func TestNewDocValue(t *testing.T) {
	const TestSourceDir = "./testdata"

	type Def struct {
		path string
		line int
		char int
	}

	type TestCase struct {
		name   string
		root   string
		file   string
		def    Def
		expect string
	}

	testCases := []TestCase{
		{
			name: "Simple definition",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "#Simple",
				line: 4,
				char: 1,
			},
			expect: `#### Description
An example of a definition

#### Type
#Simple: string`,
		},
		{
			name: "Multi line doc",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "#SimpleMultiLineDoc",
				line: 8,
				char: 1,
			},
			expect: `#### Description
Multi line
documentation

#### Type
#SimpleMultiLineDoc: number`,
		},
		{
			name: "Structure doc",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "#Struct",
				line: 11,
				char: 1,
			},
			expect: `#### Description
Structure

#### Type
#Struct: {
	bar: string

	foo: *true | bool
}`,
		},
		{
			name: "struct with reference",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "#StructWithDoc",
				line: 19,
				char: 1,
			},
			expect: `#### Description
Struct with doc in
multiline

#### Type
#StructWithDoc: {
	// first field
	bar: string

	// second field
	foo: number | [...number]
}`,
		},
		{
			name: "Struct with doc",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "#ReferencetoStruct",
				line: 29,
				char: 1,
			},
			expect: `#### Description
Multiline documentation
with reference to a definition

#### Type
#ReferenceToStruct: {
	// Multi line field doc
	// reference to a def
	ref: #Struct

	enum: "hello" | "world" | string
}`,
		},
		{
			name: "dagger type",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "#Secret",
				line: 44,
				char: 1,
			},
			expect: `#### Description
Dagger type
A reference to an external secret, for example:
 - A password
 - A SSH private key
 - An API token
Secrets are never merged in the Cue tree. They can only be used
by a special filesystem mount designed to minimize leak risk.

#### Type
#Secret: {
	$dagger: secret: _id: string
}`,
		},
		{
			name: "dagger type complex",
			root: TestSourceDir,
			file: "./main.cue",
			def: Def{
				path: "_#clientCommand",
				line: 49,
				char: 1,
			},
			expect: `#### Description
Complex command

#### Type
_#clientCommand: {
	$dagger: task: _name: "ClientCommand"

	// Name of the command to execute
	// Examples: "ls", "/bin/bash"
	name: string

	// Positional arguments to the command
	// Examples: ["/tmp"]
	args: [...string]

	// Command-line flags represented in a civilized form
	// Example: {"-l": true, "-c": "echo hello world"}
	flags: [string]: bool | string

	// Environment variables
	// Example: {"DEBUG": "1"}
	env: [string]: string | #Secret

	// Capture standard output (as a string or secret)
	stdout?: {
		@dagger(generated)
		*string | #Secret
	}

	// Capture standard error (as a string or secret)
	stderr?: {
		@dagger(generated)
		*string | #Secret
	}

	// Inject standard input (from a string or secret)
	stdin?: {
		@dagger(generated)
		string | #Secret
	}
}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			p, err := plan.New(tt.root, tt.file)
			assert.Nil(t, err)

			doc, err := p.GetDocDefinition(tt.file, tt.def.line, tt.def.char)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, doc.String())
		})
	}
}
