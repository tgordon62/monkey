package parser

import (
	"fmt"
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"monkey-interpreter/token"
)

type Parser struct {
	lex       *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	par := &Parser{
		lex:    lex,
		errors: []string{},
	}

	par.nextToken()
	par.nextToken()

	return par
}

func (par *Parser) nextToken() {
	par.curToken = par.peekToken
	par.peekToken = par.lex.NextToken()
}

func (par *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !par.curTokenIs(token.EOF) {
		stmt := par.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		par.nextToken()
	}

	return program
}

func (par *Parser) parseStatement() ast.Statement {
	switch par.curToken.Type {
	case token.LET:
		return par.parseLetStatement()
	default:
		return nil
	}
}

func (par *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: par.curToken}

	if !par.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: par.curToken, Value: par.curToken.Literal}

	if !par.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: skipping expressions until encounter semicolon
	for !par.expectPeek(token.SEMICOLON) {
		par.nextToken()
	}

	return stmt
}

func (par *Parser) curTokenIs(tok token.TokenType) bool {
	return par.curToken.Type == tok
}

func (par *Parser) peekTokenIs(tok token.TokenType) bool {
	return par.peekToken.Type == tok
}

func (par *Parser) expectPeek(tok token.TokenType) bool {
	if par.peekTokenIs(tok) {
		par.nextToken()
		return true
	} else {
		return false
	}
}

func (par *Parser) Errors() []string {
	return par.errors
}

func (par *Parser) peekError(tok token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		tok, par.peekToken.Type)
	par.errors = append(par.errors, msg)
}
