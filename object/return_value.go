package object

// ReturnValue Wrapper object for a value retured by a return statement
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// NewReturnValue Creates a new ReturnValue and returns a referece to it
func NewReturnValue(value Object) *ReturnValue {
	return &ReturnValue{
		Value: value,
	}
}
