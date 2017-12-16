package evaluator

import (
	"testing"

	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/object"
	"github.com/CzarSimon/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	return Eval(p.ParseProgram())
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	res, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj was not *object.Integer Got=%T", obj)
		return false
	}
	if res.Value != expected {
		t.Errorf("Wrong res.Value Expected=%d Got=%d", expected, res.Value)
		return false
	}
	return true
}
