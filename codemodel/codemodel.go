package codemodel

import (
	parser "github.com/daischio/daischeme/codemodel/schemaparser"
	store "github.com/daischio/daischeme/codemodel/schemastore"
)

type CodeModel struct {
	PackageName string
	ModelName   string
	Scheme      *parser.Scheme
}

func (c *CodeModel) GetSchemeStore() *store.SchemaStore {
	s := store.New(c.Scheme)
	return s
}

/* Generate a new CodeModel instance */
func New(p string, md string, schema string) *CodeModel {
	model := CodeModel{p, md, FromSchemaFile(schema)}
	return &model
}

/* Generate a CodeModel for the input schema file */
func FromSchemaFile(p string) *parser.Scheme {
	scheme := parser.ParseSchema(p)
    return scheme
}
