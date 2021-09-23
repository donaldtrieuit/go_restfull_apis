package cache

import (
	"github.com/hashicorp/golang-lru/simplelru"
	"sync"
	"time"
)

type Item struct {
	val   interface{}
	mutex sync.Mutex
	expireTime *time.Time
}

type MutexByKeyCacheSource struct {
	lru  *simplelru.LRU
	lock sync.Mutex
}

func NewMutexByKeyCacheSource(size int) (*MutexByKeyCacheSource, error) {
	lru, err := simplelru.NewLRU(size, nil)
	if err != nil {
		return nil, err
	}

	c := MutexByKeyCacheSource{lru: lru, lock: sync.Mutex{}}
	return &c, nil
}

func (c *MutexByKeyCacheSource) GetString(key string) (string, bool) {
	i, ok := c.lru.Get(key)
	if !ok {
		return "", false
	}

	item := i.(*Item)
	if item.expireTime != nil && item.expireTime.Before(time.Now()) {
		c.lru.Remove(key)
		return "", false
	}

	item.mutex.Lock()
	defer item.mutex.Unlock()

	value := item.val.(string)

	return value, true
}

func (c *MutexByKeyCacheSource) GetInt(key string) (int, bool) {
	i, ok := c.lru.Get(key)
	if !ok {
		return 0, false
	}

	item := i.(*Item)
	if item.expireTime != nil && item.expireTime.Before(time.Now()) {
		c.lru.Remove(key)
		return 0, false
	}
	item.mutex.Lock()
	defer item.mutex.Unlock()

	value := item.val.(int)

	return value, true
}

func (c *MutexByKeyCacheSource) GetInterface(key string) (interface{}, bool) {
	i, ok := c.lru.Get(key)
	if !ok {
		return nil, false
	}

	item := i.(*Item)
	if item.expireTime != nil && item.expireTime.Before(time.Now()) {
		c.lru.Remove(key)
		return nil, false
	}
	item.mutex.Lock()
	defer item.mutex.Unlock()

	return item.val, true
}

func (c *MutexByKeyCacheSource) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	i, ok := c.lru.Get(key)
	if !ok {
		c.lock.Lock()
		item := Item{val: value}
		if expiration > 0 {
			t := time.Now().Add(expiration)
			item.expireTime = &t
		}
		c.set(key, &item)
		c.lock.Unlock()
		return "", nil
	}

	item := i.(*Item)
	item.mutex.Lock()
	defer item.mutex.Unlock()

	item.val = value
	return "", nil
}

func (c *MutexByKeyCacheSource) GetAndSetAtomicString(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value string, exists bool) (string, bool)) (string, string, bool) {
	i, ok := c.lru.Get(key)
	var out string
	var update bool

	if !ok {
		if out, update = cb("", false); update {
			c.lock.Lock()
			item := &Item{val: out}
			if valueExpiration > 0 {
				t := time.Now().Add(valueExpiration)
				item.expireTime = &t
			}
			c.set(key, item)
			c.lock.Unlock()
		}

		return "", out, false
	}

	item := i.(*Item)
	item.mutex.Lock()
	defer item.mutex.Unlock()

	before := item.val.(string)
	if out, update = cb(before, true); update {
		item.val = out
	}

	return before, out, true
}

func (c *MutexByKeyCacheSource) GetAndSetAtomicInt(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value int, exists bool) (int, bool)) (int, int, bool) {
	i, ok := c.lru.Get(key)
	var out int
	var update bool

	if !ok {
		if out, update = cb(0, false); update {
			c.lock.Lock()
			item := &Item{val: out}
			if valueExpiration > 0 {
				t := time.Now().Add(valueExpiration)
				item.expireTime = &t
			}
			c.set(key, item)
			c.lock.Unlock()
		}

		return 0, out, false
	}

	item := i.(*Item)
	item.mutex.Lock()
	defer item.mutex.Unlock()

	before := item.val.(int)
	if out, update = cb(before, true); update {
		item.val = out
	}

	return before, out, true
}

func (c *MutexByKeyCacheSource) GetAndSetAtomicInterface(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value interface{}, exists bool) (interface{}, bool)) (interface{}, interface{}, bool) {
	i, ok := c.lru.Get(key)
	var out interface{}
	var update bool

	if !ok {
		if out, update = cb(nil, false); update {
			c.lock.Lock()
			item := &Item{val: out}
			if valueExpiration > 0 {
				t := time.Now().Add(valueExpiration)
				item.expireTime = &t
			}
			c.set(key, item)
			c.lock.Unlock()
		}

		return nil, out, false
	}

	item := i.(*Item)
	item.mutex.Lock()
	defer item.mutex.Unlock()

	before := item.val
	if out, update = cb(before, true); update {
		item.val = out
	}

	return before, out, true
}

func (c *MutexByKeyCacheSource) set(key string, value *Item) {
	c.lru.Add(key, value)
}

func (c *MutexByKeyCacheSource) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	panic("Method not support")
}

func (c *MutexByKeyCacheSource) SAdd(key string, members ...interface{}) (int64, error) {
	panic("Method not support")
}

func (c *MutexByKeyCacheSource) SRem(key string, members ...interface{}) (int64, error) {
	panic("Method not support")
}

func (c *MutexByKeyCacheSource) SMembers(key string) ([]string, bool) {
	panic("Method not support")
}

func (c *MutexByKeyCacheSource) HMSet(key string, fields map[string]interface{}) (string, error) {
	panic("Method not support")
}

func (c *MutexByKeyCacheSource) HGetAll(key string) (map[string]string, bool) {
	panic("Method not support")
}

func (c *MutexByKeyCacheSource) Del(keys ...string) error {
	for _, key := range keys {
		i, ok := c.lru.Get(key)
		if !ok {
			return nil
		}

		item := i.(*Item)
		item.mutex.Lock()

		c.lru.Remove(key)

		item.mutex.Unlock()
	}

	return nil
}

func (c *MutexByKeyCacheSource) BindKey() string {
	return "cache.source.mutex_by_key"
}