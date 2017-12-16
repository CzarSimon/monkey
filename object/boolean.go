package object

import "fmt"

// Boolean Object representing a boolean value
type Boolean struct {
	Value bool
}

func (boolean *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (boolean *Boolean) Inspect() string {
	return fmt.Sprintf("%t", boolean.Value)
}

// NewBoolean Creates a new Boolean object and returns a reference to it
func NewBoolean(value bool) *Boolean {
	return &Boolean{
		Value: value,
	}
}
