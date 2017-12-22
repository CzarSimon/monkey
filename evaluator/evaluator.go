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
		return evalProgram(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return object.NewInteger(node.Value)
	case *ast.Boolean:
		return nativeBoolTooBooleanObject(node.Value)
	case *ast.PrefixExpression:
		rightArg := Eval(node.Right)
		if isError(rightArg) {
			return rightArg
		}
		return evalPrefixExpression(node.Operator, rightArg)
	case *ast.InfixExpression:
		leftArg := Eval(node.Left)
		if isError(leftArg) {
			return leftArg
		}
		rightArg := Eval(node.Right)
		if isError(rightArg) {
			return rightArg
		}
		return evalInfixExpression(node.Operator, leftArg, rightArg)
	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements)
	case *ast.IFExpression:
		return evalIfExpression(node)
	case *ast.ReturnStatement:
		value := Eval(node.ReturnValue)
		if isError(value) {
			return value
		}
		return object.NewReturnValue(value)
	}
	return nil
}

// evalProgram Evaluates a series of supplied statements and
// and returns a resulting object.Object
func evalProgram(statements []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range statements {
		result = Eval(stmt)
		switch res := result.(type) {
		case *object.ReturnValue:
			return res.Value
		case *object.Error:
			return res
		}
	}
	return result
}

// evalBlockStatement Evaluates a series of supplied statements
// and returns a result
func evalBlockStatement(statements []ast.Statement) object.Object {
	var result object.Object
	for _, stmt := range statements {
		result = Eval(stmt)
		if blockShouldReturn(result) {
			return result
		}
	}
	return result
}

func blockShouldReturn(result object.Object) bool {
	if result == nil {
		return false
	}
	return (result.Type() == object.RETURN_VALUE_OBJ || result.Type() == object.ERROR_OBJ)
}

// evalPrefixExpression Evaluates a prefix expression to its resulting object
func evalPrefixExpression(operator string, argument object.Object) object.Object {
	switch operator {
	case "!":
		return evalNotOperatorExpression(argument)
	case "-":
		return evalMinusPrefixOperatorExpression(argument)
	default:
		return object.NewErrorf("Unkown operator: %s%s", operator, argument.Type())
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
		return object.NewErrorf("Unknown operator: -%s", argument.Type())
	}
	value := argument.(*object.Integer).Value
	return object.NewInteger(-value)
}

// evalInfixExpression Returns the result of performaing an operation
// denoted by the supplied operator on two argument objects
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() != right.Type():
		return object.NewErrorf("Type missmatch: %s %s %s",
			left.Type(), operator, right.Type())
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolTooBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolTooBooleanObject(left != right)
	default:
		return object.NewErrorf("Unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
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
		return object.NewErrorf("Unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

// evalIfExpression Selects one branch of an IFExpression to evaluate
func evalIfExpression(ifExpr *ast.IFExpression) object.Object {
	condition := Eval(ifExpr.Condition)
	if isError(condition) {
		return condition
	}
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

// isError Checks if a given object is an object.Error
func isError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ERROR_OBJ
}
