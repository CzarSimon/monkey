package evaluator

import (
	"testing"

	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/object"
	"github.com/CzarSimon/monkey/parser"
)

type testStruct struct {
	input    string
	expected interface{}
}

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
	env := object.NewEnvironment()
	return Eval(p.ParseProgram(), env)
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
	tests := []testStruct{
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
		expected := test.expected.(bool)
		testBooleanObject(t, evaluated, expected)
	}
}

func TestNotOperator(t *testing.T) {
	tests := []testStruct{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!4", true},
	}

	for _, test := range tests {
		evaluated := testEval(test.input)
		expected := test.expected.(bool)
		testBooleanObject(t, evaluated, expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []testStruct{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		integer, ok := test.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
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

func TestReturnStatement(t *testing.T) {
	tests := []testStruct{
		{"return 10;", 10},
		{"return 11; 9;", 11},
		{"return 2 * 6; 9;", 12},
		{"9; return 3 + 2 * 5; 10;", 13},
		{`
      if (1 > 0) {
        return 10;
      }
      return 1;
      `, 10},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		integer := test.expected.(int)
		testIntegerObject(t, evaluated, int64(integer))
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL Got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestErrorObject(t *testing.T) {
	tests := []testStruct{
		{
			"5 + true",
			"Type missmatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"Type missmatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"Unknown operator: -BOOLEAN",
		},
		{
			"false + true",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + true; 5;",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { false + true; }",
			"Unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`if (10 > 1) {
			   if (10 > 1) {
					 return true - false;
				 }
				 return true;
			 } `,
			"Unknown operator: BOOLEAN - BOOLEAN",
		},
		{
			"foobar",
			"Identifier not found: foobar",
		},
	}
	for i, test := range tests {
		evaluated := testEval(test.input)
		err, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("%d. - Expected type *object.Error Got=%T(%+v)",
				i, evaluated, evaluated)
		}
		expectedMsg := test.expected.(string)
		if err.Message != expectedMsg {
			t.Errorf("%d. - Wrong error message: Expected=%s Got=%s",
				i, expectedMsg, err.Message)
		}
	}
}

func TestLetStatement(t *testing.T) {
	tests := []testStruct{
		{"let a = 5; a;", 5},
		{"let a = 5 * 2; a;", 10},
		{"let a = -5; let b = a; b;", -5},
		{"let a = -5; let b = a; let c = a + b + 15; c;", 5},
	}
	for i, test := range tests {
		evaluated := testEval(test.input)
		expectedInt := test.expected.(int)
		if !testIntegerObject(t, evaluated, int64(expectedInt)) {
			t.Errorf("%d. - Test failed", i)
		}
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; }"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("Object is not a function Got=%T (%+v)", evaluated, evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Fatalf("Wrong number of parameters Expected=1 Got=%d", len(fn.Parameters))
	}
	if fn.Parameters[0].String() != "x" {
		t.Fatalf("Parameter is not 'x' Got='%s'", fn.Parameters[0].String())
	}
	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("Wrong body Expected=%s Got=%s", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []testStruct{
		{"let I = fn(x) { x; }; I(5);", 5},
		{"let I = fn(x) { return x; }; I(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5)", 10},
		{"let add = fn(x, y) { x + y; }; add(2, 8)", 10},
		{"let add = fn(x, y) { x + y; }; add(1 + 1, add(2, 4))", 8},
		{"fn(x) { x; }(5)", 5},
	}
	for i, test := range tests {
		evaluated := testEval(test.input)
		expectedInt := test.expected.(int)
		if !testIntegerObject(t, evaluated, int64(expectedInt)) {
			t.Errorf("%d. - Test failed", i)
		}
	}
}
