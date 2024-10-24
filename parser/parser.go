// @TODO: refactor this whole thing please
package parser

import (
	"monkey/lexer"
	"monkey/token"
	"reflect"
	"strconv"
)

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
			s = append(s, p.parseLetStatement())
		}
	}

	return s
}

type Statement interface {
	statementIsEqual(Statement) bool
}

type LetStatement struct {
	Identifier string
	Value      Expression
}

func (l *LetStatement) statementIsEqual(s Statement) bool {
	v, ok := s.(*LetStatement)
	if !ok || v == nil {
		return false
	}

	return v.Identifier == l.Identifier
}

type Expression interface {
	expressionIsEqual(e Expression) bool
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

	// body is nil for now
	Body Expression
}

func (l *NumericExpression) expressionIsEqual(e Expression) bool {
	v, ok := e.(*NumericExpression)
	if !ok || v == nil {
		return false
	}

	return v.LeftOperator == l.LeftOperator && v.Operand == l.Operand && v.RightOperator.expressionIsEqual(l.RightOperator)
}

func (l *FunctionExpression) expressionIsEqual(e Expression) bool {
	v, ok := e.(*FunctionExpression)
	if !ok {
		return false
	}

	return reflect.DeepEqual(v.Parameters, l.Parameters) && v.Body.expressionIsEqual(l.Body)
}

func (p *Parser) parseLetStatement() *LetStatement {
	tok := p.l.NextToken()
	if tok.Type != token.IDENT {
		panic("Parse error: expecting an identifier")
	}
	ident := tok.Literal

	tok = p.l.NextToken()
	if tok.Type != token.ASSIGN {
		panic("Parse error: expecting an equal sign")
	}

	return &LetStatement{
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
	tok := p.l.NextToken()
	if tok.Type != token.FUNC {
		panic("Parse error: expected func keyword")
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
	for tok.Type != token.RBRACE && tok.Type != token.EOF {
		tok = p.l.NextToken()
	}

	if tok.Type != token.RBRACE {
		panic("Parse error: expected a } got EOF")
	}

	return &FunctionExpression{
		Parameters: params,
		Body:       nil,
	}
}
