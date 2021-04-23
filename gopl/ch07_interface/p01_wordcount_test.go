package ch07_interface

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	wc := WordCount{}
	_, _ = fmt.Fprintf(&wc, "line1 word1 word2\nline2 word3 word4")
	if wc.nline != 2 {
		t.Fatal()
	}
	if wc.nword != 6 {
		t.Fatal()
	}
}
