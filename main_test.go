package main

import (
	"testing"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
	"github.com/Samezio/spreadsheet_formula_interpreter/interpreter"
	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

var basic_interpreter = interpreter.Interpreter{}

var retriveCelldata = func(column string, row int) (ast.CellData, error) {
	return ast.NewCellData("1"), nil
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestArithmaticExpression(t *testing.T) {
	var tests = []struct {
		expression  string
		result_type string
		want        interface{}
	}{
		{"1*10-9/3+-1", "integer", 1*10 - 9/3 + -1},
		{"1.2*10.8-9/3+1", "float", 1.2*10.8 - 9/3 + 1},
		{"1*(10-9)/4+-1", "float", 1*(10.0-9)/4 + -1},
		{"1>10", "boolean", 1 > 10},
		{"1<10", "boolean", 1 < 10},
		{"1<=1", "boolean", 1 <= 1},
		{"1==10", "boolean", 1 == 10},
		{"1!=10", "boolean", 1 != 10},
		{"'Hello'&' '&'World'", "string", "Hello" + " " + "World"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			if a, err := parser.Parse(tt.expression); err != nil {
				t.Errorf("Error occured[%s]: %v\n", tt.expression, err)
			} else if data, err := basic_interpreter.Interpret(a, retriveCelldata); err != nil {
				t.Errorf("Error occured[%s]: %v, %v\n", tt.expression, err, ast.AST_ToString(a, 0))
			} else {
				switch tt.result_type {
				case "integer":
					if result, err := data.ValueAsInt(); err != nil {
						t.Errorf("Error occured, expected int value[%s]: %v\n", tt.expression, err)
					} else if tt.want != result {
						t.Errorf("Result not correct. %s should give %v, but got %v", tt.expression, tt.want, result)
					}
				case "float":
					if result, err := data.ValueAsFloat(); err != nil {
						t.Errorf("Error occured, expected float value: %v\n", err)
					} else if tt.want != result {
						t.Errorf("Result not correct. %s should give %v, but got %v", tt.expression, tt.want, result)
					}
				case "string":
					if result, err := data.ValueAsString(); err != nil {
						t.Errorf("Error occured, expected string value: %v\n", err)
					} else if tt.want != result {
						t.Errorf("Result not correct. %s should give %v, but got %v", tt.expression, tt.want, result)
					}
				case "boolean":
					if result, err := data.ValueAsBoolean(); err != nil {
						t.Errorf("Error occured, expected boolean value: %v\n", err)
					} else if tt.want != result {
						t.Errorf("Result not correct. %s should give %v, but got %v", tt.expression, tt.want, result)
					}
				default:
					t.Fatalf("Invalid result type in test case")
				}
			}
		})
	}
}
