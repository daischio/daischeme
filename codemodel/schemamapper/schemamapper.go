package schemamapper

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	store "github.com/daischio/daischeme/codemodel/schemastore"
	"fmt"
	"unicode"
)

type SchemaMapper struct {
	schema *parser.Scheme
}

func New(p *parser.Scheme) *SchemaMapper{
	sm := &SchemaMapper{p}
	return sm
}

func (sm *SchemaMapper) GetMappedStructs() []*Struct{

  	structs := make([]*Struct, 0)

	// Create channels
	struct_chan := make(chan *Struct)
	done := make(chan bool)

	// Process Concurrently
	go transformSchemaToStruct(sm, struct_chan, done)

	select {
	case s := <- struct_chan:
		structs = append(structs, s)
	case <- done:
		break
	}

	return structs
}

func transformSchemaToStruct(sm *SchemaMapper, c chan *Struct, done chan bool) {
	var schemas []*parser.Scheme = store.New(sm.schema).GetSchemas()
	for _, schema := range schemas {
		s := new(Struct)
		// Map schema properties to fields
		if schema.Properties == nil {
			continue
		}

		// Make first Cap
		capitalize := func(s string) string {
			a := []rune(s)
			a[0] = unicode.ToUpper(a[0])
			s = string(a)
			return s
		}

		// Walk over properties
		for k, v := range schema.Properties {
			f := new(Field)
			f.Name = capitalize(k)
			f.Tag = fmt.Sprintf(`json:"%s"`, k)
			f.Type = v.Type
			s.Fields = append(s.Fields, f)
		}

		//todo: Add a name for the struct
		s.Name = "DummyName"

		//@todo: Use the correct types
		// Add types

		// Return the struct
		c <- s
	}
	done <- true
}

/* Structs */

type Struct struct {
	Name string
	Fields []*Field
}

type Field struct {
	Tag string
	Name string
	Type string
}

func (s *Struct) head() string {
	return fmt.Sprintf("type %s struct {\n", s.Name)
}

func (s *Struct) field(i int) string {
	return fmt.Sprintf("\t%s %s \t\t\t`%s`\n", s.Fields[i].Name, s.Fields[i].Type, s.Fields[i].Tag)
}

func (s *Struct) tail() string {
	return fmt.Sprint("\n}")
}


func (s *Struct) Code() string {
	res := ""
	res += s.head()
	for i := range s.Fields {
		res += s.field(i)
	}
	res += s.tail()
	return res
}
