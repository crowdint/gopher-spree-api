package cache

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

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

func setup(t *testing.T) {
	sock := fmt.Sprintf("/tmp/test-gomemcache-%d.sock", os.Getpid())
	cmd := exec.Command("memcached", "-s", sock)
	if err := cmd.Start(); err != nil {
		t.Skipf("skipping test; couldn't find memcached")
		return
	}
	defer cmd.Wait()
	defer cmd.Process.Kill()

	// Wait a bit for the socket to appear.
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(sock); err == nil {
			break
		}
		time.Sleep(time.Duration(25*i) * time.Millisecond)
	}

	InitCache(sock)
}

func TestCacheImplementation(t *testing.T) {
	// Init memcached
	setup(t)

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

	SetMultis(setMultis)
	missingItems, err := FindMultis(getMultis)
	errorHandler(t, err, "find multis: %v", err)
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
