package codemodel

import (
	"testing"
)

var a string = `
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

func TestGenerableStruct(t *testing.T) {

}
