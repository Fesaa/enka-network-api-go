package cache

var cache EnkaCache

type CacheType int

const (
	MEMORY CacheType = iota
)

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

func Get() EnkaCache {
	return cache
}
