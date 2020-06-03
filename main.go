package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (read *rot13Reader) Read(b []byte) (int, error) {
	end, err := read.r.Read(b)
	for i, c := range b {
		if i == end {
			break
		}
		b[i] = Rot13(c)
	}
	return end, err
}

func Rot13(c byte) byte {
	switch {
		case c >= 'A' && c < 'N':
			return c + 13
		case c >= 'N' && c < 'a':
			return c - 13
		case c >= 'a' && c < 'n':
			return c + 13
		case c >= 'n' && c <= 'z':
			return c - 13
		default:
			return c
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
