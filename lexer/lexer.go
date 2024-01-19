package lexer

import (
	"slices"
	"strings"
)

type Lexer struct {
	src      []byte
	position int
}

func New(src []byte) Lexer {
	return Lexer{
		src:      src,
		position: 0,
	}
}

func (l *Lexer) Next() Token {
	if l.position >= len(l.src) {
		return EofToken
	}

	curr := l.position
	l.position++

	switch l.src[curr] {
	case '{':
		return Token{
			Kind:  LCurlyBrace,
			Value: "{",
		}
	case '}':
		return Token{
			Kind:  RCurlyBrace,
			Value: "}",
		}
	case '[':
		return Token{
			Kind:  LSquareBrace,
			Value: "[",
		}
	case ']':
		return Token{
			Kind:  RSquareBrace,
			Value: "]",
		}
	case ':':
		return Token{
			Kind:  Colon,
			Value: ":",
		}
	case ',':
		return Token{
			Kind:  Comma,
			Value: ",",
		}
	case '"':
		var value strings.Builder
		for l.position < len(l.src) && l.src[l.position] != '"' {
			value.WriteByte(l.src[l.position])
			l.position++
		}
		l.position++

		return Token{
			Kind:  String,
			Value: value.String(),
		}
	case '\'':
		var value strings.Builder
		for l.position < len(l.src) && l.src[l.position] != '\'' {
			value.WriteByte(l.src[l.position])
			l.position++
		}
		l.position++
		return Token{
			Kind:  Char,
			Value: value.String(),
		}
	case ' ', '\t', '\n', '\r':
		return Token{
			Kind:  WhiteSpace,
			Value: string(l.src[curr]),
		}
	case 't':
		var value strings.Builder
		value.WriteByte(l.src[curr])
		for l.position < len(l.src) && value.Len() < len("true") {
			value.WriteByte(l.src[l.position])
			l.position++
		}
		return Token{
			Kind:  Bool,
			Value: value.String(),
		}
	case 'f':
		var value strings.Builder
		value.WriteByte(l.src[curr])
		for l.position < len(l.src) && value.Len() < len("false") {
			value.WriteByte(l.src[l.position])
			l.position++
		}
		return Token{
			Kind:  Bool,
			Value: value.String(),
		}
	case 'n':
		var value strings.Builder
		value.WriteByte(l.src[curr])
		for l.position < len(l.src) && value.Len() < len("null") {
			value.WriteByte(l.src[l.position])
			l.position++
		}
		return Token{
			Kind:  Null,
			Value: value.String(),
		}
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		var (
			value   strings.Builder
			isFloat bool
		)
		value.WriteByte(l.src[curr])
		for l.position < len(l.src) && l.isDigit(l.src[l.position]) {
			value.WriteByte(l.src[l.position])
			l.position++
		}

		for i := curr; i < l.position; i++ {
			if l.src[i] == '.' || l.src[i] == 'e' {
				isFloat = true
				break
			}
		}

		return Token{
			Kind:    Number,
			IsFloat: isFloat,
			Value:   value.String(),
		}
	default:
		return BadToken
	}
}

func (l *Lexer) isDigit(b byte) bool {
	return slices.Contains([]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '-', 'e'}, b)
}
