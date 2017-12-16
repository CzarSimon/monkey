package evaluator

import (
	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/object"
)

// Eval Evaluates a part of an AST from the supplied node downwards
// and returns a resulting object.Object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

// evalStatements Evaluates a series of supplied statements and
// and returns a resulting object.Object
func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range statements {
		result = Eval(stmt)
	}
	return result
}
