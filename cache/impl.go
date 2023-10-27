package cache

var cache EnkaCache

type CacheType int

const (
	// In-memory cache
	// Uses maps, so it's not thread-safe
	MEMORY CacheType = iota
)

// Init initializes the cache
// If you've initiated a EnkaNetworkAPI instance, a cache will be automatically created
// Parameters:
//
//	cacheType: The type of cache to use
//
// Returns:
//
//	error: An error if one occurred, nil otherwise
func Init(cacheType CacheType) error {

	switch cacheType {
	case MEMORY:
		c, e := newMemoryCache()
		if e != nil {
			return e
		}
		cache = c
	}

	return nil
}

// Get the current cache instance
func Get() EnkaCache {
	return cache
}
