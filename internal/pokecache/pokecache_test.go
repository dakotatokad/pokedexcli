package pokecache

import (
	"testing"
)

func TestCreateCacheNotNil(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("Expected cache to be initialized, but it is nil")
	}
}

func TestCreateCache(t *testing.T) {
	cache := NewCache()

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
