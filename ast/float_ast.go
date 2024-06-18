package ast

import "fmt"

type Float_AST struct {
	value float64
}

func (ast *Float_AST) Value() any {
	return ast.value
}
func (ast *Float_AST) Children() *[]AST {
	return &empty_children
}
func (ast *Float_AST) ValueAsInt() (int, error) {
	return int(ast.value), nil
}
func (ast *Float_AST) ValueAsFloat() (float64, error) {
	return ast.value, nil
}
func (ast *Float_AST) ValueAsString() (string, error) {
	return fmt.Sprintf("%f", ast.value), nil
}

func NewFloatData(value float64) Data {
	return &Float_AST{value: value}
}
