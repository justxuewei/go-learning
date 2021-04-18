package ch7_interface

import (
	"fmt"
	"testing"
)

func TestByteCounter(t *testing.T) {
	var bc ByteCounter
	_, _ = bc.Write([]byte("hello"))
	fmt.Println(bc)
}
