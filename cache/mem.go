package cache

import (
	"sync"
)

// Mem is a memory cache implementation
type Mem struct {
	sync.Mutex
	values map[string]string
}

func NewMem() *Mem {
	return &Mem{values: make(map[string]string)}
}

func (m *Mem) Set(key, value string) {
	m.Lock()
	defer m.Unlock()
	m.values[key] = value
}

func (m *Mem) Get(key string) string {
	m.Lock()
	defer m.Unlock()

	return m.values[key]
}
