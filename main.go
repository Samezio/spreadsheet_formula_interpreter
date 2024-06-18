package main

import (
	"fmt"

	"github.com/samezio/spreadsheet_formula_interpreter/ast"
	"github.com/samezio/spreadsheet_formula_interpreter/interpreter"
)

func main() {
	basic_interpreter := interpreter.Interpreter{}
	expression := "Hello World && Lexer"
	lexer := NewLexer(expression)
	lex := lexer()
	for lex != nil {
		fmt.Println(lex)
		lex = lexer()
	}

	//Test 2
	expression = "4-1.5 * 10 + 10 / 100"
	lexer = NewLexer(expression)
	lex = lexer()
	for lex != nil {
		fmt.Println(lex)
		lex = lexer()
	}
	if a, err := Parse(expression); err != nil {
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
	lexer = NewLexer(expression)
	lex = lexer()
	for lex != nil {
		fmt.Println(lex)
		lex = lexer()
	}
	if a, err := Parse(expression); err != nil {
		panic(err)
	} else {
		fmt.Println(ast.AST_ToString(a, 0))
		if data, err := basic_interpreter.Interpret(a); err != nil {
			panic(err)
		} else {
			fmt.Printf("Result: %v\n", data.Value())

		}
	}
}
