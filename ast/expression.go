package ast

// Expression A type of node that returns a value
type Expression interface {
	Node
	expressionNode()
}
