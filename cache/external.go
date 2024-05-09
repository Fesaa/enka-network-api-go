package cache

import (
	"encoding/json"
	"log/slog"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

type ExternalCache interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
}

type externalCache struct {
	ext ExternalCache
	log *slog.Logger
}

func ExternalCacheWrapper(ext ExternalCache, log *slog.Logger) EnkaHttpCache {
	return &externalCache{ext: ext, log: log}
}

func (c *externalCache) AddGenshinUser(user *genshin.RawGenshinUser) {
	data, err := json.Marshal(user)
	if err != nil {
		c.log.Error("Failed to marshal genshin user", "err", err)
		return
	}

	if err := c.ext.Save(user.Uid, data); err != nil {
		c.log.Error("Failed to save genshin user", "err", err)
	}
}

func (c *externalCache) GetGenshinUser(uid string) *genshin.RawGenshinUser {
	data, err := c.ext.Load(uid)
	if err != nil {
		c.log.Error("Failed to load genshin user", "err", err)
		return nil
	}

	var user genshin.RawGenshinUser
	if err := json.Unmarshal(data, &user); err != nil {
		c.log.Error("Failed to unmarshal genshin user", "err", err)
		return nil
	}

	return &user
}

func (c *externalCache) AddHonkaiUser(user *starrail.RawHonkaiUser) {
	data, err := json.Marshal(user)
	if err != nil {
		c.log.Error("Failed to marshal user", "err", err)
		return
	}

	if err := c.ext.Save(user.Uid, data); err != nil {
		c.log.Error("Failed to save user", "err", err)
	}
}

func (c *externalCache) GetHonkaiUser(uid string) *starrail.RawHonkaiUser {
	data, err := c.ext.Load(uid)
	if err != nil {
		c.log.Error("Failed to load user", "err", err)
		return nil
	}

	var user starrail.RawHonkaiUser
	if err := json.Unmarshal(data, &user); err != nil {
		c.log.Error("Failed to unmarshal user", "err", err)
		return nil
	}

	return &user
}
