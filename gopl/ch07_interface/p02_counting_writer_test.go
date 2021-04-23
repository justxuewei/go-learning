package ch07_interface

import "testing"

func TestCountingWriter(t *testing.T) {
	w, l := CountingWriter(&WordCount{})
	_, _ = w.Write([]byte("222"))
	if *l != 3 {
		t.Fatalf("l=%d", *l)
	}
}
