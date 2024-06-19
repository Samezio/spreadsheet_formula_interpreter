package ast

import "fmt"

type Integer_AST struct {
	value int
}

func (ast *Integer_AST) Value() any {
	return ast.value
}
func (ast *Integer_AST) Children() *[]AST {
	return &empty_children
}
func (ast *Integer_AST) ValueAsInt() (int, error) {
	return ast.value, nil
}
func (ast *Integer_AST) ValueAsFloat() (float64, error) {
	return float64(ast.value), nil
}
func (ast *Integer_AST) ValueAsString() (string, error) {
	return fmt.Sprintf("%d", ast.value), nil
}
func (ast *Integer_AST) ValueAsBoolean() (bool, error) {
	return ast.value > 0, nil
}

func NewIntegerData(value int) Data {
	return &Integer_AST{value: value}
}
