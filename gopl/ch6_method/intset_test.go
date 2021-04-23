package ch6_method

import (
	"testing"
)

func TestIntSet(t *testing.T) {
	set := IntSet{}
	case1 := 123
	set.Add(case1)
	t.Log(set.Has(1))
	t.Log(set.Has(case1))
}

func TestIntSet_UnionWith(t *testing.T) {
	a, b := IntSet{}, IntSet{}
	a.Add(1)
	b.Add(100)
	a.UnionWith(&b)
	t.Log(a.Has(1))
	t.Log(a.Has(100))
	t.Log(a.Has(99))
}

func TestIntSet_String(t *testing.T) {
	set := IntSet{}
	case1 := 123
	case2 := 2435
	set.Add(case1)
	set.Add(case2)
	t.Log(set.String())
}

func TestIntSet_Len(t *testing.T) {
	set := IntSet{}
	case1 := 123
	case2 := 2435
	set.Add(case1)
	set.Add(case2)
	t.Log(set.Len())
}

func TestIntSet_Remove(t *testing.T) {
	set := IntSet{}
	case1 := 123
	case2 := 2435
	set.Add(case1)
	set.Add(case2)
	t.Log(set)
	set.Remove(case2)
	t.Log(set)
}

func TestIntSet_DifferenceWith(t *testing.T) {
	a, b := IntSet{}, IntSet{}
	a.Add(1)
	a.Add(100)
	b.Add(100)
	a.DifferenceWith(&b)
	t.Log(a)
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	a, b := IntSet{}, IntSet{}
	a.Add(1)
	a.Add(100)
	b.Add(100)
	b.Add(2)
	a.SymmetricDifference(&b)
	t.Log(a)
}

func TestIntSet_Elems(t *testing.T) {
	set := IntSet{}
	case1 := 123
	case2 := 2435
	set.Add(case1)
	set.Add(case2)
	t.Log(set.Elems())
}
