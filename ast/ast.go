package ast

import "Nova/token"

// Node base node of the abstract syntax tree
// TokenLiteral provides the value of the literal in the statements
type Node interface {
	TokenLiteral() string
}

// Statement The Nodes will implement the statements and the Expressions
// Dummy Nodes are used for connecting each other
type Statement interface {
	Node
	statementNode() // dummy Node
}

type Expression interface {
	Node
	expressionNode() // dummy Node
}

// Program Root node of every AST that the parser produces
// Every valid source code will be a series of statements
// The Statements are contained in the slice which is a slice of ast
// that implement the AST interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // token.Ident token
	Value string      // holds the value of the identifier that is v = 3 holds 3
}

type LetStatement struct {
	Token token.Token // token.letvar token
	Name  *Identifier // name of the token
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
