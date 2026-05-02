package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(5 * time.Minute)
	key := "https://example.com/location-area"
	val := []byte("cached response")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected cache hit")
	}

	if !bytes.Equal(got, val) {
		t.Errorf("expected %q, got %q", val, got)
	}
}

func TestCacheReapsOldEntries(t *testing.T) {
	const interval = 10 * time.Millisecond
	cache := NewCache(interval)
	key := "https://example.com/location-area"

	cache.Add(key, []byte("cached response"))

	deadline := time.After(5 * interval)
	for {
		_, ok := cache.Get(key)
		if !ok {
			return
		}

		select {
		case <-deadline:
			t.Fatalf("expected cache entry to be reaped")
		default:
			time.Sleep(interval / 2)
		}
	}
}
