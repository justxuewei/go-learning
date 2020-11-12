package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// t.Log(a == c) Error!
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestBitClear(t *testing.T) {
	// t.Log(Readable, Writeable, Executable)
	a := 7 // 0x0111
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
