package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
	// t.Log(Readable, Writeable, Executable)
	a := 6 // 0x0110
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
