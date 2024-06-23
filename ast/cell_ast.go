package ast

import (
	"strconv"
)

type Cell_AST struct {
	Column string
	Row    int
}

func (ast *Cell_AST) Children() *[]AST {
	return &empty_children
}

func NewCellAST(value string) (*Cell_AST, error) {
	column := ""
	row := ""
	for i := 0; i < len(value); i++ {
		if (value[i] >= 'a' && value[i] <= 'z') || (value[i] >= 'A' && value[i] <= 'Z') {
			column += string(value[i])
		} else {
			row += string(value[i])
		}
	}
	if r, err := strconv.Atoi(row); err != nil {
		return nil, err
	} else {
		return &Cell_AST{Column: column, Row: r}, nil
	}
}

type CellData struct {
	value string
}

func NewCellData(value string) CellData {
	return CellData{
		value: value,
	}
}
func (ast *CellData) Value() any {
	return ast.value
}
func (ast *CellData) ValueAsInt() (int, error) {
	return strconv.Atoi(ast.value)
}
func (ast *CellData) ValueAsFloat() (float64, error) {
	return strconv.ParseFloat(ast.value, 64)
}
func (ast *CellData) ValueAsString() (string, error) {
	return ast.value, nil
}
func (ast *CellData) ValueAsBoolean() (bool, error) {
	return ast.value == "True", nil
}
