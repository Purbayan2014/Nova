package Parser

import (
	"Nova/ast"
	"Nova/lexer"
	"Nova/token"
)

// Parser Structure of our Parser
type Parser struct {
	// instance of the lexer
	l *lexer.Lexer

	curtoken  token.Token
	peektoken token.Token // peek token is used if the currtoken doesn't gives enough information
}

// NextToken is repeatedly called to the next token in the input curr token and
// the nexttoken will be the two pointers in the lexer the pointers are position
// and readpos but instead of pointing to the token directly we point to the
// currtoken and the nexttoken
func (p *Parser) nextToken() {
	p.curtoken = p.peektoken
	p.peektoken = p.l.NextToken()

}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}

// Advances both the next and the curr token
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// setting the next token and the curr token
	// by reading two tokens consecutively
	p.nextToken()
	p.nextToken()

	return p
}
