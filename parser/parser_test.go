package parser

import (
	"reflect"
	"testing"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
)

func TestParsedAST(t *testing.T) {
	expression := " 123 & \"abc\" == 100+23&'abc'"
	var expected_ast ast.AST = &ast.BinaryOperator_AST{
		Operator: "==",
		Left: &ast.BinaryOperator_AST{
			Operator: "&",
			Left:     ast.NewIntegerAST(123),
			Right:    ast.NewStringAST("abc"),
		},
		Right: &ast.BinaryOperator_AST{
			Operator: "&",
			Left: &ast.BinaryOperator_AST{
				Operator: "+",
				Left:     ast.NewIntegerAST(100),
				Right:    ast.NewIntegerAST(23),
			},
			Right: ast.NewStringAST("abc"),
		},
	}
	if a, err := Parse(expression); err != nil {
		t.Errorf("Error occured while parsing: %v", err)
	} else if !reflect.DeepEqual(a, expected_ast) {
		t.Errorf("Assertion failed. Expected:\n%v\nGot: %v", ast.AST_ToString(expected_ast, 0), ast.AST_ToString(a, 0))
	}
}
