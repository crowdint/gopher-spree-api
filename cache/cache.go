package cache

import (
	"errors"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/crowdint/gopher-spree-api/configs"
)

var (
	cache *memcache.Client

	ErrCacheInit = errors.New("Cache is not initialized.")
)

func init() {
	if cache == nil {
		url := configs.Get(configs.MEMCACHED_URL)
		servers := strings.Split(url, ",")
		cache = memcache.New(servers...)
	}
}

type Cacheable interface {
	Key() string
	KeyWithPrefix(string) string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
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

func fetchMulti(keys []string, cacheables []Cacheable) ([]Cacheable, error) {
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

// FetchMulti expects an array of Cacheable to fill it with data in the memcached server
// It returns an array of Cacheable of the items that are not in the cache and an error.
//
//	cacheables is the array to be fill. This array has to have items that implement the
// 	Cacheable interface.
//
// []Cacheable is the returned array with the missing items from the cacheables array.
// error is the Error returned if any.
func FetchMulti(cacheables []Cacheable) ([]Cacheable, error) {
	keys := []string{}
	for _, cacheable := range cacheables {
		keys = append(keys, cacheable.Key())
	}

	return fetchMulti(keys, cacheables)
}

func FetchMultiWithPrefix(prefix string, cacheables []Cacheable) ([]Cacheable, error) {
	keys := []string{}
	for _, cacheable := range cacheables {
		keys = append(keys, cacheable.KeyWithPrefix(prefix))
	}

	return fetchMulti(keys, cacheables)
}

func Delete(cacheable Cacheable) error {
	if cache == nil {
		return ErrCacheInit
	}
	return cache.Delete(cacheable.Key())
}
