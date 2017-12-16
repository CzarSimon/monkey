package object

// Null Object representing a nothing value
type Null struct{}

func (null *Null) Type() ObjectType { return NULL_OBJ }

func (null *Null) Inspect() string { return "null" }

// NewNull Creates a new Null object and returns a reference to it
func NewNull() *Null {
	return &Null{}
}
