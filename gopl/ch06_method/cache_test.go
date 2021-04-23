package ch06_method

import "testing"

func TestCache(t *testing.T) {
	case1Key := "haha"
	case1Value := "hahavalue"
	cache := NewCache()
	cache.Write(case1Key, case1Value)
	t.Log(cache.Read(case1Key))
}
