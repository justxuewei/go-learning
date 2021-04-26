package format

import "testing"

func TestAny(t *testing.T) {
	case1 := 1
	t.Log(Any(case1))
	case2 := []int{1, 2, 3}
	t.Log(Any(case2))
}
