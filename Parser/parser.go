package Parser

import (
	"Nova/ast"
	"Nova/lexer"
	"Nova/token"
	"fmt"
)

// Parser Structure of our Parser
type Parser struct {
	// instance of the lexer
	l *lexer.Lexer

	curtoken  token.Token
	peektoken token.Token // peek token is used if the currtoken doesn't gives enough information
	errors    []string    // holds the error for types that are not supported by the parser
}

// NextToken is repeatedly called to the next token in the input curr token and
// the nexttoken will be the two pointers in the lexer the pointers are position
// and readpos but instead of pointing to the token directly we point to the
// currtoken and the nexttoken
func (p *Parser) nextToken() {
	p.curtoken = p.peektoken
	p.peektoken = p.l.NextToken()

}

// parses each statement in the program
func (p *Parser) parseStatement() ast.Statement {
	switch p.curtoken.Type {
	case token.LETVAR:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// ParserProgram Construct the root node of the ast an *ast.prog then iterates over each of the token
// in the input until it gets the EOF then calls the nexttoken which advances the p.currtoken
// and the p.peektoken and the parsestatement then parses the statement at each iteration
func (p *Parser) ParserProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.curtoken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}
	return prog
}

// Constructs an *ast.letstatement node with the token its setting on the token.letvar and then
// advances the token while making assertions about the next token with calls to expectPeek
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curtoken}

	// expectPeek is used to create the *ast.Identifier node
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curtoken, Value: p.curtoken.Literal}

	// expects the equal
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO : expressions

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt

}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peektokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peektokenIs(t token.TokenType) bool {
	return p.peektoken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curtoken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peektoken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

// Advances both the next and the curr token
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// setting the next token and the curr token
	// by reading two tokens consecutively
	p.nextToken()
	p.nextToken()

	return p
}
