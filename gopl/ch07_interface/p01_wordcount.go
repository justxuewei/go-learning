package ch07_interface

import (
	"bufio"
	"bytes"
)

type WordCount struct {
	nword int
	nline int
}

func (w *WordCount) Write(b []byte) (int, error) {
	scanner1 := bufio.NewScanner(bytes.NewReader(b))
	scanner1.Split(bufio.ScanWords)
	for scanner1.Scan() {
		w.nword++
	}
	scanner2 := bufio.NewScanner(bytes.NewReader(b))
	scanner2.Split(bufio.ScanLines)
	for scanner2.Scan() {
		w.nline++
	}
	return len(b), nil
}
