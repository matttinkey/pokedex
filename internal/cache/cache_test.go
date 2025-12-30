package cache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Second * 5)
	cache.Add("test", []byte("hello"))
	val, ok := cache.Get("test")
	if !ok {
		t.Errorf("expected to find a key")
	}

	if string(val) != "hello" {
		t.Errorf("values did not match")
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(time.Millisecond * 5)
	cache.Add("test", []byte("hello"))
	_, ok := cache.Get("test")
	if !ok {
		t.Errorf("expected to find a key")
	}

	time.Sleep(time.Millisecond * 5)
	_, ok = cache.Get("test")
	if ok {
		t.Errorf("expected to not find a key")
	}
}
