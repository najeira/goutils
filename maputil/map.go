package maputil

import (
	vc "github.com/najeira/goutils/varutil"
)

func String(m map[string]interface{}, key string) (string, bool) {
	v, ok := m[key]
	if !ok {
		return "", false
	}
	return vc.String(v)
}

func Int(m map[string]interface{}, key string) (int64, bool) {
	v, ok := m[key]
	if !ok {
		return 0, false
	}
	return vc.Int(v)
}

func Float(m map[string]interface{}, key string) (float64, bool) {
	v, ok := m[key]
	if !ok {
		return 0, false
	}
	return vc.Float(v)
}

func Bool(m map[string]interface{}, key string) (bool, bool) {
	v, ok := m[key]
	if !ok {
		return false, false
	}
	return vc.Bool(v)
}
