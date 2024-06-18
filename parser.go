package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
)

func Parse(expression string) (ast.AST, error) {
	parseState := parseState{
		lexer: NewLexer(expression),
	}
	parseState.currentLex = parseState.lexer()
	return parseState.ParseExpression()
}

type parseState struct {
	lexer      func() *Lex
	currentLex *Lex
}

func (p *parseState) ParseExpression() (ast.AST, error) {
	return p.parseTerm2()
}
func (p *parseState) skipWhiteSpaces() {
	for p.currentLex != nil && p.currentLex.Type == SPECIAL_LEX && p.currentLex.Token == " " {
		p.currentLex = p.lexer()
	}
}
func (p *parseState) consume(ttype string) error {
	if p.currentLex == nil {
		return fmt.Errorf("expected %s, got EOF", ttype)
	}
	if p.currentLex.Type != ttype {
		return fmt.Errorf("expected %s, but got %s", ttype, p.currentLex.Type)
	}
	p.currentLex = p.lexer()
	p.skipWhiteSpaces()

	return nil
}
func (p *parseState) parseTerm2() (ast.AST, error) {
	left, err := p.parseTerm1()
	if err != nil {
		return nil, err
	}
	for p.currentLex != nil && p.currentLex.Type == SPECIAL_LEX && (p.currentLex.Token == "+" || p.currentLex.Token == "-") {
		operator := p.currentLex.Token
		err = p.consume(SPECIAL_LEX)
		if err != nil {
			return nil, err
		}
		right, err := p.parseTerm1()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryOperator_AST{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left, nil
}
func (p *parseState) parseTerm1() (ast.AST, error) {
	left, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.currentLex != nil && p.currentLex.Type == SPECIAL_LEX && (p.currentLex.Token == "*" || p.currentLex.Token == "/") {
		operator := p.currentLex.Token
		err = p.consume(SPECIAL_LEX)
		if err != nil {
			return nil, err
		}
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryOperator_AST{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left, nil
}
func (p *parseState) factor() (ast.AST, error) {
	if p.currentLex == nil {
		return nil, nil
	}
	//TODO: remove Debug print
	//fmt.Printf("Factor: %v\n", p.currentLex)
	if p.currentLex.Type == NUMBER_LEX {
		token := p.currentLex.Token
		if err := p.consume(NUMBER_LEX); err != nil {
			return nil, err
		}
		if strings.ContainsRune(token, '.') {
			if value, err := strconv.ParseFloat(token, 64); err != nil {
				return nil, err
			} else {
				return ast.NewValueAST(value)
			}
		} else {
			if value, err := strconv.Atoi(token); err != nil {
				return nil, err
			} else {
				return ast.NewValueAST(value)
			}
		}
	} else if p.currentLex.Type == SPECIAL_LEX {
		if p.currentLex.Token == "\"" || p.currentLex.Token == "'" {
			// String start
			closingChar := "'"
			if p.currentLex.Token == "\"" {
				closingChar = "\""
			}
			if err := p.consume(SPECIAL_LEX); err != nil {
				return nil, err
			}
			value := ""
			for p.currentLex != nil && p.currentLex.Token != closingChar {
				value += p.currentLex.Token
				if err := p.consume(p.currentLex.Type); err != nil {
					return nil, err
				}
			}
			if err := p.consume(SPECIAL_LEX); err != nil { //Closing quote
				return nil, err
			}
			return ast.NewValueAST(value)
		} else if p.currentLex.Token == "(" || p.currentLex.Token == "[" || p.currentLex.Token == "{" {
			/*closingChar := "("
			if p.currentLex.Token == "[" {
				closingChar = "["
			} else if p.currentLex.Token == "{" {
				closingChar = "{"
			}*/

			if err := p.consume(SPECIAL_LEX); err != nil {
				return nil, err
			}
			if exp, err := p.ParseExpression(); err != nil {
				return nil, err
			} else if err := p.consume(SPECIAL_LEX); err != nil { //Closing
				return nil, err
			} else {
				return exp, nil
			}
		}
	}
	return nil, fmt.Errorf("no valid factor found, got: %v", p.currentLex)
}
