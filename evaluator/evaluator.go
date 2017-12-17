package evaluator

import (
	"github.com/CzarSimon/monkey/ast"
	"github.com/CzarSimon/monkey/object"
)

var (
	NULL  = object.NewNull()
	TRUE  = object.NewBoolean(true)
	FALSE = object.NewBoolean(false)
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
		return object.NewInteger(node.Value)
	case *ast.Boolean:
		return nativeBoolTooBooleanObject(node.Value)
	case *ast.PrefixExpression:
		rightArg := Eval(node.Right)
		return evalPrefixExpression(node.Operator, rightArg)
	case *ast.InfixExpression:
		leftArg := Eval(node.Left)
		rightArg := Eval(node.Right)
		return evalInfixExpression(node.Operator, leftArg, rightArg)
	case *ast.BlockStatement:
		return evalStatements(node.Statements)
	case *ast.IFExpression:
		return evalIfExpression(node)
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

// evalPrefixExpression Evaluates a prefix expression to its resulting object
func evalPrefixExpression(operator string, argument object.Object) object.Object {
	switch operator {
	case "!":
		return evalNotOperatorExpression(argument)
	case "-":
		return evalMinusPrefixOperatorExpression(argument)
	default:
		return NULL
	}
}

// evalNotOperatorExpression Performs boolean negation on the
// the supplied argument and returns the result
func evalNotOperatorExpression(argument object.Object) object.Object {
	switch argument {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// evalMinusPrefixOperatorExpression Multiplies the supplied argument
// with -1 and returns the result
func evalMinusPrefixOperatorExpression(argument object.Object) object.Object {
	if argument.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := argument.(*object.Integer).Value
	return object.NewInteger(-value)
}

// evalInfixExpression Returns the result of performaing an operation
// denoted by the supplied operator on two argument objects
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolTooBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolTooBooleanObject(left != right)
	default:
		return NULL
	}
}

// evalIntegerInfixExpression Retruns the Integer result of performing an
// arithmetic inifx operation on the two supplied Integer arguments
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value
	switch operator {
	case "+":
		return object.NewInteger(leftValue + rightValue)
	case "-":
		return object.NewInteger(leftValue - rightValue)
	case "*":
		return object.NewInteger(leftValue * rightValue)
	case "/":
		return object.NewInteger(leftValue / rightValue)
	case "<":
		return nativeBoolTooBooleanObject(leftValue < rightValue)
	case ">":
		return nativeBoolTooBooleanObject(leftValue > rightValue)
	case "==":
		return nativeBoolTooBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBoolTooBooleanObject(leftValue != rightValue)
	default:
		return NULL
	}
}

// evalIfExpression Selects one branch of an IFExpression to evaluate
func evalIfExpression(ifExpr *ast.IFExpression) object.Object {
	condition := Eval(ifExpr.Condition)
	if isTruthy(condition) {
		return Eval(ifExpr.Consequence)
	} else if ifExpr.Alternative != nil {
		return Eval(ifExpr.Alternative)
	}
	return NULL
}

// isTruthy Checks if a supplied object is is truty
func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// nativeBoolTooBooleanObject Maps a boolean value to
// one of the boolean objects TRUE or FALSE
func nativeBoolTooBooleanObject(value bool) object.Object {
	if value {
		return TRUE
	}
	return FALSE
}
