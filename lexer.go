package main

import "strings"

const STRING_LEX string = "STRING"
const NUMBER_LEX string = "NUMBER"
const SPECIAL_LEX string = "SPECIAL"

type Lex struct {
	Token string
	Type  string
}

func NewLexer(expression string) func() *Lex {
	index := 0
	return func() *Lex {
		if len(expression) <= index {
			return nil
		}
		token := ""
		ttype := ""
		for ; index < len(expression); index++ {
			if expression[index] >= '0' && expression[index] <= '9' { //Number
				if ttype == NUMBER_LEX || ttype == STRING_LEX {
					token += string(expression[index])
				} else if ttype == SPECIAL_LEX {
					if token == "." {
						token += string(expression[index])
						ttype = NUMBER_LEX
					} else {
						return &Lex{
							Token: token,
							Type:  ttype,
						}
					}
				} else {
					token += string(expression[index])
					ttype = NUMBER_LEX
				}
			} else if (expression[index] >= 'a' && expression[index] <= 'z') || (expression[index] >= 'A' && expression[index] <= 'Z') {
				if ttype == NUMBER_LEX || ttype == STRING_LEX {
					token += string(expression[index])
					ttype = STRING_LEX
				} else if ttype == SPECIAL_LEX {
					return &Lex{
						Token: token,
						Type:  ttype,
					}
				} else {
					token += string(expression[index])
					ttype = STRING_LEX
				}
			} else {
				if ttype == SPECIAL_LEX {
					return &Lex{
						Token: token,
						Type:  ttype,
					}
				} else if ttype == NUMBER_LEX && expression[index] == '.' && !strings.Contains(token, ".") {
					token += string(expression[index])
				} else if ttype == NUMBER_LEX || ttype == STRING_LEX {
					return &Lex{
						Token: token,
						Type:  ttype,
					}
				} else {
					token += string(expression[index])
					ttype = SPECIAL_LEX
				}
			}
		}
		if token == "" || ttype == "" {
			return nil
		} else {
			return &Lex{
				Token: token,
				Type:  ttype,
			}
		}
	}
}
