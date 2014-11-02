package main

import (
	"fmt"
	parser "github.com/daischio/daischeme/schemaparser"
)

func main() {
	// Parse a json schema
	scheme := parser.ParseSchema("./schema/example1.json")

	// Print the structure
	fmt.Printf("%+v", scheme)
}
