package cache

// Cache is an interface for cache service
type Cache interface {
	Get(key string) string
	Set(key, value string)
}
