package parser

import (
	"reflect"
	"testing"
)

func TestParser_ParseIntSlice(t *testing.T) {
	// int test
	{
		var (
			src  = []byte(`[1, 2, 3, 4]`)
			dest = []int{1, 2, 3, 4}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, []int{1, 2, 3, 4}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int8 test
	{
		var (
			src  = []byte(`[1, 2, 3, 4]`)
			dest = []int8{1, 2, 3, 4}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, []int8{1, 2, 3, 4}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int16 test
	{
		var (
			src  = []byte(`[1, 2, 3, 4]`)
			dest = []int16{1, 2, 3, 4}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, []int16{1, 2, 3, 4}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int32 test
	{
		var (
			src  = []byte(`[1, 2, 3, 4]`)
			dest = []int32{1, 2, 3, 4}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, []int32{1, 2, 3, 4}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int64 test
	{
		var (
			src  = []byte(`[1, 2, 3, 4]`)
			dest = []int64{1, 2, 3, 4}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, []int64{1, 2, 3, 4}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// custom int test
	{
		type Num int
		var (
			src  = []byte(`[1, 2, 3, 4]`)
			dest = []Num{1, 2, 3, 4}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, []Num{1, 2, 3, 4}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// [][]int test
	{
		var (
			src  = []byte(`[[1, 2, 3, 4], [1,2,3,4]]`)
			dest = [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}}
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if reflect.DeepEqual(dest, [][]int{{1, 2, 3, 4}, {1, 2, 3, 4}}) {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
}

func TestParser_ParseBool(t *testing.T) {
	// bool test
	{
		var (
			src  = []byte(`true`)
			dest bool
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != true {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	{
		var (
			src  = []byte(`false`)
			dest bool
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != false {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
}

func TestParser_ParseString(t *testing.T) {
	// string test
	{
		var (
			src  = []byte(`"hello world"`)
			dest string
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != "hello world" {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
}

func TestParser_ParseInt(t *testing.T) {
	// int test
	{
		var (
			src  = []byte(`10`)
			dest int
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
	// int8 test
	{
		var (
			src  = []byte(`10`)
			dest int8
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int16 test
	{
		var (
			src  = []byte(`10`)
			dest int16
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int32 test
	{
		var (
			src  = []byte(`10`)
			dest int32
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// int64 test
	{
		var (
			src  = []byte(`10`)
			dest int64
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
}

func TestParser_ParseUInt(t *testing.T) {
	// uint test
	{
		var (
			src  = []byte(`10`)
			dest uint
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
	// uint8 test
	{
		var (
			src  = []byte(`10`)
			dest uint8
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// uint16 test
	{
		var (
			src  = []byte(`10`)
			dest uint16
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// uint32 test
	{
		var (
			src  = []byte(`10`)
			dest uint32
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}

	// uint64 test
	{
		var (
			src  = []byte(`10`)
			dest uint64
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
}

func TestParser_ParseFloat(t *testing.T) {
	// float32 test
	{
		var (
			src  = []byte(`10.10`)
			dest float32
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10.10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
	// float64 test
	{
		var (
			src  = []byte(`10.10`)
			dest float64
		)

		p := New()

		err := p.Parse(src, &dest)
		if err != nil {
			t.Fatalf("parse failed: %v", err)
		}

		if dest != 10.10 {
			t.Fatalf("parser can't set value: %v", dest)
		}
	}
}
