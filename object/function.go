package object

import (
	"bytes"
	"strings"

	"github.com/CzarSimon/monkey/ast"
)

// Function Object wrapping a function
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// NewFunction Creates a new function based on a function literal and
// an environment and passes a reference to it
func NewFunction(fnLit *ast.FunctionLiteral, env *Environment) *Function {
	return &Function{
		Parameters: fnLit.Parameters,
		Body:       fnLit.Body,
		Env:        env,
	}
}

func (fn *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (fn *Function) Inspect() string {
	var out bytes.Buffer
	params := make([]string, 0)
	for _, param := range fn.Parameters {
		params = append(params, param.String())
	}
	out.WriteString("fn (")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(fn.Body.String())
	out.WriteString("\n}")
	return out.String()
}
