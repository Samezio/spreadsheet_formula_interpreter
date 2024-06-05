package main

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
				if ttype == "NUMBER" || ttype == "TEXT" {
					token += string(expression[index])
				} else if ttype == "SPECIAL" {
					if token == "." {
						token += string(expression[index])
					} else {
						return &Lex{
							Token: token,
							Type:  ttype,
						}
					}
				} else {
					token += string(expression[index])
					ttype = "NUMBER"
				}
			} else if (expression[index] >= 'a' && expression[index] <= 'z') || (expression[index] >= 'A' && expression[index] <= 'Z') {
				if ttype == "NUMBER" || ttype == "TEXT" {
					token += string(expression[index])
					ttype = "TEXT"
				} else if ttype == "SPECIAL" {
					return &Lex{
						Token: token,
						Type:  ttype,
					}
				} else {
					token += string(expression[index])
					ttype = "TEXT"
				}
			} else {
				if ttype == "SPECIAL" {
					return &Lex{
						Token: token,
						Type:  ttype,
					}
				} else if ttype == "NUMBER" || ttype == "TEXT" {
					return &Lex{
						Token: token,
						Type:  ttype,
					}
				} else {
					token += string(expression[index])
					ttype = "SPECIAL"
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
