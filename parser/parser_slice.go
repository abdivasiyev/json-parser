package parser

import (
	"errors"
	"github.com/abdivasiyev/json-parser/lexer"
	"reflect"
)

func (p *Parser) parseSlice(v reflect.Value) error {
	var (
		curr        = p.peek()
		kind        = reflect.TypeOf(v.Interface()).Elem().Kind()
		values      []reflect.Value
		aliasedType = reflect.TypeOf(v.Interface()).Elem()
		value       reflect.Value
	)
	if curr.Kind != lexer.LSquareBrace {
		return errors.New("invalid json syntax")
	}

	for curr = p.next(); curr.Kind != lexer.RSquareBrace; curr = p.next() {
		switch curr.Kind {
		case lexer.Colon, lexer.Comma, lexer.WhiteSpace:
			continue
		default:
		}

		switch kind {
		case reflect.Int:
			v, err := Int(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Int8:
			v, err := Int8(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Int16:
			v, err := Int16(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Int32:
			v, err := Int32(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Int64:
			v, err := Int64(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Uint:
			v, err := Uint(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Uint8:
			v, err := Uint8(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Uint16:
			v, err := Uint16(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Uint32:
			v, err := Uint32(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Uint64:
			v, err := Uint64(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.String:
			v, err := String(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		case reflect.Bool:
			v, err := Bool(curr)
			if err != nil {
				return err
			}
			value = reflect.ValueOf(v)
		default:
			return errors.New("invalid array type")
		}

		if value.Type().AssignableTo(aliasedType) {
			values = append(values, value)
		} else if value.CanConvert(aliasedType) {
			values = append(values, value.Convert(aliasedType))
		}
	}
	v.Set(reflect.Append(v, values...))

	return nil
}
