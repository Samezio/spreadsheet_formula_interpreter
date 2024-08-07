package parser_test

import (
	"testing"

	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

func TestNumberToken(t *testing.T) {
	expression := "123"
	lexer := parser.NewLexer(expression)
	token := lexer()

	if token == nil {
		t.Fatalf("Expected a token but got nil")
	}

	if token.Type != parser.NUMBER_LEX {
		t.Errorf("Expected token type %s but got %s", parser.NUMBER_LEX, token.Type)
	}

	if token.Token != "123" {
		t.Errorf("Expected token value '123' but got '%s'", token.Token)
	}
}

func TestStringToken(t *testing.T) {
	expression := "abc"
	lexer := parser.NewLexer(expression)
	token := lexer()

	if token == nil {
		t.Fatalf("Expected a token but got nil")
	}

	if token.Type != parser.STRING_LEX {
		t.Errorf("Expected token type %s but got %s", parser.STRING_LEX, token.Type)
	}

	if token.Token != "abc" {
		t.Errorf("Expected token value 'abc' but got '%s'", token.Token)
	}
}

func TestSpecialToken(t *testing.T) {
	expression := "+-*/"
	lexer := parser.NewLexer(expression)

	for _, char := range expression {
		token := lexer()

		if token == nil {
			t.Fatalf("Expected a token but got nil")
		}

		if token.Type != parser.SPECIAL_LEX {
			t.Errorf("Expected token type %s but got %s", parser.SPECIAL_LEX, token.Type)
		}

		if token.Token != string(char) {
			t.Errorf("Expected token value '%s' but got '%s'", string(char), token.Token)
		}
	}
}

func TestMixedTokens(t *testing.T) {
	expression := "123+abc-456"
	lexer := parser.NewLexer(expression)

	expectedTokens := []parser.Lex{
		{Token: "123", Type: parser.NUMBER_LEX},
		{Token: "+", Type: parser.SPECIAL_LEX},
		{Token: "abc", Type: parser.STRING_LEX},
		{Token: "-", Type: parser.SPECIAL_LEX},
		{Token: "456", Type: parser.NUMBER_LEX},
	}

	for _, expectedToken := range expectedTokens {
		token := lexer()

		if token == nil {
			t.Fatalf("Expected a token but got nil")
		}

		if token.Type != expectedToken.Type {
			t.Errorf("Expected token type %s but got %s", expectedToken.Type, token.Type)
		}

		if token.Token != expectedToken.Token {
			t.Errorf("Expected token value '%s' but got '%s'", expectedToken.Token, token.Token)
		}
	}
}

func TestWhitespaceHandling(t *testing.T) {
	expression := " 123 + abc "
	lexer := parser.NewLexer(expression)

	expectedTokens := []parser.Lex{
		{Token: " ", Type: parser.SPECIAL_LEX},
		{Token: "123", Type: parser.NUMBER_LEX},
		{Token: " ", Type: parser.SPECIAL_LEX},
		{Token: "+", Type: parser.SPECIAL_LEX},
		{Token: " ", Type: parser.SPECIAL_LEX},
		{Token: "abc", Type: parser.STRING_LEX},
		{Token: " ", Type: parser.SPECIAL_LEX},
	}

	for _, expectedToken := range expectedTokens {
		token := lexer()

		if token == nil {
			t.Fatalf("Expected a token but got nil")
		}

		if token.Type != expectedToken.Type {
			t.Errorf("Expected token type %s but got %s", expectedToken.Type, token.Type)
		}

		if token.Token != expectedToken.Token {
			t.Errorf("Expected token value '%s' but got '%s'", expectedToken.Token, token.Token)
		}
	}
}
