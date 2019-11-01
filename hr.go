package hr

import (
	"strings"
)

// HorizontalRule produces a string with a length matched to the terminal dimensions
func HorizontalRule(line ...string) string {
	cols, _, err := TerminalDimensions()
	if err != nil {
		return "Error: " + err.Error()
	}

	var w strings.Builder

	for _, str := range line {
		var runes int
		r := NewInfiniteRuneReader(str)
		if r == nil {
			continue
		}

		for runes < cols {
			ch, _, err := r.ReadRune()
			if err != nil {
				panic(err)
			}
			if _, err := w.WriteRune(ch); err != nil {
				panic(err)
			}
			runes++
		}
		w.Write([]byte(newline))
	}

	return w.String()
}
