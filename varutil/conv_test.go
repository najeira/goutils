package varutil

import (
	"database/sql"
	"testing"
)

func TestStringString(t *testing.T) {
	var s string = "hoge"
	if r, ok := String(interface{}(s)); ok {
		if r != s {
			t.Errorf("String invalid")
		}
	} else {
		t.Errorf("String failed")
	}
}

func TestStringStringPointer(t *testing.T) {
	var s string = "hoge"
	if r, ok := String(interface{}(&s)); ok {
		if r != s {
			t.Errorf("String invalid")
		}
	} else {
		t.Errorf("String failed")
	}
}

func TestStringNullString(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "hoge"}
	if r, ok := String(interface{}(s)); ok {
		if r != s.String {
			t.Errorf("String invalid")
		}
	} else {
		t.Errorf("String failed")
	}
}

func TestStringNullStringPointer(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "hoge"}
	if r, ok := String(interface{}(&s)); ok {
		if r != s.String {
			t.Errorf("String invalid")
		}
	} else {
		t.Errorf("String failed")
	}
}

func TestStringBytesFail(t *testing.T) {
	var s []byte = []byte("hoge")
	if _, ok := String(interface{}(s)); ok {
		t.Errorf("String returned")
	}
}

func TestIntInt64(t *testing.T) {
	var s int64 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != s {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt32(t *testing.T) {
	var s int32 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt16(t *testing.T) {
	var s int16 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt8(t *testing.T) {
	var s int8 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt(t *testing.T) {
	var s int = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt64Pointer(t *testing.T) {
	var s int64 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != s {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt32Pointer(t *testing.T) {
	var s int32 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt16Pointer(t *testing.T) {
	var s int16 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntInt8Pointer(t *testing.T) {
	var s int8 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntIntPointer(t *testing.T) {
	var s int = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUint64Fail(t *testing.T) {
	var s uint64 = 123
	if _, ok := Int(interface{}(s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntUint32(t *testing.T) {
	var s uint32 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUint16(t *testing.T) {
	var s uint16 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUint8(t *testing.T) {
	var s uint8 = 123
	if r, ok := Int(interface{}(s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUintFail(t *testing.T) {
	var s uint = 123
	if _, ok := Int(interface{}(s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntUint64PointerFail(t *testing.T) {
	var s uint64 = 123
	if _, ok := Int(interface{}(&s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntUint32Pointer(t *testing.T) {
	var s uint32 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUint16Pointer(t *testing.T) {
	var s uint16 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUint8Pointer(t *testing.T) {
	var s uint8 = 123
	if r, ok := Int(interface{}(&s)); ok {
		if r != int64(s) {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntUintPointerFail(t *testing.T) {
	var s uint = 123
	if _, ok := Int(interface{}(&s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntFloat64Fail(t *testing.T) {
	var s float64 = 123.456
	if _, ok := Int(interface{}(s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntFloat32Fail(t *testing.T) {
	var s float32 = 123.456
	if _, ok := Int(interface{}(s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntNullInt64(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if r, ok := Int(interface{}(s)); ok {
		if r != s.Int64 {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntNullFloat64Fail(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if _, ok := Int(interface{}(s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntNullStringFail(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if _, ok := Int(interface{}(s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntFloat64PointerFail(t *testing.T) {
	var s float64 = 123.456
	if _, ok := Int(interface{}(&s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntFloat32PointerFail(t *testing.T) {
	var s float32 = 123.456
	if _, ok := Int(interface{}(&s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntNullInt64Pointer(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if r, ok := Int(interface{}(&s)); ok {
		if r != s.Int64 {
			t.Errorf("Int invalid")
		}
	} else {
		t.Errorf("Int failed")
	}
}

func TestIntNullFloat64PointerFail(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if _, ok := Int(interface{}(&s)); ok {
		t.Errorf("Int returned")
	}
}

func TestIntNullStringPointerFail(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if _, ok := Int(interface{}(&s)); ok {
		t.Errorf("Int returned")
	}
}

func TestFloatInt64Fail(t *testing.T) {
	var s int64 = 123
	if _, ok := Float(interface{}(s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatInt32(t *testing.T) {
	var s int32 = 123
	if r, ok := Float(interface{}(s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatInt16(t *testing.T) {
	var s int16 = 123
	if r, ok := Float(interface{}(s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatInt8(t *testing.T) {
	var s int8 = 123
	if r, ok := Float(interface{}(s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatIntFail(t *testing.T) {
	var s int = 123
	if _, ok := Float(interface{}(s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatUint64Fail(t *testing.T) {
	var s uint64 = 123
	if _, ok := Float(interface{}(s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatUint32(t *testing.T) {
	var s uint32 = 123
	if r, ok := Float(interface{}(s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatUint16(t *testing.T) {
	var s uint16 = 123
	if r, ok := Float(interface{}(s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatUint8(t *testing.T) {
	var s uint8 = 123
	if r, ok := Float(interface{}(s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatUintFail(t *testing.T) {
	var s uint = 123
	if _, ok := Float(interface{}(s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatNullInt64Fail(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if _, ok := Float(interface{}(s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatNullFloat64(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if r, ok := Float(interface{}(s)); ok {
		if r != s.Float64 {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatNullStringFail(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if _, ok := Float(interface{}(s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatInt64PointerFail(t *testing.T) {
	var s int64 = 123
	if _, ok := Float(interface{}(&s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatInt32Pointer(t *testing.T) {
	var s int32 = 123
	if r, ok := Float(interface{}(&s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatInt16Pointer(t *testing.T) {
	var s int16 = 123
	if r, ok := Float(interface{}(&s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatInt8Pointer(t *testing.T) {
	var s int8 = 123
	if r, ok := Float(interface{}(&s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatIntPointerFail(t *testing.T) {
	var s int = 123
	if _, ok := Float(interface{}(&s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatUint64PointerFail(t *testing.T) {
	var s uint64 = 123
	if _, ok := Float(interface{}(&s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatUint32Pointer(t *testing.T) {
	var s uint32 = 123
	if r, ok := Float(interface{}(&s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatUint16Pointer(t *testing.T) {
	var s uint16 = 123
	if r, ok := Float(interface{}(&s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatUint8Pointer(t *testing.T) {
	var s uint8 = 123
	if r, ok := Float(interface{}(&s)); ok {
		if r != float64(s) {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatUintPointerFail(t *testing.T) {
	var s uint = 123
	if _, ok := Float(interface{}(&s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatNullInt64PointerFail(t *testing.T) {
	var s sql.NullInt64 = sql.NullInt64{Valid: true, Int64: 123}
	if _, ok := Float(interface{}(&s)); ok {
		t.Errorf("Float returned")
	}
}

func TestFloatNullFloat64Pointer(t *testing.T) {
	var s sql.NullFloat64 = sql.NullFloat64{Valid: true, Float64: 123.456}
	if r, ok := Float(interface{}(&s)); ok {
		if r != s.Float64 {
			t.Errorf("Float invalid")
		}
	} else {
		t.Errorf("Float failed")
	}
}

func TestFloatNullStringPointerFail(t *testing.T) {
	var s sql.NullString = sql.NullString{Valid: true, String: "123"}
	if _, ok := Float(interface{}(&s)); ok {
		t.Errorf("Float returned")
	}
}
