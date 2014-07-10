package jsonutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrInvalidType = errors.New("invalid type")
)

func Version() string {
	return "0.1.0"
}

type Value struct {
	Value interface{}
}

func (v *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(&v.Value)
}

func (v *Value) UnmarshalJSON(b []byte) error {
	d := json.NewDecoder(bytes.NewBuffer(b))
	return d.Decode(&v.Value)
}

func DecodeString(s string) (*Value, error) {
	v := &Value{}
	err := v.UnmarshalJSON([]byte(s))
	if err != nil {
		return nil, err
	}
	return v, nil
}

func DecodeReader(r io.Reader) (*Value, error) {
	d := json.NewDecoder(r)
	v := &Value{}
	err := d.Decode(&v.Value)
	return v, err
}

type Map map[string]interface{}

func (v Value) Map() (Map, error) {
	switch d := v.Value.(type) {
	case map[string]interface{}:
		return d, nil
	}
	return nil, ErrInvalidType
}

func (v Value) GetPath(name ...string) (*Value, bool) {
	m, err := v.Map()
	if err != nil || m == nil {
		return nil, false
	}
	return m.GetPath(name...)
}

func (m Map) Get(name string) (*Value, bool) {
	v, ok := m[name]
	if !ok {
		return nil, false
	}
	return &Value{v}, true
}

func (m Map) GetPath(name ...string) (*Value, bool) {
	ret := m
	for _, n := range name {
		v, ok := ret.Get(n)
		if !ok {
			return nil, false
		}
		ret, err := v.Map()
		if err != nil || ret == nil {
			return nil, false
		}
	}
	return &Value{ret}, true
}

func (m Map) Set(name string, value interface{}) {
	m[name] = &Value{value}
}

func (m Map) Del(name string) {
	delete(m, name)
}

type Array []interface{}

func (v Value) Array() (Array, error) {
	switch d := v.Value.(type) {
	case []interface{}:
		return d, nil
	}
	return nil, ErrInvalidType
}

func (a Array) Get(index int) *Value {
	return &Value{a[index]}
}

func (v Value) String() (string, error) {
	switch d := v.Value.(type) {
	case string:
		return d, nil
	}
	return "", ErrInvalidType
}

func (v Value) Int() (int64, error) {
	switch d := v.Value.(type) {
	case json.Number:
		return d.Int64()
	case int64:
		return d, nil
	case int:
	case int8:
	case int16:
	case int32:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
		return int64(d), nil
	case string:
		return strconv.ParseInt(d, 10, 64)
	}
	return 0, ErrInvalidType
}

func (v Value) Float() (float64, error) {
	switch d := v.Value.(type) {
	case json.Number:
		return d.Float64()
	case float64:
		return d, nil
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case float32:
		return float64(d), nil
	case string:
		return strconv.ParseFloat(d, 64)
	}
	return 0, ErrInvalidType
}

func (v Value) Bool() (bool, error) {
	switch d := v.Value.(type) {
	case bool:
		return d, nil
	case int, int8, int16, int32, int64:
	case uint, uint8, uint16, uint32, uint64:
	case float32, float64:
		return d != 0, nil
	case string:
		return strconv.ParseBool(d)
	}
	return false, ErrInvalidType
}
