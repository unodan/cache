/*
# File cache.go
# Author: Dan Huckson
# Date: 20160323
*/
package cache

import (
	"sync"
)

type Store struct {
	lck sync.RWMutex
	cfg map[string]interface{}
}

func (c *Store) New() *Store {
	c = &Store{}
	c.cfg = make(map[string]interface{})
	return c
}
func (c *Store) Get(k string) interface{} {
	c.lck.RLock()
	defer c.lck.RUnlock()
	v, ok := c.cfg[k]

	if !ok {
		return ""
	}
	return v
}
func (c *Store) Set(k string, v interface{}) error {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.cfg[k] = v
	return nil
}
func (c *Store) Delete(k string) error {
	c.lck.Lock()

	defer c.lck.Unlock()
	delete(c.cfg, k)
	return nil
}
