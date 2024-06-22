package ast

type UnaryOperator_AST struct {
	Operator string
	Operand  AST
}

func (ast *UnaryOperator_AST) Children() *[]AST {
	return &[]AST{ast.Operand}
}
