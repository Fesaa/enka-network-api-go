package cache

import (
	"github.com/rs/zerolog"
	"time"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

func Default(log zerolog.Logger) EnkaHttpCache {
	return &httpMemoryCache{
		hsrUsers:     utils.NewMap[string, CachedData[*starrail.RawUser]](),
		GenshinUsers: utils.NewMap[string, CachedData[*genshin.RawUser]](),
		log:          log.With().Str("handler", "http-cache").Logger(),
	}
}

type httpMemoryCache struct {
	hsrUsers     *utils.Map[string, CachedData[*starrail.RawUser]]
	GenshinUsers *utils.Map[string, CachedData[*genshin.RawUser]]

	log zerolog.Logger
}

func (http *httpMemoryCache) AddGenshinUser(user *genshin.RawUser) {
	http.GenshinUsers.Set(user.Uid, NewCachedData[*genshin.RawUser](user))
}

func (http *httpMemoryCache) GetGenshinUser(uid string) *genshin.RawUser {
	if cache, ok := http.GenshinUsers.Get(uid); ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		http.GenshinUsers.Delete(uid)
		return nil
	}
	return nil
}

func (http *httpMemoryCache) GetHsrUser(uid string) *starrail.RawUser {
	if cache, ok := http.hsrUsers.Get(uid); ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		http.hsrUsers.Delete(uid)
		return nil
	}
	return nil
}

func (http *httpMemoryCache) AddHSRUser(user *starrail.RawUser) {
	http.hsrUsers.Set(user.Uid, NewCachedData[*starrail.RawUser](user))
}

func (http *httpMemoryCache) cleaner(d time.Duration) {
	for range time.Tick(d) {
		http.log.Debug().Str("game", "hsr").Msg("cleaning hsr cache")
		http.hsrUsers.ForEachModifySafe(func(key string, value CachedData[*starrail.RawUser]) {
			if value.IsExpired() {
				http.hsrUsers.Delete(key)
			}
		})
		http.log.Debug().Str("game", "genshin").Msg("cleaning genshin cache")
		http.GenshinUsers.ForEachModifySafe(func(key string, value CachedData[*genshin.RawUser]) {
			if value.IsExpired() {
				http.GenshinUsers.Delete(key)
			}
		})
	}
}
