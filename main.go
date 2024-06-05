package main

import "fmt"

func main() {
	lexer := NewLexer("Hello World && Lexer")
	lex := lexer()
	for lex != nil {
		fmt.Println(lex)
		lex = lexer()
	}
}
