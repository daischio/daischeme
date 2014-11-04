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
	schemaStore := m.GetSchemaStore()
	fmt.Printf("main: %+v\n", *schemaStore.GetSchema())

	schemas := m.GetSchemaStore().GetSchemas()
	fmt.Printf("\nSchemas: %+v\n", schemas)

	structs := m.GetSchemaMapper().GetMappedStructs()
	fmt.Printf("\nStructs: %+v\n", structs)

}
