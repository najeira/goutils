package stringutil

import (
	"bytes"
	"io"
	"testing"
)

func TestReadLine(t *testing.T) {
	b := bytes.NewBuffer([]byte(`おはよう
こんにちは
こんばんは`))
	s, err := ReadLine(b)
	if err != nil {
		t.Error(err)
	}
	if s != "おはよう" {
		t.Error("invalid line")
	}

	s, err = ReadLine(b)
	if err != nil {
		t.Error(err)
	}
	if s != "こんにちは" {
		t.Error("invalid line")
	}

	s, err = ReadLine(b)
	if err != io.EOF {
		t.Error(err)
	}
	if s != "こんばんは" {
		t.Error("invalid line")
	}
}
