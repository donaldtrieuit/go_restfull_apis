package cache

import "time"

type ICacheSource interface {
	GetString(key string) (string, bool)
	GetInt(key string) (int, bool)
	GetInterface(key string) (interface{}, bool)
	Set(key string, value interface{}, expiration time.Duration) (string, error)
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)
	SAdd(key string, members ...interface{}) (int64, error)
	SRem(key string, members ...interface{}) (int64, error)
	SMembers(key string) ([]string, bool)
	HMSet(key string, fields map[string]interface{}) (string, error)
	HGetAll(key string) (map[string]string, bool)
	GetAndSetAtomicString(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value string, exists bool) (string, bool)) (string, string, bool)
	GetAndSetAtomicInt(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value int, exists bool) (int, bool)) (int, int, bool)
	GetAndSetAtomicInterface(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value interface{}, exists bool) (interface{}, bool)) (interface{}, interface{}, bool)
	Del(keys ...string) error
}

type ICache interface {
	Set(key string, value interface{}, expiration time.Duration) (string, error)
	GetString(key string) (string, bool)
	GetInt(key string) (int, bool)
	GetInterface(key string) (interface{}, bool)
	HMSet(key string, fields map[string]interface{}) (string, error)
	GetMapOrUpdate(key string, f func() (map[string]interface{}, error)) (map[string]interface{}, error)
	GetIntOrUpdate(key string, valueExpiration time.Duration, f func() (int, error)) (int, error)
	GetStringOrUpdate(key string, valueExpiration time.Duration, f func() (string, error)) (string, error)
	GetAndSetAtomicString(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value string, exists bool) (string, bool)) (string, string, bool)
	GetAndSetAtomicInt(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value int, exists bool) (int, bool)) (int, int, bool)
	GetAndSetAtomicInterface(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value interface{}, exists bool) (interface{}, bool)) (interface{}, interface{}, bool)
	SMembers(key string) ([]string, bool)
	Del(keys ...string) error
}

type Cache struct {
	CacheSource ICacheSource
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return c.CacheSource.Set(key, value, expiration)
}

func (c *Cache) GetString(key string) (string, bool) {
	return c.CacheSource.GetString(key)
}

func (c *Cache) GetInt(key string) (int, bool) {
	return c.CacheSource.GetInt(key)
}

func (c *Cache) GetInterface(key string) (interface{}, bool) {
	return c.CacheSource.GetInterface(key)
}

func (c *Cache) HMSet(key string, fields map[string]interface{}) (string, error) {
	return c.CacheSource.HMSet(key, fields)
}

func (c *Cache) GetMapOrUpdate(key string, f func() (map[string]interface{}, error)) (map[string]interface{}, error) {
	cache, ok := c.CacheSource.HGetAll(key)

	if ok {
		out := make(map[string]interface{})
		for k, v := range cache {
			out[k] = v
		}
		return out, nil
	}

	out, err := f()

	if err != nil {
		return make(map[string]interface{}), err
	}

	_, _ = c.CacheSource.HMSet(key, out)
	return out, err
}

func (c *Cache) GetIntOrUpdate(key string, valueExpiration time.Duration, f func() (int, error)) (int, error) {
	cache, ok := c.CacheSource.GetInt(key)

	if ok {
		return cache, nil
	}

	out, err := f()

	if err != nil {
		return 0, err
	}

	_, _ = c.CacheSource.Set(key, out, valueExpiration)
	return out, err
}

func (c *Cache) GetStringOrUpdate(key string, valueExpiration time.Duration, f func() (string, error)) (string, error) {
	cache, ok := c.CacheSource.GetString(key)

	if ok {
		return cache, nil
	}

	out, err := f()

	if err != nil {
		return "", err
	}

	_, _ = c.CacheSource.Set(key, out, valueExpiration)
	return out, err
}

func (c *Cache) GetAndSetAtomicString(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value string, exists bool) (string, bool)) (string, string, bool) {
	return c.CacheSource.GetAndSetAtomicString(key, lockExpiration, valueExpiration, cb)
}

func (c *Cache) GetAndSetAtomicInt(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value int, exists bool) (int, bool)) (int, int, bool) {
	return c.CacheSource.GetAndSetAtomicInt(key, lockExpiration, valueExpiration, cb)
}

func (c *Cache) GetAndSetAtomicInterface(key string, lockExpiration time.Duration, valueExpiration time.Duration, cb func(value interface{}, exists bool) (interface{}, bool)) (interface{}, interface{}, bool) {
	return c.CacheSource.GetAndSetAtomicInterface(key, lockExpiration, valueExpiration, cb)
}

func (c *Cache) SMembers(key string) ([]string, bool) {
	return c.CacheSource.SMembers(key)
}

func (c *Cache) Del(keys ...string) error {
	return c.CacheSource.Del(keys...)
}

func (c *Cache) BindKey() string {
	return "cache.service"
}
