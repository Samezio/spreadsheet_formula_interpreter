package interpreter_test

import (
	"testing"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
	"github.com/Samezio/spreadsheet_formula_interpreter/interpreter"
	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

func TestArithmaticExpression(t *testing.T) {
	retriveCelldata := func(column string, row int) (ast.CellData, error) {
		return ast.NewCellData("1"), nil
	}
	basic_interpreter := interpreter.Interpreter{
		RetriveCelldata: retriveCelldata,
	}
	function_registry := parser.CreateDefaultFunctionRegistry()

	expression := "1+1"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 2 {
		t.Errorf("Expected 2 but got %v", result)
	}

	expression = "1-1"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 0 {
		t.Errorf("Expected 0 but got %v", result)
	}

	expression = "1.09*100"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 109 {
		t.Errorf("Expected 109 but got %v", result)
	}

	expression = "1-1*100"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != -99 {
		t.Errorf("Expected -99 but got %v", result)
	}

	expression = "10*100-99"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 901 {
		t.Errorf("Expected 901 but got %v", result)
	}

	expression = "1*(100-99)"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 1 {
		t.Errorf("Expected 1 but got %v", result)
	}

	expression = "10+-8"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsFloat(); err != nil {
		t.Errorf("Failed convert to float: %v", err)
	} else if value != 2 {
		t.Errorf("Expected 0.2 but got %v", result)
	}
}

func TestEqulityExpression(t *testing.T) {
	retriveCelldata := func(column string, row int) (ast.CellData, error) {
		return ast.NewCellData("1"), nil
	}
	basic_interpreter := interpreter.Interpreter{
		RetriveCelldata: retriveCelldata,
	}
	function_registry := parser.CreateDefaultFunctionRegistry()

	expression := "1==1"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 1 {
		t.Errorf("Expected 1 but got %v", result)
	} else if value, err := result.ValueAsBoolean(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if !value {
		t.Errorf("Expected true but got %v", result)
	}

	expression = "1.1>1"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 1 {
		t.Errorf("Expected 1 but got %v", result)
	} else if value, err := result.ValueAsBoolean(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if !value {
		t.Errorf("Expected true but got %v", result)
	}

	expression = "1>=1"
	if a, err := parser.Parse(expression, function_registry); err != nil {
		t.Errorf("Parsing error: %v", err)
	} else if result, err := basic_interpreter.Interpret(a); err != nil {
		t.Errorf("Interpreter error: %v", err)
	} else if value, err := result.ValueAsInt(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if value != 1 {
		t.Errorf("Expected 1 but got %v", result)
	} else if value, err := result.ValueAsBoolean(); err != nil {
		t.Errorf("Failed convert to int: %v", err)
	} else if !value {
		t.Errorf("Expected true but got %v", result)
	}
}
