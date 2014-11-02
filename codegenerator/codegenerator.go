package codegenerator

type Generable struct {
	Name string
	String[]
}

type CodeGenerator interface {
	GenerateHead() string
	GenerateField() string
	GenerateMethod() string
}

