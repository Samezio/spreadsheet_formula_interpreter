package ast

import (
	"strconv"
)

type String_AST struct {
	value string
}

func (ast *String_AST) Value() any {
	return ast.value
}
func (ast *String_AST) Children() *[]AST {
	return &empty_children
}
func (ast *String_AST) ValueAsInt() (int, error) {
	return strconv.Atoi(ast.value)
}
func (ast *String_AST) ValueAsFloat() (float64, error) {
	return strconv.ParseFloat(ast.value, 64)
}
func (ast *String_AST) ValueAsString() (string, error) {
	return ast.value, nil
}
