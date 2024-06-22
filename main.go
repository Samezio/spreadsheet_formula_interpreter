package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Samezio/spreadsheet_formula_interpreter/interpreter"
	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	basic_interpreter := interpreter.Interpreter{}
	fmt.Print(">>")
	for scanner.Scan() {
		exp := scanner.Text()
		if exp == "exit" {
			break
		}
		if a, err := parser.Parse(exp); err != nil {
			fmt.Println(err)
		} else if data, err := basic_interpreter.Interpret(a); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data.Value())
		}
		fmt.Print(">>")
	}
}
