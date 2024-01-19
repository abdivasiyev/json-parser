package lexer

import (
	"testing"
)

func TestLexer_Next(t *testing.T) {
	testCases := []struct {
		name   string
		src    []byte
		tokens []Token
	}{
		{
			name: "Empty",
			src:  []byte(``),
		},
		{
			name: "Empty object",
			src:  []byte(`{}`),
			tokens: []Token{
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
			},
		},
		{
			name: "Single key-value pair",
			src:  []byte(`{"key": "value"}`),
			tokens: []Token{
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  String,
					Value: "key",
				},
				{
					Kind:  Colon,
					Value: ":",
				},
				{
					Kind:  String,
					Value: "value",
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
			},
		},
		{
			name: "Key-value with integer",
			src:  []byte(`{"key": 10}`),
			tokens: []Token{
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  String,
					Value: "key",
				},
				{
					Kind:  Colon,
					Value: ":",
				},
				{
					Kind:  Number,
					Value: "10",
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
			},
		},
		{
			name: "Key-value with float",
			src:  []byte(`{"key": 10.3}`),
			tokens: []Token{
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  String,
					Value: "key",
				},
				{
					Kind:  Colon,
					Value: ":",
				},
				{
					Kind:    Number,
					Value:   "10.3",
					IsFloat: true,
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
			},
		},
		{
			name: "Key-value with bool",
			src:  []byte(`{"key": true}`),
			tokens: []Token{
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  String,
					Value: "key",
				},
				{
					Kind:  Colon,
					Value: ":",
				},
				{
					Kind:  Bool,
					Value: "true",
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
			},
		},
		{
			name: "Array test",
			src:  []byte(`[1, 2, 3]`),
			tokens: []Token{
				{
					Kind:  LSquareBrace,
					Value: "[",
				},
				{
					Kind:  Number,
					Value: "1",
				},
				{
					Kind:  Comma,
					Value: ",",
				},
				{
					Kind:  Number,
					Value: "2",
				},
				{
					Kind:  Comma,
					Value: ",",
				},
				{
					Kind:  Number,
					Value: "3",
				},
				{
					Kind:  RSquareBrace,
					Value: "]",
				},
			},
		},
		{
			name: "Only string",
			src:  []byte(`"value"`),
			tokens: []Token{
				{
					Kind:  String,
					Value: "value",
				},
			},
		},
		{
			name: "Only integer",
			src:  []byte(`10`),
			tokens: []Token{
				{
					Kind:  Number,
					Value: "10",
				},
			},
		},
		{
			name: "Only float",
			src:  []byte(`10.34`),
			tokens: []Token{
				{
					Kind:    Number,
					Value:   "10.34",
					IsFloat: true,
				},
			},
		},
		{
			name: "Negative",
			src:  []byte(`-10`),
			tokens: []Token{
				{
					Kind:  Number,
					Value: "-10",
				},
			},
		},
		{
			name: "Null value",
			src:  []byte(`null`),
			tokens: []Token{
				{
					Kind:  Null,
					Value: "null",
				},
			},
		},
		{
			name: "Invalid",
			src:  []byte(`.34`),
			tokens: []Token{
				BadToken,
				{
					Kind:  Number,
					Value: "34",
				},
			},
		},
		{
			name: "Complex structure",
			src:  []byte(`{"key": ["a", "b", 1, 2, {"key2": "value2"}]}`),
			tokens: []Token{
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  String,
					Value: "key",
				},
				{
					Kind:  Colon,
					Value: ":",
				},
				{
					Kind:  LSquareBrace,
					Value: "[",
				},
				{
					Kind:  String,
					Value: "a",
				},
				{
					Kind:  Comma,
					Value: ",",
				},
				{
					Kind:  String,
					Value: "b",
				},
				{
					Kind:  Comma,
					Value: ",",
				},
				{
					Kind:  Number,
					Value: "1",
				},
				{
					Kind:  Comma,
					Value: ",",
				},
				{
					Kind:  Number,
					Value: "2",
				},
				{
					Kind:  Comma,
					Value: ",",
				},
				{
					Kind:  LCurlyBrace,
					Value: "{",
				},
				{
					Kind:  String,
					Value: "key2",
				},
				{
					Kind:  Colon,
					Value: ":",
				},
				{
					Kind:  String,
					Value: "value2",
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
				{
					Kind:  RSquareBrace,
					Value: "]",
				},
				{
					Kind:  RCurlyBrace,
					Value: "}",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := New(tc.src)
			i := 0
			for curr := l.Next(); curr.Kind != Eof; curr = l.Next() {
				if curr.Kind == WhiteSpace {
					continue
				}

				if i >= len(tc.tokens) {
					t.Fatalf("incorrect length of tokens, expected: %d, got: %d", len(tc.tokens), i)
				}
				if curr.Kind != tc.tokens[i].Kind {
					t.Fatalf("incorrect kind, expected: %s, got: %s", tc.tokens[i].Kind, curr.Kind)
				}

				if curr.Value != tc.tokens[i].Value {
					t.Fatalf("incorrect value, expected: %s, got: %s", tc.tokens[i].Value, curr.Value)
				}

				if curr.IsFloat != tc.tokens[i].IsFloat {
					t.Fatalf("incorrect value, expected: %t, got: %t", tc.tokens[i].IsFloat, curr.IsFloat)
				}
				i++
			}
		})
	}
}
