package schemamapper

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
)

type SchemaMapper struct {
	schema *parser.Scheme
}

type Field struct {
	Tag string
	Name string
	Type string
}

type Struct struct {
	Name string
	Fields []*Field
}

func New(p *parser.Scheme) *SchemaMapper{
	sm := &SchemaMapper{p}
	return sm
}

func (sm *SchemaMapper) GetMappedStructs() []*Struct{
  return make([]*Struct, 0)
}
