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