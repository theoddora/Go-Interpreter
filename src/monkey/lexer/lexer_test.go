package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "frequently used",
			input: `=+(){},;`,
			expected: []token.Token{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "addition function",
			input: `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
`,
			expected: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)

		for i, expectedToken := range tt.expected {
			inputToken := l.NextToken()
			if inputToken.Type != expectedToken.Type {
				t.Fatalf("test %s - tests[%d] - tokentype wrong. expected=%q, got=%q",
					tt.name, i, expectedToken.Type, inputToken.Type)
			}

			if inputToken.Literal != expectedToken.Literal {
				t.Fatalf("test %s - tests[%d] - literal wrong. expected=%q, got=%q",
					tt.name, i, expectedToken.Literal, inputToken.Literal)
			}
		}

	}
}

func TestLexer(t *testing.T) {

}
