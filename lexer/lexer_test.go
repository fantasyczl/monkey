package lexer

import (
	"testing"

	"github.com/fantasyczl/monkey/token"
)

func TestNextToken(t *testing.T) {
	type tokenList []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	tests := []struct {
		input    string
		expected tokenList
	}{
		{
			input: `=+(){},;`,
			expected: tokenList{
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
			input: `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
"hello \"world\""
"hello\n world"
"hello\n\t\t world"
"hello\n\t\t \"world"
[1, 2];
{"foo": "bar"}
macro(x, y) { x + y; };
`,
			expected: tokenList{
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
				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.INT, "10"},
				{token.NOT_EQ, "!="},
				{token.INT, "9"},
				{token.SEMICOLON, ";"},
				{token.STRING, "foobar"},
				{token.STRING, "foo bar"},
				{token.STRING, "hello \"world\""},
				{token.STRING, "hello\n world"},
				{token.STRING, "hello\n\t\t world"},
				{token.STRING, "hello\n\t\t \"world"},
				{token.LBRACKET, "["},
				{token.INT, "1"},
				{token.COMMA, ","},
				{token.INT, "2"},
				{token.RBRACKET, "]"},
				{token.SEMICOLON, ";"},
				{token.LBRACE, "{"},
				{token.STRING, "foo"},
				{token.COLON, ":"},
				{token.STRING, "bar"},
				{token.RBRACE, "}"},
				{token.MACRO, "macro"},
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
				{token.EOF, ""},
			},
		},
	}

	for i, p := range tests {
		l := New(p.input)
		for j, tt := range p.expected {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d.%d] - tokentype wrong. expected=%q, got=%q, literal=%q",
					i, j, tt.expectedType, tok.Type, tok.Literal)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d.%d] - literal wrong. expected=%q, got=%q",
					i, j, tt.expectedLiteral, tok.Literal)
			}
		}
	}
}
