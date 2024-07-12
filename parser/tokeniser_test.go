package parser

import "testing"

func TestExpression(t *testing.T) {
	expression := " 123 & \"abc\" == 100+23&'abc'"
	tokeniser := NewTokeniser(expression)

	expectedTokens := []ParseToken{
		{Token: "123", Type: INTEGER},
		{Token: " ", Type: SPECIAL},
		{Token: "&", Type: SPECIAL},
		{Token: " ", Type: SPECIAL},
		{Token: "abc", Type: STRING},
		{Token: " ", Type: SPECIAL},
		{Token: "==", Type: SPECIAL},
		{Token: " ", Type: SPECIAL},
		{Token: "100", Type: INTEGER},
		{Token: "+", Type: SPECIAL},
		{Token: "23", Type: INTEGER},
		{Token: "&", Type: SPECIAL},
		{Token: "abc", Type: STRING},
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
