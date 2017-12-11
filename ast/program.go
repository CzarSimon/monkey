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

// AddStatements Adds statements to the program
func (program *Program) AddStatements(stmts ...Statement) {
	program.Statements = append(program.Statements, stmts...)
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
