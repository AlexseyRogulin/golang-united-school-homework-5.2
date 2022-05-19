package cache

import "time"

var infinite time.Time = time.Time{}

type cacheValue struct {
	value    string
	deadline time.Time
}

type Cache struct {
	cache map[string]cacheValue
}

func NewCache() Cache {
	return Cache{cache: make(map[string]cacheValue)}
}

func (c Cache) Get(key string) (string, bool) {
	cacheValue, ok := c.cache[key]
	switch {
	case ok && cacheValue.deadline != infinite && time.Now().After(cacheValue.deadline):
		delete(c.cache, key)
		return "", false
	case ok:
		return cacheValue.value, true
	default:
		return "", false
	}
}

func (c Cache) Put(key, value string) {
	c.cache[key] = cacheValue{value: value, deadline: infinite}
}

func (c Cache) Keys() []string {
	var keys []string
	for k := range c.cache {
		if _, ok := c.Get(k); ok {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = cacheValue{value: value, deadline: deadline}
}
