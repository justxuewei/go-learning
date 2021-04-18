package channel_close

import (
	"fmt"
	"testing"
)

func TestMake(t *testing.T) {
	a1 := make([]int, 5)
	fmt.Printf("%T %p\n", a1, &a1)
	sliceTest(a1)
	fmt.Println(a1[0])
}

func sliceTest(a []int) {
	fmt.Printf(">>> sliceTest(): %T %p\n", a, &a)
	a[0] = 1000
}

func TestMakeMap(t *testing.T) {
	m1 := make(map[int]int, 10)
	fmt.Printf("%T %p\n", m1, &m1)
	mapAsParameter(m1)
}

func mapAsParameter(m1 map[int]int) {
	fmt.Printf("%T %p\n", m1, &m1)
}
