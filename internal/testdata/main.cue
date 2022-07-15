package main

// An example of a definition
#Simple: string

// Multi line
// documentation
#SimpleMultiLineDoc: number

// Structure
#Struct: {
	bar: string

	foo: *true | bool
}

// Struct with doc in
// multiline
#StructWithDoc: {
	// first field
	bar: string

	// second field
	foo: number | [...number]
}

// Multiline documentation
// with reference to a definition
#ReferenceToStruct: {
	// Multi line field doc
	// reference to a def
	ref: #Struct

	enum: "hello" | "world" | string
}

// Dagger type
// A reference to an external secret, for example:
//  - A password
//  - A SSH private key
//  - An API token
// Secrets are never merged in the Cue tree. They can only be used
// by a special filesystem mount designed to minimize leak risk.
#Secret: {
	$dagger: secret: _id: string
}

// Complex command
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
}