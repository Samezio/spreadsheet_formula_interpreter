package main

import (
	"fmt"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
	"github.com/Samezio/spreadsheet_formula_interpreter/interpreter"
	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

func main() {
	basic_interpreter := interpreter.Interpreter{}
	//Test 2
	expression := "4-1.5 * 10 + 10 / 100"
	if a, err := parser.Parse(expression); err != nil {
		panic(err)
	} else {
		fmt.Println(ast.AST_ToString(a, 0))
		if data, err := basic_interpreter.Interpret(a); err != nil {
			panic(err)
		} else {
			fmt.Printf("Result: %v\n", data.Value())

		}
	}

	//Test 3
	expression = "4-1.5*(10+10)/100"
	if a, err := parser.Parse(expression); err != nil {
		panic(err)
	} else {
		fmt.Println(ast.AST_ToString(a, 0))
		if data, err := basic_interpreter.Interpret(a); err != nil {
			panic(err)
		} else {
			fmt.Printf("Result: %v\n", data.Value())

		}
	}

	//Test 4
	expression = "4-1.5*(10+10)/100 == 3.9"
	if a, err := parser.Parse(expression); err != nil {
		panic(err)
	} else {
		fmt.Println(ast.AST_ToString(a, 0))
		if data, err := basic_interpreter.Interpret(a); err != nil {
			panic(err)
		} else {
			fmt.Printf("Result: %v\n", data.Value())

		}
	}
	//Test 5
	expression = "4-1.5 * 10 + 10 / 100 == -10.9"
	if a, err := parser.Parse(expression); err != nil {
		panic(err)
	} else {
		fmt.Println(ast.AST_ToString(a, 0))
		if data, err := basic_interpreter.Interpret(a); err != nil {
			panic(err)
		} else {
			fmt.Printf("Result: %v\n", data.Value())

		}
	}
	/*scanner := bufio.NewScanner(os.Stdin)

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
	}*/
}
