package internal

import (
	"sync"
)

type MutexMap struct {
	M  map[string][]string
	mu sync.RWMutex
}

func (m *MutexMap) GetMutexMap(key string) ([]string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.M[key]
	return val, ok
}

func (m *MutexMap) SetMutexMap(key string, val []string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.M[key] = val
}

func NewMutexMap() *MutexMap {
	return &MutexMap{
		M: make(map[string][]string),
	}
}
