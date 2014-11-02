package main

import (
	"./parser"
	"fmt"
)

func main() {
	// Parse a json schema
	scheme := parser.ParseSchema("./schema/example1.json")

	// Print the structure
	fmt.Printf("%+v", scheme)
}
