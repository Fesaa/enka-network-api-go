package cache

import (
	"log/slog"
	"time"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

func Default(log *slog.Logger) EnkaHttpCache {
	return &httpMemoryCache{
		HonkaiUsers:  utils.NewMap[string, CachedData[*starrail.RawHonkaiUser]](),
		GenshinUsers: utils.NewMap[string, CachedData[*genshin.RawGenshinUser]](),
		log:          log,
	}
}

type httpMemoryCache struct {
	HonkaiUsers  *utils.Map[string, CachedData[*starrail.RawHonkaiUser]]
	GenshinUsers *utils.Map[string, CachedData[*genshin.RawGenshinUser]]

	log *slog.Logger
}

func (m *httpMemoryCache) AddGenshinUser(user *genshin.RawGenshinUser) {
	m.GenshinUsers.Set(user.Uid, NewCachedData[*genshin.RawGenshinUser](user))
}

func (m *httpMemoryCache) GetGenshinUser(uid string) *genshin.RawGenshinUser {
	if cache, ok := m.GenshinUsers.Get(uid); ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		m.GenshinUsers.Delete(uid)
		return nil
	}
	return nil
}

func (m *httpMemoryCache) GetHonkaiUser(uid string) *starrail.RawHonkaiUser {
	if cache, ok := m.HonkaiUsers.Get(uid); ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		m.HonkaiUsers.Delete(uid)
		return nil
	}
	return nil
}

func (m *httpMemoryCache) AddHonkaiUser(user *starrail.RawHonkaiUser) {
	m.HonkaiUsers.Set(user.Uid, NewCachedData[*starrail.RawHonkaiUser](user))
}

func (http *httpMemoryCache) cleaner(d time.Duration) {
	for range time.Tick(d) {
		http.log.Debug(" (Honkai) Cleaning cache")
		http.HonkaiUsers.ForEachModifySafe(func(key string, value CachedData[*starrail.RawHonkaiUser]) {
			if value.IsExpired() {
				http.HonkaiUsers.Delete(key)
			}
		})
		http.log.Debug(" (Genshin) Cleaning cache")
		http.GenshinUsers.ForEachModifySafe(func(key string, value CachedData[*genshin.RawGenshinUser]) {
			if value.IsExpired() {
				http.GenshinUsers.Delete(key)
			}
		})
	}
}
