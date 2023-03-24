package lexer

import (
	"Nova/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectType      token.TokenType
		expectedLiteral string
	}{
		// TEST 1
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},

		// {token.ILLEGAL, "ILLEGAL"},
		// {token.IDENT, "ident"},
		// {token.INT, "int"},
		// {token.FRAC, "frac"},
		// {token.COMMA, ","},
		// {token.SEMICOLON, ";"},
		// {token.LPAREN, "("},
		// {token.RPAREN, ")"},
		// {token.LBRACE, "{"},
		// {token.RBRACE, "}"},
		// {token.LSQUARE, "["},
		// {token.RSQUARE, "]"},
		// {token.ASSIGN, "="},
		// {token.PLUS, "+"},
		// {token.SUB, "-"},
		// {token.MULTI, "*"},
		// {token.DIV, "/"},
		// {token.LESS, "<"},
		// {token.GREATER, ">"},
		// {token.FUNCTION, "fn"},
		// {token.LETVAR, "letvar"},
		// {token.BANG, "!"},
		// {token.GREATERTHANEQ, ">="},
		// {token.LESSTHANEQ, "<="},
		// {token.NOTEQ, "!="},
		// {token.BANG, "!"},
		// {token.EQ, "=="},
		// {token.TRUE, "true"},
		// {token.FALSE, "false"},
		// {token.IF, "if"},
		// {token.ELSE, "else"},
		// {token.RETURN, "return"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong , expected=%q got=%q",
				i, tt.expectType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
