package jsonutil

import (
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	s := `{"users": [{"id": 123, "name": "Alice"}]}`

	// decode
	v, err := DecodeString(s)
	if err != nil {
		t.Error(err)
	}
	if v == nil {
		t.Error("DecodeString() retuns nil")
	}

	// v should be a Map
	m, err := v.Map()
	if err != nil {
		t.Error(err)
	}
	if m == nil {
		t.Error("Map() retuns nil")
	}

	// m should have users
	va, ok := m.Get("users")
	if !ok {
		t.Error("Map.Get() has not value")
	}
	if va == nil {
		t.Error("Map.Get() retuns nil")
	}

	// va should be Array
	a, err := va.Array()
	if err != nil {
		t.Error(err)
	}
	if a == nil {
		t.Error("Array() retuns nil")
	}
	if len(a) != 1 {
		t.Error("a array has invalid length")
	}

	// get a value
	vu := a.Get(0)
	if vu == nil {
		t.Error("a array does not have a value")
	}

	// vu should be a Map
	u, err := vu.Map()
	if err != nil {
		t.Error(err)
	}
	if u == nil {
		t.Error("Map() retuns nil")
	}

	vi, ok := u.Get("id")
	if !ok {
		t.Error("id not found")
	}
	if vi == nil {
		t.Error("id is nil")
	}

	i, err := vi.Int()
	if err != nil {
		t.Error(err)
	}
	if i != 123 {
		t.Error("invalid id")
	}

	vn, ok := u.Get("name")
	if !ok {
		t.Error("name not found")
	}
	if vn == nil {
		t.Error("name is nil")
	}

	n, err := vn.String()
	if err != nil {
		t.Error(err)
	}
	if n != "Alice" {
		t.Error("invalid name")
	}
}
