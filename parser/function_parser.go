package parser

import (
	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
	"github.com/Samezio/spreadsheet_formula_interpreter/interpreter"
)

type FunctionCreator func(parameters *[]ast.AST) interpreter.Function
type FunctionRegistry map[string]FunctionCreator

func CreateDefaultFunctionRegistry() FunctionRegistry {
	functionRegistry := make(FunctionRegistry)
	functionRegistry["MAX"] = interpreter.NewMaxFunction
	return functionRegistry
}
