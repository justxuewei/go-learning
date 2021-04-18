package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}
