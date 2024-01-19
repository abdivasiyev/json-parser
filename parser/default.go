package parser

import (
	"errors"
	"github.com/abdivasiyev/json-parser/lexer"
	"strconv"
)

var (
	Int Fn[int] = func(curr lexer.Token) (int, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseInt(curr.Value, 10, 32)
		return int(value), nil
	}

	Int8 Fn[int8] = func(curr lexer.Token) (int8, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseInt(curr.Value, 10, 8)
		return int8(value), nil
	}

	Int16 Fn[int16] = func(curr lexer.Token) (int16, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseInt(curr.Value, 10, 16)
		return int16(value), nil
	}
	Int32 Fn[int32] = func(curr lexer.Token) (int32, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseInt(curr.Value, 10, 32)
		return int32(value), nil
	}
	Int64 Fn[int64] = func(curr lexer.Token) (int64, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseInt(curr.Value, 10, 64)
		return value, nil
	}

	Uint Fn[uint] = func(curr lexer.Token) (uint, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseUint(curr.Value, 10, 32)
		return uint(value), nil
	}
	Uint8 Fn[uint8] = func(curr lexer.Token) (uint8, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseUint(curr.Value, 10, 8)
		return uint8(value), nil
	}
	Uint16 Fn[uint16] = func(curr lexer.Token) (uint16, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseUint(curr.Value, 10, 16)
		return uint16(value), nil
	}
	Uint32 Fn[uint32] = func(curr lexer.Token) (uint32, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseUint(curr.Value, 10, 32)
		return uint32(value), nil
	}
	Uint64 Fn[uint64] = func(curr lexer.Token) (uint64, error) {
		if curr.Kind != lexer.Number || curr.IsFloat {
			return 0, errors.New("invalid integer format")
		}

		value, _ := strconv.ParseUint(curr.Value, 10, 64)
		return value, nil
	}

	Float32 Fn[float32] = func(curr lexer.Token) (float32, error) {
		if curr.Kind != lexer.Number || !curr.IsFloat {
			return 0, errors.New("invalid float32 format")
		}

		value, _ := strconv.ParseFloat(curr.Value, 32)
		return float32(value), nil
	}
	Float64 Fn[float64] = func(curr lexer.Token) (float64, error) {
		if curr.Kind != lexer.Number || !curr.IsFloat {
			return 0, errors.New("invalid float64 format")
		}

		value, _ := strconv.ParseFloat(curr.Value, 64)
		return value, nil
	}
	String Fn[string] = func(curr lexer.Token) (string, error) {
		if curr.Kind != lexer.String {
			return "", errors.New("invalid string")
		}
		return curr.Value, nil
	}

	Bool Fn[bool] = func(curr lexer.Token) (bool, error) {
		if curr.Kind != lexer.Bool {
			return false, errors.New("invalid boolean value")
		}
		return curr.Value == "true", nil
	}
)
