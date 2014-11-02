package schemaparser

import (
	"fmt"
	"reflect"
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

func TestType(t *testing.T) {
	scheme := Parse(a)
	if scheme.Type != "object" {
		t.Error(scheme.Type, "should have been object")
	}
}

func TestSchema(t *testing.T) {
	scheme := Parse(a)
	if scheme.Schema != "http://json-schema.org/draft-03/schema" {
		t.Error(scheme.Schema, "not expected schema")
	}
}

func TestId(t *testing.T) {
	scheme := Parse(a)
	if scheme.Id != "http://jsonschema.net" {
		t.Error(scheme.Id, "not expected id")
	}
}

func TestRequired(t *testing.T) {
	scheme := Parse(a)
	if scheme.Required != false {
		t.Error("schema.Required should be false")
	}
}

func TestProperties(t *testing.T) {
	scheme := Parse(a)
	typeOf := reflect.TypeOf(scheme.Properties["firstName"])
	s := fmt.Sprintf("%v", typeOf)
	if s != "schemaparser.Scheme" {
		t.Error("Type of the key it's not schemaparser.Scheme")
	}
}
