// @TODO: refactor this whole thing please
package parser

import (
	"fmt"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type Statement interface {
	statementNode() // used for debugging
}

type Expression interface {
	expressionNode() // used for debugging
}

// A numeric expression is simply a number & an operand plus another numeric expression
type NumericExpression struct {
	// @TODO: precise the operands
	Operand       string
	LeftOperator  int
	RightOperator *NumericExpression
}

type FunctionExpression struct {
	// for now i'm just capturing their names
	Parameters []string

	// bosdy is nil for now
	Body Expression
}

func (l *NumericExpression) expressionNode() {
	fmt.Printf("NumberExpression %+v\n", l)
}

func (l *FunctionExpression) expressionNode() {
	fmt.Printf("FunctionExpression %+v\n", l)
}

type LetExpression struct {
	Identifier string
	Value      Expression
}

func (l *LetExpression) expressionNode() {
	fmt.Printf("LetExpression %+v\n", l)
}

type Parser struct {
	l *lexer.Lexer
}

func NewParser(code string) Parser {
	l := lexer.NewLexer(code)
	return Parser{
		l: &l,
	}
}

func (p *Parser) ParseProgram() []Statement {
	s := make([]Statement, 0)

	for {
		tok := p.l.NextToken()
		if tok.Type == token.EOF {
			break
		}

		if tok.Type == token.LET {

		}
	}

	return s
}

func (p *Parser) parseLetExpression() *LetExpression {
	tok := p.l.NextToken()
	if tok.Type != token.IDENT {
		panic("Parse error: expecting an identifier")
	}
	ident := tok.Literal

	tok = p.l.NextToken()
	if tok.Type != token.ASSIGN {
		panic("Parse error: expecting an equal sign")
	}

	return &LetExpression{
		Identifier: ident,
		Value:      p.parseExpression(),
	}
}

func (p *Parser) parseExpression() Expression {
	tok := p.l.PeekToken()
	if tok.Type == token.FUNC {
		return p.parseFunctionExpression()
	} else if tok.Type == token.INT {
		return p.parseNumericExpression()
	}

	panic("Parse error: cannot parse let expression")
}

func (p *Parser) parseNumericExpression() *NumericExpression {
	tok := p.l.NextToken()
	if tok.Type != token.INT {
		panic("Parse error: expected a int found something else")
	}

	tokVal, _ := strconv.Atoi(tok.Literal)
	expr := &NumericExpression{
		LeftOperator: tokVal,
	}

	tok = p.l.NextToken()
	if tok.Type == token.SEMICOLON {
		return expr
	} else if tok.Type == token.PLUS || tok.Type == token.MINUS || tok.Type == token.ASTERISK {
		expr.Operand = tok.Literal
		expr.RightOperator = p.parseNumericExpression()
		return expr
	}

	panic("Parse error: cannot parse numeric epxression")
}

func (p *Parser) parseFunctionExpression() *FunctionExpression {
	// this might sounds stupid since i'm calling this func after checking if the next token is the 'func' keyword
	// but i'm assuming i might use this somewhere else where i assume what comes next should be a function expr
	// @TODO: if the above statement is false please delete this
	tok := p.l.NextToken()
	if tok.Type != token.LPAREN {
		panic("Parse error: expected a (")
	}

	tok = p.l.NextToken()
	if tok.Type != token.LPAREN {
		panic("Parse error: expected a (")
	}

	// parse parameters
	params := make([]string, 0)
	tok = p.l.NextToken()
	for tok.Type == token.IDENT {
		params = append(params, tok.Literal)
		tok = p.l.NextToken()
		if tok.Type == token.RPAREN {
			break
		} else if tok.Type == token.COMMA {
			tok = p.l.NextToken()
		} else {
			panic("Parse error: expected a ) or a parameter")
		}
	}

	tok = p.l.NextToken()
	if tok.Type != token.LBRACE {
		panic("Parse error: expected a {")
	}

	// @TODO: skip function body for now

	tok = p.l.NextToken()
	if tok.Type != token.RBRACE {
		panic("Parse error: expected a }")
	}

	return &FunctionExpression{
		Parameters: params,
		Body:       nil,
	}
}
