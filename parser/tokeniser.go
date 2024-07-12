package parser

import (
	"fmt"
	"strings"
)

type ParseTokenType int

const (
	STRING ParseTokenType = iota
	BOOLEAN
	INTEGER
	DECIMAL
	SPECIAL
)

type ParseToken struct {
	Token string
	Type  ParseTokenType
}

type Tokeniser struct {
	getLex       func() *Lex
	currentLex   *Lex
	currentToken *ParseToken
}

func NewTokeniser(expression string) Tokeniser {
	return Tokeniser{
		getLex:       NewLexer(expression),
		currentLex:   nil,
		currentToken: nil}
}
func (t *Tokeniser) HasNext() bool {
	if t.currentToken == nil {
		t.Next()
	}
	return t.currentToken != nil
}
func (t *Tokeniser) Peek() *ParseToken {
	if t.currentToken == nil {
		t.Next()
	}
	return t.currentToken
}
func (t *Tokeniser) Next() *ParseToken {
	if t.currentLex == nil {
		t.currentLex = t.getLex()
	}
	if t.currentLex == nil {
		t.currentToken = nil
		return nil
	}

	if t.currentLex.Type == STRING_LEX {
		if t.currentLex.Token == "True" || t.currentLex.Token == "False" {
			t.currentToken = &ParseToken{
				Token: t.currentLex.Token,
				Type:  BOOLEAN,
			}
		} else {
			t.currentToken = &ParseToken{
				Token: t.currentLex.Token,
				Type:  SPECIAL,
			}
		}
	} else if t.currentLex.Type == NUMBER_LEX {
		if strings.Contains(t.currentLex.Token, ".") {
			t.currentToken = &ParseToken{
				Token: t.currentLex.Token,
				Type:  DECIMAL,
			}
		} else {
			t.currentToken = &ParseToken{
				Token: t.currentLex.Token,
				Type:  INTEGER,
			}
		}
	} else if t.currentLex.Type == SPECIAL_LEX {
		if t.currentLex.Token == ">" {
			t.currentLex = t.getLex()
			if t.currentLex.Token == "=" && t.currentLex.Type == SPECIAL_LEX {
				t.currentToken = &ParseToken{
					Token: ">=",
					Type:  SPECIAL,
				}
			} else {
				t.currentToken = &ParseToken{
					Token: ">",
					Type:  SPECIAL,
				}
				return t.currentToken
			}
		} else if t.currentLex.Token == "<" {
			t.currentLex = t.getLex()
			if t.currentLex.Token == "=" && t.currentLex.Type == SPECIAL_LEX {
				t.currentToken = &ParseToken{
					Token: "<=",
					Type:  SPECIAL,
				}
			} else {
				t.currentToken = &ParseToken{
					Token: "<",
					Type:  SPECIAL,
				}
				return t.currentToken
			}
		} else if t.currentLex.Token == "!" {
			t.currentLex = t.getLex()
			if t.currentLex.Token == "=" && t.currentLex.Type == SPECIAL_LEX {
				t.currentToken = &ParseToken{
					Token: "!=",
					Type:  SPECIAL,
				}
			} else {
				t.currentToken = &ParseToken{
					Token: "!",
					Type:  SPECIAL,
				}
				return t.currentToken
			}
		} else if t.currentLex.Token == "=" {
			t.currentLex = t.getLex()
			if t.currentLex.Token == "=" && t.currentLex.Type == SPECIAL_LEX {
				t.currentToken = &ParseToken{
					Token: "==",
					Type:  SPECIAL,
				}
			} else {
				t.currentToken = &ParseToken{
					Token: "=",
					Type:  SPECIAL,
				}
				return t.currentToken
			}
		} else if t.currentLex.Token == "\"" || t.currentLex.Token == "'" { //String
			quote := t.currentLex.Token
			t.currentLex = t.getLex()
			s := ""
			//Implement escape operator
			for ; t.currentLex != nil && (t.currentLex.Token != quote || t.currentLex.Type != SPECIAL_LEX); t.currentLex = t.getLex() {
				s += t.currentLex.Token
			}
			t.currentToken = &ParseToken{
				Token: s,
				Type:  STRING,
			}
		} else {
			t.currentToken = &ParseToken{
				Token: t.currentLex.Token,
				Type:  SPECIAL,
			}
		}

	} else {
		panic(fmt.Errorf("unknown lex: %v", t.currentLex))
	}
	t.currentLex = t.getLex()
	return t.currentToken
}
