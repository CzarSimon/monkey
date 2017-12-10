package ast

// Program Slice of statements representing a full program
type Program struct {
	Statements []Statement
}

// NewProgram Creates a new program structure
func NewProgram() *Program {
	return &Program{
		Statements: make([]Statement, 0),
	}
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
