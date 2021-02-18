package client

import "sync"

type Cache struct {
	Values []map[string]interface{} `json:"values"`

	lock *sync.RWMutex
}

func InitCacheClient() *Cache {
	return &Cache{
		Values: []map[string]interface{}{},
		lock:   &sync.RWMutex{},
	}
}

func (c *Cache) Find(k string, v interface{}) map[string]interface{} {
	var res map[string]interface{}
	c.lock.RLock()
	defer c.lock.RUnlock()
	for _, item := range c.Values {
		if item[k] != nil {
			if item[k] == v {
				res = item
			}
		}
	}
	return res
}

func (c *Cache) Add(value map[string]interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Values = append(c.Values, value)
}

func (c *Cache) Clean() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.Values = []map[string]interface{}{}
}

func (c *Cache) Delete(k string, v interface{}) {
	var newValues []map[string]interface{}
	c.lock.Lock()

	for _, item := range c.Values {
		if item[k] == nil || item[k] != v {
			newValues = append(newValues, item)
		}
	}

	c.Values = newValues
	c.lock.Unlock()
}
