package main

import (
	"fmt"
	model "github.com/daischio/daischeme/codegenerator"
	parser "github.com/daischio/daischeme/schemaparser"
)

func main() {
	// Parse a json schema
	scheme := parser.ParseSchema("./schema/example1.json")

	// Print the structure
	fmt.Printf("%+v\n\n", scheme)

	// Generate a model
	m := model.New("models", "Testmodel", scheme)
	fmt.Println("%+v", m)

	fmt.Println(model.GenCode(m))
}
