package hr

import (
	"unicode/utf8"
)

// InfiniteRuneReader produces a rune from a string for infinity
type InfiniteRuneReader struct {
	s string
	i int64 // current reading index
}

// NewInfiniteRuneReader creates a new InfiniteRuneReader object
func NewInfiniteRuneReader(s string) *InfiniteRuneReader {
	if len(s) == 0 {
		return nil
	}
	return &InfiniteRuneReader{s, 0}
}

// ReadRune reads the next rune from the InfiniteRuneReader
func (r *InfiniteRuneReader) ReadRune() (ch rune, size int, err error) {
	if r.i >= int64(len(r.s)) {
		r.i = 0
	}
	if c := r.s[r.i]; c < utf8.RuneSelf {
		r.i++
		return rune(c), 1, nil
	}
	ch, size = utf8.DecodeRuneInString(r.s[r.i:])
	r.i += int64(size)
	return
}
