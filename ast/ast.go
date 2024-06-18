package ast

import (
	"fmt"
	"reflect"
	"strings"
)

type AST interface {
	Children() *[]AST
}
type Data interface {
	Value() any
	ValueAsInt() (int, error)
	ValueAsFloat() (float64, error)
	ValueAsString() (string, error)
}

var empty_children = []AST{}

func NewValueAST(value any) (AST, error) {
	switch v := value.(type) {
	case int:
		return &Integer_AST{
			value: v,
		}, nil
	case float64:
		return &Float_AST{
			value: v,
		}, nil
	case string:
		return &String_AST{
			value: v,
		}, nil
	default:
		return nil, fmt.Errorf("not matching ast for this data")
	}
}
func getType(myvar interface{}) string {
	valueOf := reflect.ValueOf(myvar)
	if valueOf.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(valueOf).Type().Name()
	}
	return valueOf.Type().Name()
}
func AST_ToString(ast AST, numTabs int) string {
	s := strings.Repeat("\t", numTabs) + getType(ast)
	switch v := ast.(type) {
	case Data:
		s += fmt.Sprintf("[%v]", v.Value())
	case *BinaryOperator_AST:
		s += fmt.Sprintf("[%v]", v.Operator)
	}
	s += "\n"
	if ast.Children() == &empty_children || ast.Children() == nil {
		return s
	}
	for i := 0; i < len(*ast.Children()); i++ {
		s += AST_ToString((*ast.Children())[i], numTabs+1)
	}
	return s
}
