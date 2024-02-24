package lexer

import (
	"reflect"
	"testing"
)

// TestLexer tests the lexer.
func TestLexer(t *testing.T) {
	tests := []struct {
		input  string
		tokens []Token
	}{
		{
			input: `{"name": "John Doe", "age": 30}`,
			tokens: []Token{
				{Type: LBRACE, Value: "{", Pos: 0},
				{Type: STRING, Value: "name", Pos: 1},
				{Type: COLON, Value: ":", Pos: 7},
				{Type: STRING, Value: "John Doe", Pos: 9},
				{Type: COMMA, Value: ",", Pos: 19},
				{Type: STRING, Value: "age", Pos: 21},
				{Type: COLON, Value: ":", Pos: 26},
				{Type: NUMBER, Value: "30", Pos: 28},
				{Type: RBRACE, Value: "}", Pos: 30},
			},
		},
		{
			input: `[1, 2, 3, 4, 5]`,
			tokens: []Token{
				{Type: LBRACKET, Value: "[", Pos: 0},
				{Type: NUMBER, Value: "1", Pos: 1},
				{Type: COMMA, Value: ",", Pos: 2},
				{Type: NUMBER, Value: "2", Pos: 4},
				{Type: COMMA, Value: ",", Pos: 5},
				{Type: NUMBER, Value: "3", Pos: 7},
				{Type: COMMA, Value: ",", Pos: 8},
				{Type: NUMBER, Value: "4", Pos: 10},
				{Type: COMMA, Value: ",", Pos: 11},
				{Type: NUMBER, Value: "5", Pos: 13},
				{Type: RBRACKET, Value: "]", Pos: 14},
			},
		},
		{
			input: `true`,
			tokens: []Token{
				{Type: BOOLEAN, Value: "true", Pos: 0},
			},
		},
		{
			input: `false`,
			tokens: []Token{
				{Type: BOOLEAN, Value: "false", Pos: 0},
			},
		},
		{
			input: `null`,
			tokens: []Token{
				{Type: NULL, Value: "null", Pos: 0},
			},
		},
	}

	for _, test := range tests {
		l := NewLexer([]rune(test.input))
		err := l.Lex()
		if err != nil {
			t.Errorf("Error thrown by Lexer: %v", err)
		}

		if !reflect.DeepEqual(l.Tokens(), test.tokens) {
			t.Errorf("Expected tokens: %v, got: %v", test.tokens, l.Tokens())
		}
	}
}
