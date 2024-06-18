package ast

type BinaryOperator_AST struct {
	Operator string
	Left     AST
	Right    AST
}

func (ast *BinaryOperator_AST) Children() *[]AST {
	return &[]AST{ast.Left, ast.Right}
}
