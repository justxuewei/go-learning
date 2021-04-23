package ch06_method

import "sync"

type Cache struct {
	sync.Mutex
	mapping map[string]string
}

func NewCache() *Cache {
	return &Cache{
		mapping: make(map[string]string, 0),
	}
}

func (c *Cache) Write(key, value string) {
	c.Lock()
	c.mapping[key] = value
	c.Unlock()
}

func (c *Cache) Read(key string) (string, bool) {
	ret1, ret2 := "", false
	c.Lock()
	if v, ok := c.mapping[key]; ok {
		ret1 = v
		ret2 = true
	}
	c.Unlock()
	return ret1, ret2
}
