package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
	"github.com/Samezio/spreadsheet_formula_interpreter/interpreter"
	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	basic_interpreter := interpreter.Interpreter{}
	retriveCelldata := func(column string, row int) (ast.CellData, error) {
		return ast.NewCellData("1"), nil
	}
	fmt.Print(">>")
	for scanner.Scan() {
		exp := scanner.Text()
		if exp == "exit" {
			break
		}
		if a, err := parser.Parse(exp); err != nil {
			fmt.Printf("Error occured: %v\n", err)
		} else if data, err := basic_interpreter.Interpret(a, retriveCelldata); err != nil {
			fmt.Printf("Error occured: %v\n", err)
		} else {
			fmt.Printf("Result: %v\n", data.Value())
		}
		fmt.Print(">>")
	}
}
