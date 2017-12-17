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
		{"-5", -5},
		{"-10", -10},
		{"--11", 11},
		{"-0", 0},
		{"5 + 5 + 5 - 10 + 1", 6},
		{"2 * 2 * 2 * 2", 16},
		{"2 * 3 / 3 + 4 - 5", 1},
		{"2 * 4 + 5", 13},
		{"2 * (4 + 5)", 18},
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

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"false == false", true},
		{"true == false ", false},
		{"true != false", true},
		{"false != true", true},
		{"true == true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func TestNotOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!4", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	res, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj was not *object.Boolean Got=%T", obj)
		return false
	}
	if res.Value != expected {
		t.Errorf("Wrong res.Value Expected=%d Got=%d", expected, res.Value)
		return false
	}
	return true
}
