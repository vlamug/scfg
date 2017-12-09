package cache

import (
	"sync"

	"github.com/vlamug/scfg/metrics"
)

// Mem is a memory cache implementation
type Mem struct {
	sync.Mutex
	values map[string]string
}

// NewMem creates new memory cache service
func NewMem() *Mem {
	return &Mem{values: make(map[string]string)}
}

// Set sets new config to memory
func (m *Mem) Set(key, value string) {
	m.Lock()
	defer m.Unlock()
	m.values[key] = value

	metrics.MemCfgStored.Set(float64(len(m.values)))
}

// Get gets config from memory
func (m *Mem) Get(key string) string {
	m.Lock()
	defer m.Unlock()

	return m.values[key]
}
