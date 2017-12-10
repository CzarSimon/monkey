package ast

// Statement A type of node that does not return a value
type Statement interface {
	Node
	statementNode()
}
