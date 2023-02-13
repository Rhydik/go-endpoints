package mutexmap

import (
	"sync"
)

type MutexMap struct {
	m  map[string]string
	mu sync.RWMutex
}

func (m *MutexMap) get(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.m[key]
	return val, ok
}

func (m *MutexMap) set(key, val string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = val
}

func newMutexMap() *MutexMap {
	return &MutexMap{
		m: make(map[string]string),
	}
}