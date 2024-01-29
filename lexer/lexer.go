package lexer

import (
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

type TokenType int

const (
	LBRACKET TokenType = iota
	RBRACKET
	LBRACE
	RBRACE
	COLON
	COMMA
	STRING
	NUMBER
	BOOLEAN
	NULL
)

var symbolMap = map[rune]TokenType{
	'[': LBRACKET,
	']': RBRACKET,
	'{': LBRACE,
	'}': RBRACE,
	':': COLON,
	',': COMMA,
}

type Token struct {
	Type  TokenType
	pos   int
	Value string
}

type Lexer struct {
	input  []rune
	pos    int
	tokens []Token
}

func isSymbol(r rune) bool {
	_, ok := symbolMap[r]
	return ok
}

func (l *Lexer) Lex() error {
	for l.pos < len(l.input) {
		if unicode.IsSpace(l.input[l.pos]) {
			l.pos++
			continue
		}
		if isSymbol(l.input[l.pos]) {
			l.tokens = append(l.tokens, Token{symbolMap[l.input[l.pos]], l.pos, string(l.input[l.pos])})
			l.pos++
			continue
		}
		switch l.input[l.pos] {
		case '"':
			stringStart := l.pos
			var buf bytes.Buffer
			buf.WriteRune('"')
			l.pos++
			for l.pos < len(l.input) && (l.input[l.pos] != '"' || l.input[l.pos-1] == '\\') {
				buf.WriteRune(l.input[l.pos])
				l.pos++
			}
			if l.pos >= len(l.input) {
				return errors.New("unterminated string at position: " + strconv.Itoa(stringStart))
			}
			buf.WriteRune(l.input[l.pos])
			l.pos++
			l.tokens = append(l.tokens, Token{STRING, stringStart, buf.String()})
			continue

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			numStart := l.pos
			var buf bytes.Buffer
			buf.WriteRune(l.input[l.pos])
			l.pos++
			for l.pos < len(l.input) && (unicode.IsDigit(l.input[l.pos]) || l.input[l.pos] == '.') {
				buf.WriteRune(l.input[l.pos])
				l.pos++
			}
			l.tokens = append(l.tokens, Token{NUMBER, numStart, buf.String()})
			continue

		case 't', 'f':
			boolStart := l.pos
			nextFour := string(l.input[l.pos : l.pos+4])
			if nextFour != "true" && nextFour != "false" {
				return errors.New("unknown character: " + string(l.input[l.pos]) + " at position: " + strconv.Itoa(l.pos))
			}
			l.tokens = append(l.tokens, Token{BOOLEAN, boolStart, nextFour})
			l.pos += 4
		case 'n':
			nullStart := l.pos
			nextFour := string(l.input[l.pos : l.pos+4])
			if nextFour != "null" {
				return errors.New("unknown character: " + string(l.input[l.pos]) + " at position: " + strconv.Itoa(l.pos))
			}
			l.tokens = append(l.tokens, Token{NULL, nullStart, nextFour})
			l.pos += 4
		default:
			return errors.New("unknown character: " + string(l.input[l.pos]) + " at position: " + strconv.Itoa(l.pos))
		}
	}

	return nil
}

func (l *Lexer) Tokens() []Token {
	return l.tokens
}

func NewLexer(input []rune) *Lexer {
	return &Lexer{input: input}
}
