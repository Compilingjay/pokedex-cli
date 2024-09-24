package pokecache

import (
	"fmt"
	"testing"
	"time"
)

const (
	interval = 5 * time.Second
	baseURL  = "https://pokeapi.co/api/v2"

	baseTime = 5 * time.Millisecond
	waitTime = baseTime + 5*time.Millisecond
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: fmt.Sprintf("%s/location-area", baseURL),
			val: []byte("testdata"),
		},
		{
			key: fmt.Sprintf("%s/location-area/eterna-city-area", baseURL),
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(baseTime)
	cache.Add("https://www.youtube.com/watch?v=jNQXAC9IVRw", []byte("Me at the zoo"))

	_, ok := cache.Get("https://www.youtube.com/watch?v=jNQXAC9IVRw")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://www.youtube.com/watch?v=jNQXAC9IVRw")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
