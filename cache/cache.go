package cache

import (
	"errors"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	cache *memcache.Client

	ErrCacheInit = errors.New("Cache is not initialized.")
)

type Cacheable interface {
	Key() string
	KeyWithPrefix(string) string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func Init(servers ...string) {
	if cache == nil {
		cache = memcache.New(servers...)
	}
}

func set(key string, cacheable Cacheable) error {
	if cache == nil {
		return ErrCacheInit
	}

	data, err := cacheable.Marshal()
	if err != nil {
		return err
	}

	return cache.Set(&memcache.Item{Key: key, Value: data, Expiration: 3600})
}

func Set(cacheable Cacheable) error {
	return set(cacheable.Key(), cacheable)
}

func SetWithPrefix(prefix string, cacheable Cacheable) error {
	return set(cacheable.KeyWithPrefix(prefix), cacheable)
}

func SetMulti(cacheables []Cacheable) {
	for _, cacheable := range cacheables {
		go set(cacheable.Key(), cacheable)
	}
}

func SetMultisWithPrefix(prefix string, cacheables []Cacheable) {
	for _, cacheable := range cacheables {
		go set(cacheable.KeyWithPrefix(prefix), cacheable)
	}
}

func find(key string, cacheable Cacheable) error {
	if cache == nil {
		return ErrCacheInit
	}

	item, err := cache.Get(cacheable.Key())
	if err != nil {
		return err
	}

	return cacheable.Unmarshal(item.Value)
}

func Find(cacheable Cacheable) error {
	return find(cacheable.Key(), cacheable)
}

func FindWithPrefix(prefix string, cacheable Cacheable) error {
	return find(cacheable.KeyWithPrefix(prefix), cacheable)
}

func findMulti(keys []string, cacheables []Cacheable) ([]Cacheable, error) {
	if cache == nil {
		return cacheables, ErrCacheInit
	}

	items, err := cache.GetMulti(keys)
	if err != nil {
		return cacheables, err
	}

	missingItems := []Cacheable{}
	for i, cacheable := range cacheables {
		if item, ok := items[keys[i]]; ok {
			if err := cacheable.Unmarshal(item.Value); err != nil {
				missingItems = append(missingItems, cacheable)
			}
		} else {
			missingItems = append(missingItems, cacheable)
		}
	}

	return missingItems, nil
}

func FindMulti(cacheables []Cacheable) ([]Cacheable, error) {
	keys := []string{}
	for _, cacheable := range cacheables {
		keys = append(keys, cacheable.Key())
	}

	return findMulti(keys, cacheables)
}

func FindMultiWithPrefix(prefix string, cacheables []Cacheable) ([]Cacheable, error) {
	keys := []string{}
	for _, cacheable := range cacheables {
		keys = append(keys, cacheable.KeyWithPrefix(prefix))
	}

	return findMulti(keys, cacheables)
}

func Delete(cacheable Cacheable) error {
	if cache == nil {
		return ErrCacheInit
	}
	return cache.Delete(cacheable.Key())
}
