package object

import "fmt"

// Integer Object representing an integer value
type Integer struct {
	Value int64
}

func (integer *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (integer *Integer) Inspect() string {
	return fmt.Sprintf("%d", integer.Value)
}
