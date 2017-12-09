package loader

import (
	"sync"
	"strings"

	cachepkg "github.com/vlamug/scfg/cache"
	storagepkg "github.com/vlamug/scfg/storage"
	"github.com/vlamug/scfg/request"
)

// Loader is service for loading config from database using memory cache
type Loader struct {
	storage storagepkg.Storage

	sync.Mutex
	cache cachepkg.Cache
}

// NewLoader creates new load service
func NewLoader(storage storagepkg.Storage, cache cachepkg.Cache) *Loader {
	return &Loader{storage: storage, cache: cache}
}

// Load loads config from cache if exists otherwise from database
func (ld *Loader) Load(req request.GetRequest) string {
	ld.Lock()
	defer ld.Unlock()

	key := strings.Join([]string{req.Type, req.Data}, ":")

	value := ld.cache.Get(key)
	if value != "" {
		return value
	}

	cfg := ld.storage.Get(key)
	ld.cache.Set(key, cfg.CSet)

	return ld.cache.Get(key)
}
