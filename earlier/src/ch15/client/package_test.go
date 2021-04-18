package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	if v, err := series.GetFibonacci(5); err == nil {
		t.Log(v)
	}
}
