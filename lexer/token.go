package lexer

type Kind byte

func (k Kind) String() string {
	return mapping[k]
}

const (
	Bad Kind = iota
	LSquareBrace
	LCurlyBrace
	RSquareBrace
	RCurlyBrace
	Colon
	Comma
	WhiteSpace
	String
	Number
	Char
	Bool
	Null
	Eof
)

type Token struct {
	Kind    Kind
	IsFloat bool
	Value   string
}

var (
	BadToken = Token{
		Kind: Bad,
	}
	EofToken = Token{
		Kind: Eof,
	}

	mapping = map[Kind]string{
		Bad:          "Bad",
		LSquareBrace: "LSquareBrace",
		LCurlyBrace:  "LCurlyBrace",
		RSquareBrace: "RSquareBrace",
		RCurlyBrace:  "RCurlyBrace",
		Colon:        "Colon",
		Comma:        "Comma",
		WhiteSpace:   "WhiteSpace",
		String:       "String",
		Number:       "Number",
		Char:         "Char",
		Bool:         "Bool",
		Eof:          "Eof",
	}
)
