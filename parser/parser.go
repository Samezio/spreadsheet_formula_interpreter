package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Samezio/spreadsheet_formula_interpreter/ast"
)

func Parse(expression string) (ast.AST, error) {
	parseState := parseState{
		tokeniser: Tokeniser{
			getLex:       NewLexer(expression),
			currentLex:   nil,
			currentToken: nil},
	}
	parseState.currentToken = parseState.tokeniser.Next()
	return parseState.ParseExpression()
}

type parseState struct {
	tokeniser    Tokeniser
	currentToken *ParseToken
}

func (p *parseState) ParseExpression() (ast.AST, error) {
	return p.parseTerm5()
}
func (p *parseState) skipWhiteSpaces() {
	for p.currentToken != nil && p.currentToken.Type == SPECIAL && p.currentToken.Token == " " {
		p.currentToken = p.tokeniser.Next()
	}
}
func (p *parseState) consume(ttype ParseTokenType) error {
	if p.currentToken == nil {
		return fmt.Errorf("expected %v, got EOF", ttype)
	}
	if p.currentToken.Type != ttype {
		return fmt.Errorf("expected %v, but got %v", ttype, p.currentToken.Type)
	}
	p.currentToken = p.tokeniser.Next()
	p.skipWhiteSpaces()

	return nil
}
func (p *parseState) parseTerm5() (ast.AST, error) { //Inequalities
	left, err := p.parseTerm4()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && p.currentToken.Type == SPECIAL && p.currentToken.Token == "&" {
		operator := p.currentToken.Token
		err = p.consume(SPECIAL)
		if err != nil {
			return nil, err
		}
		right, err := p.parseTerm4()
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
func (p *parseState) parseTerm4() (ast.AST, error) { //Inequalities
	left, err := p.parseTerm3()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && p.currentToken.Type == SPECIAL && (p.currentToken.Token == ">=" || p.currentToken.Token == "<=" || p.currentToken.Token == ">" || p.currentToken.Token == "<") {
		operator := p.currentToken.Token
		err = p.consume(SPECIAL)
		if err != nil {
			return nil, err
		}
		right, err := p.parseTerm3()
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
func (p *parseState) parseTerm3() (ast.AST, error) { //Equalities
	left, err := p.parseTerm2()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && p.currentToken.Type == SPECIAL && (p.currentToken.Token == "==" || p.currentToken.Token == "!=") {
		operator := p.currentToken.Token
		err = p.consume(SPECIAL)
		if err != nil {
			return nil, err
		}
		right, err := p.parseTerm2()
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
func (p *parseState) parseTerm2() (ast.AST, error) {
	left, err := p.parseTerm1()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && p.currentToken.Type == SPECIAL && (p.currentToken.Token == "+" || p.currentToken.Token == "-") {
		operator := p.currentToken.Token
		err = p.consume(SPECIAL)
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
	left, err := p.unaryOperator()
	if err != nil {
		return nil, err
	}
	for p.currentToken != nil && p.currentToken.Type == SPECIAL && (p.currentToken.Token == "*" || p.currentToken.Token == "/") {
		operator := p.currentToken.Token
		err = p.consume(SPECIAL)
		if err != nil {
			return nil, err
		}
		right, err := p.unaryOperator()
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
func (p *parseState) unaryOperator() (ast.AST, error) {
	if p.currentToken == nil {
		return nil, nil
	}
	if p.currentToken.Type == SPECIAL && p.currentToken.Token == "-" { //Negation
		if err := p.consume(SPECIAL); err != nil {
			return nil, err
		} else if operand, err := p.factor(); err != nil {
			return nil, err
		} else {
			return &ast.UnaryOperator_AST{
				Operator: "-",
				Operand:  operand,
			}, nil
		}
	}
	return p.factor()
}
func (p *parseState) factor() (ast.AST, error) {
	if p.currentToken == nil {
		return nil, nil
	}
	switch p.currentToken.Type {
	case STRING:
		token := p.currentToken.Token
		p.consume(STRING)
		return ast.NewValueAST(token)
	case DECIMAL:
		token := p.currentToken.Token
		p.consume(DECIMAL)
		if value, err := strconv.ParseFloat(token, 64); err != nil {
			return nil, err
		} else {
			return ast.NewValueAST(value)
		}
	case INTEGER:
		token := p.currentToken.Token
		p.consume(INTEGER)
		if value, err := strconv.Atoi(token); err != nil {
			return nil, err
		} else {
			return ast.NewValueAST(value)
		}
	case BOOLEAN:
		token := p.currentToken.Token
		p.consume(BOOLEAN)
		return ast.NewValueAST(token == "True")
	case SPECIAL:
		if p.currentToken.Token == "(" || p.currentToken.Token == "{" || p.currentToken.Token == "[" {
			if err := p.consume(SPECIAL); err != nil {
				return nil, err
			} else if exp, err := p.ParseExpression(); err != nil {
				return nil, err
			} else if err := p.consume(SPECIAL); err != nil { //Closing
				return nil, err
			} else {
				return exp, nil
			}
		} else if match, err := regexp.MatchString("(\\w+\\d+)", p.currentToken.Token); err != nil { //TODO: correct it
			return nil, err
		} else if match {
			if cell_ast, err := ast.NewCellAST(p.currentToken.Token); err != nil {
				return nil, err
			} else if err := p.consume(SPECIAL); err != nil {
				return nil, err
			} else {
				return cell_ast, nil
			}
		}
	}
	return nil, fmt.Errorf("no valid factor found, got: %v", p.currentToken)
}
