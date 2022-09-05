package anycache

import "sync"

type Map struct {
	store map[int]string
	lock  sync.RWMutex
}

func NewMap() Map {
	return Map{
		store: map[int]string{},
		lock:  sync.RWMutex{},
	}
}

func (m *Map) Get(key int) string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.store[key]
}

func (m *Map) ReloadSnapshot(s map[int]string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.store = s
}
