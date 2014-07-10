package stringutil

import (
	"bytes"
	"io"
)

type Reader interface {
	ReadRune() (r rune, size int, err error)
	UnreadRune() error
}

func ReadLine(r Reader) (string, error) {
	var b bytes.Buffer
	for {
		r1, err := readRune(r)
		if err != nil {
			if err == io.EOF {
				return b.String(), io.EOF
			}
			return "", err
		} else if r1 == '\n' {
			return b.String(), nil
		}
		b.WriteRune(r1)
	}
	panic("unreachable")
}

func readRune(r Reader) (rune, error) {
	r1, _, err := r.ReadRune()
	if r1 == '\r' {
		r1, _, err = r.ReadRune()
		if err == nil {
			if r1 != '\n' {
				r.UnreadRune()
				r1 = '\r'
			}
		}
	}
	return r1, err
}
