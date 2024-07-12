package interpreter

import "github.com/Samezio/spreadsheet_formula_interpreter/ast"

type Function interface {
	Name() string
	Solve(i AbstractInterpreter) (ast.Data, error)
	Update(i AbstractInterpreter, updated_arg_index int) (ast.Data, error)
	Parameters() *[]ast.AST
	Children() *[]ast.AST
}

type Max_Function struct {
	parameters *[]ast.AST
}

func (f *Max_Function) Name() string {
	return "MAX"
}
func (f *Max_Function) Solve(i AbstractInterpreter) (ast.Data, error) {
	var max *float64 = nil
	for _, parameter := range *f.Parameters() {
		if r, err := i.Interpret(parameter); err != nil {
			return nil, err
		} else if value, err := r.ValueAsFloat(); err != nil {
			return nil, err
		} else if max == nil || *max < value {
			max = &value
		}
	}
	return ast.NewFloatData(*max), nil
}
func (f *Max_Function) Update(i AbstractInterpreter, updated_arg_index int) (ast.Data, error) {
	return f.Solve(i)
}
func (f *Max_Function) Parameters() *[]ast.AST {
	return f.parameters
}
func (f *Max_Function) Children() *[]ast.AST {
	return f.parameters
}

func NewMaxFunction(parameters *[]ast.AST) Function {
	return &Max_Function{
		parameters: parameters,
	}
}
