package parser

import (
	"errors"
	"github.com/abdivasiyev/json-parser/lexer"
	"reflect"
)

type Fn[T any] func(curr lexer.Token) (T, error)

type Parser struct {
	position int
	tokens   []lexer.Token
}

func New() Parser {
	return Parser{}
}

func (p *Parser) Parse(src []byte, dest any) error {
	l := lexer.New(src)

	for curr := l.Next(); curr != lexer.EofToken; curr = l.Next() {
		if curr == lexer.BadToken {
			return errors.New("invalid json signature")
		}

		p.tokens = append(p.tokens, curr)
	}

	return p.parse(dest)
}

func (p *Parser) parse(dest any) error {
	v := reflect.ValueOf(dest)

	if v.Kind() != reflect.Pointer {
		return errors.New("destination must be a pointer")
	}

	v = v.Elem()

	if !v.CanSet() {
		return errors.New("destination must be settable")
	}

	switch v.Kind() {
	case reflect.Pointer:
		return p.parse(v.Interface())
		// int parse
	case reflect.Int:
		return parseAndSetValue[int](p, v, Int)
	case reflect.Int8:
		return parseAndSetValue[int8](p, v, Int8)
	case reflect.Int16:
		return parseAndSetValue[int16](p, v, Int16)
	case reflect.Int32:
		return parseAndSetValue[int32](p, v, Int32)
	case reflect.Int64:
		return parseAndSetValue[int64](p, v, Int64)

		// uint parse
	case reflect.Uint:
		return parseAndSetValue[uint](p, v, Uint)
	case reflect.Uint8:
		return parseAndSetValue[uint8](p, v, Uint8)
	case reflect.Uint16:
		return parseAndSetValue[uint16](p, v, Uint16)
	case reflect.Uint32:
		return parseAndSetValue[uint32](p, v, Uint32)
	case reflect.Uint64:
		return parseAndSetValue[uint64](p, v, Uint64)

	// float parse
	case reflect.Float32:
		return parseAndSetValue[float32](p, v, Float32)
	case reflect.Float64:
		return parseAndSetValue[float64](p, v, Float64)

		// string parse
	case reflect.String:
		return parseAndSetValue[string](p, v, String)
	// bool parse
	case reflect.Bool:
		return parseAndSetValue[bool](p, v, Bool)
	// array
	case reflect.Array:
		return nil
	// slice
	case reflect.Slice:
		return p.parseSlice(v)
	// map
	case reflect.Map:
		return nil
	// struct
	case reflect.Struct:
		return nil
	default:
		return errors.New("unsupported type")
	}
}

func (p *Parser) peek() lexer.Token {
	if p.position >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[p.position]
}

func (p *Parser) next() lexer.Token {
	p.position++

	if p.position >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[p.position]
}

func parseAndSetValue[T any](p *Parser, v reflect.Value, fn Fn[T]) error {
	curr := p.peek()
	value, err := fn(curr)
	if err != nil {
		return err
	}

	rValue := reflect.ValueOf(value)
	cType := v.Type()

	switch {
	case rValue.Type().AssignableTo(cType):
		v.Set(rValue)
	case rValue.CanConvert(cType):
		v.Set(rValue.Convert(cType))
	default:
		return errors.New("could not convert json type to go's type")
	}

	return nil
}
