package ast

type Boolean_AST struct {
	value bool
}

func (ast *Boolean_AST) Value() any {
	return ast.value
}
func (ast *Boolean_AST) Children() *[]AST {
	return &empty_children
}
func (ast *Boolean_AST) ValueAsInt() (int, error) {
	if ast.value {
		return 1, nil
	} else {
		return 0, nil
	}
}
func (ast *Boolean_AST) ValueAsFloat() (float64, error) {
	if ast.value {
		return 1.0, nil
	} else {
		return 0, nil
	}
}
func (ast *Boolean_AST) ValueAsString() (string, error) {
	if ast.value {
		return "TRUE", nil
	} else {
		return "FALSE", nil
	}
}
func (ast *Boolean_AST) ValueAsBoolean() (bool, error) {
	return ast.value, nil
}

func NewBooleanData(value bool) Data {
	return &Boolean_AST{value: value}
}
func NewBooleanAST(value bool) AST {
	return &Boolean_AST{value: value}
}
