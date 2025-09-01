package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	NewCache(3 * time.Second)
}

func TestCacheAdd(t *testing.T) {
	cache := NewCache(2 * time.Second)
	cache.Add("taiyaki", []byte{20, 26})
}

func TestCacheGet(t *testing.T) {
	cases := map[string][]byte{
		"takoyaki": {20, 26},
		"taiyaki":  {4, 30, 35},
		"sushi":    nil,
	}

	cache := NewCache(3 * time.Second)

	cache.Add("sushi", cases["sushi"])
	cache.Add("taiyaki", cases["taiyaki"])
	cache.Add("takoyaki", cases["takoyaki"])

	for expected_key, expected_val := range cases {
		val, ok := cache.Get(expected_key)
		if !bytes.Equal(val, expected_val) {
			t.Errorf(
				`failed for "%s", expected: %v, got: %v, ok: %v`,
				expected_key,
				expected_val,
				val,
				ok,
			)
		}
	}
}

func TestReapLoop(t *testing.T) {
	t.Parallel()
	cases := map[string][]byte{
		"bun":  {93, 31},
		"node": {34, 11, 24},
		"deno": nil,
	}
	cache := NewCache(3 * time.Second)

	cache.Add("deno", []byte{3, 3, 80}) // should become nil after timeout
	time.Sleep(2 * time.Second)
	cache.Add("bun", cases["bun"])
	time.Sleep(2 * time.Second)

	cache.Add("node", cases["node"])

	for expected_key, expected_val := range cases {
		val, ok := cache.Get(expected_key)
		if !bytes.Equal(val, expected_val) {
			t.Errorf(
				`failed for "%s", expected: %v, got: %v, ok: %v, time: %v`,
				expected_key,
				expected_val,
				val,
				ok,
				time.Now(),
			)
		}
	}
}
