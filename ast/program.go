package ast

import "bytes"

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

// TokenLiteral Retruns the node token literal
func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String Returns a string representation of the Program node
func (program *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range program.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}
