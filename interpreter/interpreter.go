package interpreter

import (
	"fmt"
	"math"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
)

type RetriveCelldata func(string, int) (ast.CellData, error)

// Abstract interpreter
type AbstractInterpreter interface {
	Interpret(ast ast.AST) (ast.Data, error)
}

type Interpreter struct {
	RetriveCelldata RetriveCelldata
}

func (i *Interpreter) Interpret(node ast.AST) (ast.Data, error) {
	//visitor pattern
	switch v := node.(type) {
	case ast.Data:
		return v, nil
	case *ast.UnaryOperator_AST:
		if operand, err := i.Interpret(v.Operand); err != nil {
			return nil, err
		} else if operand == nil {
			return nil, fmt.Errorf("operand node of unary operation is nil")
		} else if value, err := operand.ValueAsFloat(); err != nil {
			return nil, err
		} else {
			value = value * (-1)
			if value-math.Trunc(value) != 0 {
				return ast.NewFloatData(value), nil
			}
			return ast.NewIntegerData(int(value)), nil
		}
	case *ast.Cell_AST:
		if cell_data, err := i.RetriveCelldata(v.Column, v.Row); err != nil {
			return nil, err
		} else {
			return &cell_data, nil
		}
	case *ast.BinaryOperator_AST:
		if left, err := i.Interpret(v.Left); err != nil {
			return nil, err
		} else if right, err := i.Interpret(v.Right); err != nil {
			return nil, err
		} else if left == nil {
			return nil, fmt.Errorf("left node of binary operation is nil")
		} else if right == nil {
			return nil, fmt.Errorf("right node of binary operation is nil")
		} else {
			result := 0.0
			switch v.Operator {
			case "&":
				if left_val, err := left.ValueAsString(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsString(); err != nil {
					return nil, err
				} else {
					return ast.NewStringData(left_val + right_val), nil
				}
			case "+":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					result = left_val + right_val
				}
			case "-":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					result = left_val - right_val
				}
			case "*":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					result = left_val * right_val
				}
			case "/":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					result = left_val / right_val
				}
			case "==":
				if left_val, err := left.ValueAsFloat(); err == nil {
					if right_val, err := right.ValueAsFloat(); err == nil {
						return ast.NewBooleanData(left_val == right_val), nil
					}
				}
				if left_val, err := left.ValueAsString(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsString(); err != nil {
					return nil, err
				} else {
					return ast.NewBooleanData(left_val == right_val), nil
				}
			case ">=":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					return ast.NewBooleanData(left_val >= right_val), nil
				}
			case "<=":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					return ast.NewBooleanData(left_val <= right_val), nil
				}
			case "!=":
				if left_val, err := left.ValueAsFloat(); err == nil {
					if right_val, err := right.ValueAsFloat(); err == nil {
						return ast.NewBooleanData(left_val != right_val), nil
					}
				}
				if left_val, err := left.ValueAsString(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsString(); err != nil {
					return nil, err
				} else {
					return ast.NewBooleanData(left_val != right_val), nil
				}
			case ">":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					return ast.NewBooleanData(left_val > right_val), nil
				}
			case "<":
				if left_val, err := left.ValueAsFloat(); err != nil {
					return nil, err
				} else if right_val, err := right.ValueAsFloat(); err != nil {
					return nil, err
				} else {
					return ast.NewBooleanData(left_val < right_val), nil
				}
			}
			if result-math.Trunc(result) != 0 {
				return ast.NewFloatData(result), nil
			}
			return ast.NewIntegerData(int(result)), nil
		}
	case Function:
		return v.Solve(i)
	default:
		return nil, fmt.Errorf("invalid ast for interpretation: %v", node)
	}
}
