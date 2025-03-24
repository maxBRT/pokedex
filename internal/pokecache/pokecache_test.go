package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cases := []struct {
		input struct {
			key string
			val []byte
		}
		expected bool
	}{
		{
			input: struct {
				key string
				val []byte
			}{"example.com", []byte("testdata")},
			expected: true,
		},
	}
	for _, c := range cases {
		cache := NewCache(5 * time.Second)
		cache.Add(c.input.key, c.input.val)
		_, exist := cache.Get(c.input.key)
		if !exist {
			t.Errorf("Expected true, got %v", exist)
		}
		time.Sleep(6 * time.Second)
		_, exist = cache.Get(c.input.key)
		if exist {
			t.Errorf("Expected false, got %v", exist)
		}
	}
}
