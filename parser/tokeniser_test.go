package parser_test

import (
	"testing"

	"github.com/Samezio/spreadsheet_formula_interpreter/parser"
)

func TestExpression(t *testing.T) {
	expression := " 123 & \"abc\" == 100+23&'abc'"
	tokeniser := parser.NewTokeniser(expression)

	expectedTokens := []parser.ParseToken{
		{Token: "123", Type: parser.INTEGER},
		{Token: " ", Type: parser.SPECIAL},
		{Token: "&", Type: parser.SPECIAL},
		{Token: " ", Type: parser.SPECIAL},
		{Token: "abc", Type: parser.STRING},
		{Token: " ", Type: parser.SPECIAL},
		{Token: "==", Type: parser.SPECIAL},
		{Token: " ", Type: parser.SPECIAL},
		{Token: "100", Type: parser.INTEGER},
		{Token: "+", Type: parser.SPECIAL},
		{Token: "23", Type: parser.INTEGER},
		{Token: "&", Type: parser.SPECIAL},
		{Token: "abc", Type: parser.STRING},
	}

	for _, expectedToken := range expectedTokens {
		if !tokeniser.HasNext() {
			t.Fatalf("Expected a token, but tokeniser is ended")
		}
		token := tokeniser.Next()
		t.Logf("token: %v\n", token)
		if token == nil {
			t.Fatalf("Expected a token but got nil")
		}

		if token.Type != expectedToken.Type {
			t.Errorf("Expected token type %v but got %v", expectedToken.Type, token.Type)
		}

		if token.Token != expectedToken.Token {
			t.Errorf("Expected token value '%s' but got '%s'", expectedToken.Token, token.Token)
		}
	}
}
