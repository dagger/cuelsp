package internal

import "strings"

// Definition is a simple abstraction of a CUE definition
// to simplify management of each part of the definition
type Definition struct {
	isImported bool
	pkg        string
	def        string
}

// StringToDef convert a string into definition
func StringToDef(def string) *Definition {
	isImported := strings.Contains(def, ".")

	if !isImported {
		return &Definition{
			isImported: isImported,
			def:        def,
		}
	}

	contents := strings.Split(def, ".")
	return &Definition{
		isImported: isImported,
		pkg:        contents[0],
		def:        contents[1],
	}
}

func (d *Definition) IsImported() bool {
	return d.isImported
}

func (d *Definition) Pkg() string {
	return d.pkg
}

func (d *Definition) Def() string {
	return d.def
}
