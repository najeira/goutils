package varconv

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

const (
	Version = "0.1.0"
)

func String(v interface{}) (string, bool) {
	switch d := v.(type) {
	case string:
		return d, true
	case sql.NullString:
		return d.String, true
	}
	return "", false
}

func TryString(v interface{}) (string, bool) {
	if n, ok := String(v); ok {
		return n, true
	}
	switch d := v.(type) {
	case []byte:
		return string(d), true
	}
	return "", false
}

func Int(v interface{}) (int64, bool) {
	switch d := v.(type) {
	case int64:
		return d, true
	case int:
		return int64(d), true
	case int8:
		return int64(d), true
	case int16:
		return int64(d), true
	case int32:
		return int64(d), true
	case uint8:
		return int64(d), true
	case uint16:
		return int64(d), true
	case uint32:
		return int64(d), true
	case sql.NullInt64:
		return d.Int64, true
	}
	return 0, false
}

func TryInt(v interface{}) (int64, bool) {
	if n, ok := Int(v); ok {
		return n, true
	}
	switch d := v.(type) {
	case uint64:
		return int64(d), true
	case float32:
		return int64(d), true
	case float64:
		return int64(d), true
	case json.Number:
		if n, err := d.Int64(); err == nil {
			return n, true
		}
	case sql.NullFloat64:
		return int64(d.Float64), true
	case sql.NullBool:
		if d.Bool {
			return 1, true
		} else {
			return 0, true
		}
	}
	if s, ok := TryString(v); ok {
		if n, err := strconv.ParseInt(s, 10, 64); err == nil {
			return n, true
		}
	}
	return 0, false
}

func Float(v interface{}) (float64, bool) {
	switch d := v.(type) {
	case float64:
		return d, true
	case int8:
		return float64(d), true
	case int16:
		return float64(d), true
	case int32:
		return float64(d), true
	case uint8:
		return float64(d), true
	case uint16:
		return float64(d), true
	case uint32:
		return float64(d), true
	case float32:
		return float64(d), true
	case sql.NullFloat64:
		return d.Float64, true
	}
	return 0, false
}

func TryFloat(v interface{}) (float64, bool) {
	if n, ok := Float(v); ok {
		return n, true
	}
	switch d := v.(type) {
	case int64:
		return float64(d), true
	case uint64:
		return float64(d), true
	case json.Number:
		if n, err := d.Float64(); err == nil {
			return n, true
		}
	case sql.NullInt64:
		return float64(d.Int64), true
	case sql.NullBool:
		if d.Bool {
			return 1, true
		} else {
			return 0, true
		}
	}
	if s, ok := TryString(v); ok {
		if n, err := strconv.ParseFloat(s, 64); err == nil {
			return n, true
		}
	}
	return 0, false
}

func Bool(v interface{}) (bool, bool) {
	switch d := v.(type) {
	case bool:
		return d, true
	case sql.NullBool:
		return d.Bool, true
	}
	return false, false
}

func TryBool(v interface{}) (bool, bool) {
	if n, ok := Bool(v); ok {
		return n, true
	}
	switch d := v.(type) {
	case int:
		return d != 0, true
	case int8:
		return d != 0, true
	case int16:
		return d != 0, true
	case int32:
		return d != 0, true
	case int64:
		return d != 0, true
	case uint:
		return d != 0, true
	case uint8:
		return d != 0, true
	case uint16:
		return d != 0, true
	case uint32:
		return d != 0, true
	case uint64:
		return d != 0, true
	case float32:
		return d != 0, true
	case float64:
		return d != 0, true
	}
	if s, ok := TryString(v); ok {
		if n, err := strconv.ParseBool(s); err == nil {
			return n, true
		}
	}
	return false, false
}
