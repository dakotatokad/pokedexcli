package pokecache

import "testing"

func TestCreateCacheNotNil(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("Expected cache to be initialized, but it is nil")
	}
}

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	cache.Add("testKey", []byte("testValue"))
	actual, ok := cache.Get("testKey")
	if !ok {
		t.Error("Expected to find 'testKey' in cache, but it was not found")
	}
	if string(actual) != "testValue" {
		t.Errorf("Expected 'testValue', got '%s'", actual)
	}
}