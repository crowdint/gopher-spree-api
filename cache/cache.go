package cache

import (
	"errors"

	"github.com/bradfitz/gomemcache/memcache"
)

var (
	cache *memcache.Client
)

type Cacheable interface {
	Key() string
	KeyWithPrefix(string) string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}

func InitCache(servers ...string) {
	if cache == nil {
		cache = memcache.New(servers...)
	}
}

func set(key string, cacheable Cacheable) error {
	if cache == nil {
		return errors.New("Cache is not initialized.")
	}

	data, err := cacheable.Marshal()
	if err != nil {
		return err
	}

	cache.Set(&memcache.Item{Key: key, Value: data})
	return nil
}

func Set(cacheable Cacheable) error {
	return set(cacheable.Key(), cacheable)
}

func SetWithPrefix(prefix string, cacheable Cacheable) error {
	return set(cacheable.KeyWithPrefix(prefix), cacheable)
}

func SetMultis(cacheables []Cacheable) {
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
		return errors.New("Cache is not initialized.")
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

func findMultis(keys []string, cacheables []Cacheable) ([]Cacheable, error) {
	if cache == nil {
		return cacheables, errors.New("Cache is not initialized.")
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

func FindMultis(cacheables []Cacheable) ([]Cacheable, error) {
	keys := []string{}
	for _, cacheable := range cacheables {
		keys = append(keys, cacheable.Key())
	}

	return findMultis(keys, cacheables)
}

func FindMultisWithPrefix(prefix string, cacheables []Cacheable) ([]Cacheable, error) {
	keys := []string{}
	for _, cacheable := range cacheables {
		keys = append(keys, cacheable.KeyWithPrefix(prefix))
	}

	return findMultis(keys, cacheables)
}
