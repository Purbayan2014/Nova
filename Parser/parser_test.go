package Parser

import (
	"Nova/ast"
	"Nova/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `letvar x = 4;` +
		`letvar y = 10;` +
		`letvar foobar = 83883;`

	// error stress testing (sucess)
	//input := `letvar x 4;` +
	//	`letvar = 10;` +
	//	`letvar 83883;`

	l := lexer.New(input)
	p := New(l)

	prog := p.ParserProgram()
	checkParseErrors(t, p)
	if prog == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(prog.Statements) != 3 {
		t.Fatalf("program.Statements doesnt contains 1 declared statements. got %d",
			len(prog.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := prog.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error : %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "letvar" {
		t.Errorf("s.TokenLiteral not 'letvar', got=%q", s.TokenLiteral())
		return false
	}

	letsmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	if letsmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letsmt.Name.Value)
		return false
	}

	if letsmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s' , got=%s", name, letsmt.Name)
		return false
	}

	return true
}
