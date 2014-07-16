package sqlutil

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var ErrNotFound = errors.New("not found")
var ErrInvalidType = errors.New("invalid type")

type NullBool struct {
	sql.NullBool
}

func (v NullBool) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Bool)
}

func (v *NullBool) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp == nil {
		v.Bool = false
		v.Valid = false
	} else {
		switch conv := tmp.(type) {
		case bool:
			v.Bool = conv
		default:
			return fmt.Errorf("sql: unmarshaling %T, got %T", v, tmp)
		}
		v.Valid = true
	}
	return nil
}

type NullFloat64 struct {
	sql.NullFloat64
}

func (v NullFloat64) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Float64)
}

func (v *NullFloat64) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp == nil {
		v.Float64 = 0
		v.Valid = false
	} else {
		switch conv := tmp.(type) {
		case float32:
			v.Float64 = float64(conv)
		case float64:
			v.Float64 = conv
		default:
			return fmt.Errorf("sql: unmarshaling %T, got %T", v, tmp)
		}
		v.Valid = true
	}
	return nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Int64)
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp == nil {
		v.Int64 = 0
		v.Valid = false
	} else {
		switch conv := tmp.(type) {
		case int:
		case int8:
		case int16:
		case int32:
			v.Int64 = int64(conv)
		case int64:
			v.Int64 = conv
		default:
			return fmt.Errorf("sql: unmarshaling %T, got %T", v, tmp)
		}
		v.Valid = true
	}
	return nil
}

type NullString struct {
	sql.NullString
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.String)
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if tmp == nil {
		v.String = ""
		v.Valid = false
	} else {
		switch conv := tmp.(type) {
		case string:
			v.String = conv
			v.Valid = true
		default:
			return fmt.Errorf("sql: unmarshaling %T, got %T", v, tmp)
		}
	}
	return nil
}

type Value interface {
	sql.Scanner
	driver.Valuer
}

type Row map[string]Value

func (r Row) Value(name string) (Value, error) {
	v, ok := r[name]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (r Row) String(name string) (NullString, error) {
	n := NullString{}
	v, ok := r[name]
	if !ok {
		return n, ErrNotFound
	}
	var i interface{} = v
	switch d := i.(type) {
	case *NullString:
		return *d, nil
	case string:
		n.String = d
		n.Valid = true
		return n, nil
	}
	return n, ErrInvalidType
}

func (r Row) Int64(name string) (NullInt64, error) {
	n := NullInt64{}
	v, ok := r[name]
	if !ok {
		return n, ErrNotFound
	}
	var i interface{} = v
	switch d := i.(type) {
	case *NullInt64:
		return *d, nil
	case int64:
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
		n.Int64 = int64(d)
		n.Valid = true
		return n, nil
	case string:
		di, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			return n, err
		}
		n.Int64 = di
		n.Valid = true
		return n, nil
	}
	return n, ErrInvalidType
}

func (r Row) Float64(name string) (NullFloat64, error) {
	n := NullFloat64{}
	v, ok := r[name]
	if !ok {
		return n, ErrNotFound
	}
	var i interface{} = v
	switch d := i.(type) {
	case *NullFloat64:
		return *d, nil
	case float64:
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
		n.Float64 = float64(d)
		n.Valid = true
		return n, nil
	case string:
		di, err := strconv.ParseFloat(d, 64)
		if err != nil {
			return n, err
		}
		n := NullFloat64{}
		n.Float64 = di
		n.Valid = true
		return n, nil
	}
	return n, ErrInvalidType
}

func (r Row) Bool(name string) (NullBool, error) {
	n := NullBool{}
	v, ok := r[name]
	if !ok {
		return n, ErrNotFound
	}
	var i interface{} = v
	switch d := i.(type) {
	case *NullBool:
		return *d, nil
	case bool:
		n.Bool = d
		n.Valid = true
		return n, nil
	case int, int8, int16, int32, int64:
	case uint, uint8, uint16, uint32, uint64:
	case float32, float64:
		n.Bool = d != 0
		n.Valid = true
		return n, nil
	case string:
		di, err := strconv.ParseBool(d)
		if err != nil {
			return n, err
		}
		n.Bool = di
		n.Valid = true
		return n, nil
	}
	return n, ErrInvalidType
}

func (r Row) Bytes(name string) ([]byte, error) {
	v, ok := r[name]
	if !ok {
		return nil, ErrNotFound
	}
	var i interface{} = v
	switch d := i.(type) {
	case []byte:
		return d, nil
	}
	return nil, ErrInvalidType
}

type Rows struct {
	*sql.Rows
	columns  []string
	scanners []interface{}
	row      Row
}

type scanner struct {
	rows   *Rows
	column string
}

func (s *scanner) Scan(src interface{}) error {
	column, ok := s.rows.row[s.column]
	if !ok {
		return nil // ignore unknown columns
	}
	return column.Scan(src)
}

func NewRows(sqlRows *sql.Rows) (*Rows, error) {
	columns, err := sqlRows.Columns()
	if err != nil {
		return nil, err
	}
	r := &Rows{sqlRows, nil, nil, nil}
	scanners := make([]interface{}, len(columns))
	for i := range scanners {
		scanners[i] = &scanner{rows: r, column: columns[i]}
	}
	r.columns = columns
	r.scanners = scanners
	return r, nil
}

func (r *Rows) Scan(row Row) error {
	if row == nil {
		return errors.New("row is nil")
	}
	r.row = row
	err := r.Rows.Scan(r.scanners...)
	r.row = nil
	return err
}

func RowsToMaps(sqlRows *sql.Rows, newRow func() Row) ([]Row, error) {
	rows, err := NewRows(sqlRows)
	if err != nil {
		return nil, err
	}
	rets := make([]Row, 0)
	for rows.Next() {
		row := newRow()
		err = rows.Scan(row)
		if err != nil {
			return nil, err
		}
		rets = append(rets, row)
	}
	return rets, nil
}
