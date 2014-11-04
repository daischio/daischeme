package schemastore

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	"fmt"
)

type SchemaStore struct {
	schema *parser.Scheme
}

func New(s *parser.Scheme) *SchemaStore {
	st := &SchemaStore{s};
	return st
}

func (s *SchemaStore) GetSchema() *parser.Scheme {
	return s.schema
}

func (s *SchemaStore) GetSchemas() []*parser.Scheme {
	// Create
	schemas := make([]*parser.Scheme, 0)

	// Add the first schema
	schemas = append(schemas, s.schema)

	// Define a walking function
	var walk func(s *parser.Scheme)
	walk = func(s *parser.Scheme) {
		// Iterate over properties and get nested schemas
		for k, v := range s.Properties {
			fmt.Printf("%v: %+v\n",k,v)
			if v.Type == "object" && v.Properties != nil {
				// append the schema
				schemas = append(schemas, &v)
				// walk that schema also
				walk(&v)
			}
		}
	}
	// Iterate
	walk(s.schema)
	return schemas
}
