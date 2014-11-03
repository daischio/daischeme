package codegenerator

import (
	"fmt"
	"github.com/daischio/daischeme/schemaparser"
)

type CodeGenerator interface {
	GenerateHead() string
	GenerateField() string
	GenerateMethod() string
}

type Model struct {
	PackageName string
	ModelName   string
	Scheme      *schemaparser.Scheme
}

func New(p string, md string, s *schemaparser.Scheme) *Model {
	model := Model{p, md, s}
	return &model
}

func main() {
	a := `
    {
      "type":"object",
      "$schema": "http://json-schema.org/draft-03/schema",
      "id": "http://jsonschema.net",
      "required":false,
      "properties":{
        "firstName": {
          "type":"string",
          "id": "http://jsonschema.net/firstName",
          "required":false
        },
        "name": {
          "type":"string",
          "id": "http://jsonschema.net/name",
          "required":false
        }
      }
    }`

	// Get the schema object
	scheme := schemaparser.Parse(a)
	m := New("models", "Testmodel", scheme)
	fmt.Println("%v", m)

}
