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

func (m *Model) GenerateHead() string {
	h := fmt.Sprintf("package %s\n\nimport()\n\n", m.PackageName)
	return h
}

func (m *Model) GenerateField() string {
	return ""
}

func (m *Model) GenerateMethod() string {
	return ""
}

func GenCode(m CodeGenerator) string {
	return m.GenerateHead()
}
