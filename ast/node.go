package ast

// Node Interface of types in the ast
type Node interface {
	TokenLiteral() string
	String() string
}
