package interpreter

import (
	"fmt"
	"math"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
)

// Abstract interpreter
type AbstractInterpreter interface {
	Interpret(ast ast.AST) (ast.Data, error)
}

type Interpreter struct{}

func (i *Interpreter) Interpret(node ast.AST) (ast.Data, error) {
	//visitor pattern
	switch v := node.(type) {
	case ast.Data:
		return v, nil
	case *ast.BinaryOperator_AST:
		if left, err := i.Interpret(v.Left); err != nil {
			return nil, err
		} else if right, err := i.Interpret(v.Right); err != nil {
			return nil, err
		} else if left == nil {
			return nil, fmt.Errorf("left node of binary operation is nil")
		} else if right == nil {
			return nil, fmt.Errorf("right node of binary operation is nil")
		} else if left_val, err := left.ValueAsFloat(); err != nil {
			return nil, err
		} else if right_val, err := right.ValueAsFloat(); err != nil {
			return nil, err
		} else {
			result := 0.0
			switch v.Operator {
			case "+":
				result = left_val + right_val
			case "-":
				result = left_val - right_val
			case "*":
				result = left_val * right_val
			case "/":
				result = left_val / right_val
			}
			if result-math.Trunc(result) != 0 {
				return ast.NewFloatData(result), nil
			}
			return ast.NewIntegerData(int(result)), nil
		}
	default:
		return nil, fmt.Errorf("invalid ast for interpretation: %v", node)
	}
}
