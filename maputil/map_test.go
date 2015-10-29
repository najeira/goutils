package maputil

import (
	"testing"
)

var m = map[string]interface{}{
	"i": int64(123),
	"f": float64(123.456),
	"s": string("foo"),
}

func TestString(t *testing.T) {
	if r, ok := String(m, "s"); ok {
		if r != "foo" {
			t.Errorf("String invalid")
		}
	} else {
		t.Errorf("String failed")
	}
}

func TestInt(t *testing.T) {
	if r, ok := Int(m, "i"); ok {
		if r != 123 {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestFloat(t *testing.T) {
	if r, ok := Float(m, "f"); ok {
		if r != 123.456 {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}
