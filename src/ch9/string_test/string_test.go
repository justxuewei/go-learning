package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF %x", s)
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStrConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
