package nlog

import (
	"bytes"
	"fmt"
	"testing"
)

func trimLF(s string) string {
	if s != "" {
		tail := len(s) - 1
		if s[tail] == '\n' {
			return s[:tail]
		}
	}
	return s
}

func TestPrintf(t *testing.T) {
	var buf bytes.Buffer
	config := Config{Out: &buf, Level: Warn}
	l := NewLogger(&config)

	var w, r, e string

	buf.Reset()
	w = "Test for Info"
	l.Infof(w)
	e = ""
	r = buf.String()
	if r != e {
		t.Errorf("%s != %s", trimLF(r), trimLF(e))
	}

	buf.Reset()
	w = "Test for Warn"
	l.Warnf(w)
	e = fmt.Sprintf("[WARN] %s\n", w)
	r = buf.String()
	if r != e {
		t.Errorf("%s != %s", trimLF(r), trimLF(e))
	}

	buf.Reset()
	w = "Test for Error"
	l.Errorf(w)
	e = fmt.Sprintf("[ERROR] %s\n", w)
	r = buf.String()
	if r != e {
		t.Errorf("%s != %s", trimLF(r), trimLF(e))
	}
}
