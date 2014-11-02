package main

import (
	"fmt"
	"github.com/daischio/daischeme/parser"
)

func main() {
	// Parse a json schema
	scheme := parser.ParseSchema("./schema/example1.json")

	// Print the structure
	fmt.Printf("%+v", scheme)
}
