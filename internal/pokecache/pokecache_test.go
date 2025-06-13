package pokecache

import (
	"testing"
	"time"
)

func TestCreateCacheNotNil(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("Expected cache to be initialized, but it is nil")
	}
}

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{inputKey: "testKey", inputVal: []byte("testValue")},
		{inputKey: "anotherKey", inputVal: []byte("anotherValue")},
		{inputKey: "emptyKey", inputVal: []byte("")},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("Expected to find '%s' in cache, but it was not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("Expected '%s', got '%s'", cas.inputVal, actual)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("testKey", []byte("testValue"))

	time.Sleep(interval + time.Millisecond * 5)

	_, ok := cache.Get("testKey")
	if ok {
		t.Errorf("Expected 'testKey' to be removed from cache after reap, but it was still found")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("testKey", []byte("testValue"))

	time.Sleep(interval / 2)

	_, ok := cache.Get("testKey")
	if !ok {
		t.Errorf("Expected 'testKey' to be in cache but it was reaped")
	}
}