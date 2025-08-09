package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache CacheEentry
	mu    sync.Mutex
}

type CacheEentry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {

}

func (g Cache) Get(key string) ([]byte, bool) {

}

func (r Cache) reapLoop() {

}
