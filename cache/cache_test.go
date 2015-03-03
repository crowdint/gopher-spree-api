package cache

import (
	"testing"

	"github.com/bradfitz/gomemcache/memcache"
)

type CacheItem struct {
	Value string
}

func (item *CacheItem) Key() string {
	return "CacheItem"
}

func (item *CacheItem) KeyWithPrefix(prefix string) string {
	return prefix + "CacheItem"
}

func (item *CacheItem) Marshal() ([]byte, error) {
	return []byte(item.Value), nil
}

func (item *CacheItem) Unmarshal(data []byte) error {
	item.Value = string(data)
	return nil
}

func TestCacheImplementation(t *testing.T) {
	// Set
	foo := &CacheItem{"foo"}
	err := Set(foo)
	errorHandler(t, err, "first set(foo): %v", err)
	err = Set(foo)
	errorHandler(t, err, "second set(foo): %v", err)

	// Get
	cachedFoo := &CacheItem{}
	err = Find(cachedFoo)
	errorHandler(t, err, "find(foo): %v", err)
	if cachedFoo.Value != foo.Value {
		t.Errorf("The cached foo value is %s, but should be %s", cachedFoo.Value, foo.Value)
	}

	// SetMultis && FindMultis
	setMultis := []Cacheable{
		&CacheItem{"bar"},
		&CacheItem{"waka"},
	}
	getMultis := []Cacheable{}
	for i := 0; i < len(setMultis); i++ {
		getMultis = append(getMultis, &CacheItem{})
	}

	SetMulti(setMultis)
	missingItems, err := FetchMulti(getMultis)
	errorHandler(t, err, "find multi: %v", err)
	if len(missingItems) > 0 {
		t.Errorf("There are some missing items %+v", missingItems)
	}

	// Delete
	err = Delete(foo)
	errorHandler(t, err, "delete(foo): %v", err)
	err = Find(foo)
	if err != memcache.ErrCacheMiss {
		t.Errorf("delete(foo) is not missing")
	}
}

func errorHandler(t *testing.T, err error, format string, args ...interface{}) {
	if err != nil {
		t.Fatalf(format, args...)
	}
}
