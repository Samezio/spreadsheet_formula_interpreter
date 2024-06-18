package main

import (
	"fmt"

	"github.com/samezio/spreadsheet_formula_interpreter/ast"
)

func main() {
	expression := "Hello World && Lexer"
	lexer := NewLexer(expression)
	lex := lexer()
	for lex != nil {
		fmt.Println(lex)
		lex = lexer()
	}

	//Test 2
	expression = "1+2+40.2*2/4-9"
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
	}

	//Test 3
	expression = "1+(2+40.2)*2/(4-9)"
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
	}
}
