package main

import (
	"fmt"
	"github.com/daischio/daischeme/codemodel"
)

func main() {
	// Generate a model from a schema
	m := codemodel.New("MyPackage", "MyModel", "./assets/json_example_schema.json")
	fmt.Printf("%+v\n", m)

	// Generate the SchemaStore from a schema
	schemaStore := m.GetSchemeStore()
	fmt.Printf("main: %+v\n", *schemaStore.GetSchema())

	schemas := m.GetSchemeStore().GetSchemas()
	fmt.Printf("\nSchemas: %+v\n", schemas)

}
