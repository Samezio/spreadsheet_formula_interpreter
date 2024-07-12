package ast

import (
	"fmt"
	"strconv"
	"strings"
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
func (ast *String_AST) ValueAsBoolean() (bool, error) {
	t := strings.ToUpper(ast.value)
	if t == "TRUE" {
		return true, nil
	} else if t == "FALSE" {
		return false, nil
	} else {
		return false, fmt.Errorf("can't be converted to boolean")
	}
}

func NewStringData(value string) Data {
	return &String_AST{value: value}
}
func NewStringAST(value string) AST {
	return &String_AST{value: value}
}
