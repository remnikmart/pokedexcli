package pokecache

import (
	"bytes"
	"testing"
	"time"
)

type Cmd string

const (
	Add    Cmd = "Add"
	Get    Cmd = "Get"
	Remove Cmd = "Remove"
	Wait   Cmd = "Wait"
)

func perform(cache *Cache, cmd Cmd, key string, val []byte, waitTime int) []byte {
	empty := []byte{}
	switch cmd {
	case Add:
		cache.Add(key, val)
	case Get:
		res, ok := cache.Get(key)
		if ok {
			return res
		}
	case Remove:
		cache.removeKey(key)
	case Wait:
		time.Sleep(time.Duration(float64(waitTime)*1100) * time.Millisecond)
	}
	return empty
}

func TestPokecacheCrud(t *testing.T) {
	cases := []struct {
		ttl             int
		commands        []Cmd
		keys            []string
		values          [][]byte
		expectedResults [][]byte
		expectedCounts  []int
	}{
		{
			ttl:             1,
			commands:        []Cmd{Add, Wait, Get},
			keys:            []string{"a", "", "a"},
			values:          [][]byte{{1, 2, 3}, {}, {}},
			expectedResults: [][]byte{{}, {}, {}},
			expectedCounts:  []int{1, 0, 0},
		},
		{
			ttl:             100,
			commands:        []Cmd{Add, Get, Remove},
			keys:            []string{"a", "a", "a"},
			values:          [][]byte{{1, 2, 3}, {}, {}},
			expectedResults: [][]byte{{}, {1, 2, 3}, {}},
			expectedCounts:  []int{1, 1, 0},
		},
	}
	for _, c := range cases {
		cache := NewCache(c.ttl)
		for i := range c.commands {
			res := perform(cache, c.commands[i], c.keys[i], c.values[i], c.ttl)
			if !bytes.Equal(res, c.expectedResults[i]) {
				t.Errorf("Wrong result after command %s\n  Expected: %v\n  Actual: %v",
					c.commands[i], c.expectedResults[i], res)
			}
			cnt := cache.Count()
			if cnt != c.expectedCounts[i] {
				t.Errorf("Wrong cache count after command %s\n  Expected: %v\n  Actual: %v",
					c.commands[i], c.expectedCounts[i], cnt)
			}
		}
	}
}
