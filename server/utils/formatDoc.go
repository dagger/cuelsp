package utils

import (
	"fmt"

	"github.com/dagger/daggerlsp/loader"
)

// FormatDefinitionDoc will prettify the doc of a CUE value by providing any
// useful information
// For now, it is formatted in Markdown with the following pattern.
//
//  ##### Description
//   <Insert CUE value doc>
//  ##### Type
//   <Describe CUE value type>
func FormatDefinitionDoc(v *loader.Value) string {
	doc := "#### Description\n"

	for _, d := range v.Doc() {
		doc = fmt.Sprintf("%s%s\n", doc, d.Text())
	}

	if fieldsDoc, err := v.ListFieldDoc(); err == nil {
		if fieldsDoc > "" {
			doc += "#### Type  \n"
			doc += fieldsDoc
		}
	}

	return doc
}
