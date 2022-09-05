package anycache

import (
	"sync"
	"sync/atomic"
)

type anyCache struct {
	store atomic.Value
	lock  sync.RWMutex
}

func New() anyCache {
	return anyCache{
		store: atomic.Value{},
		lock:  sync.RWMutex{},
	}
}

func (a *anyCache) Get(key int) any {
	m := a.store.Load().(map[any]any)
	return m[key]
}

func (a *anyCache) ReloadSnapshot(T any) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.store.Store(T)
}
