package object

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

// ObjectType String denoting the type of an object
type ObjectType string

// Object
type Object interface {
	Type() ObjectType
	Inspect() string
}
