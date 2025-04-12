package cache

import (
	"encoding/json"
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/rs/zerolog"
)

type ExternalCache interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
}

type externalCache struct {
	ext ExternalCache
	log zerolog.Logger
}

func ExternalCacheWrapper(ext ExternalCache, log zerolog.Logger) EnkaHttpCache {
	return &externalCache{ext: ext, log: log}
}

func (c *externalCache) AddGenshinUser(user *genshin.RawUser) {
	data, err := json.Marshal(user)
	if err != nil {
		c.log.Error().Err(err).Msg("Failed to marshal genshin user")
		return
	}

	if err = c.ext.Save(user.Uid, data); err != nil {
		c.log.Error().Err(err).Msg("Failed to save genshin user")
	}
}

func (c *externalCache) GetGenshinUser(uid string) *genshin.RawUser {
	data, err := c.ext.Load(uid)
	if err != nil {
		c.log.Error().Err(err).Msg("Failed to get genshin user")
		return nil
	}

	var user genshin.RawUser
	if err = json.Unmarshal(data, &user); err != nil {
		c.log.Error().Err(err).Msg("Failed to unmarshal genshin user")
		return nil
	}

	return &user
}

func (c *externalCache) AddHSRUser(user *starrail.RawUser) {
	data, err := json.Marshal(user)
	if err != nil {
		c.log.Error().Err(err).Msg("Failed to marshal hsr user")
		return
	}

	if err = c.ext.Save(user.Uid, data); err != nil {
		c.log.Error().Err(err).Msg("Failed to save user")
	}
}

func (c *externalCache) GetHsrUser(uid string) *starrail.RawUser {
	data, err := c.ext.Load(uid)
	if err != nil {
		c.log.Error().Err(err).Msg("Failed to load user")
		return nil
	}

	var user starrail.RawUser
	if err = json.Unmarshal(data, &user); err != nil {
		c.log.Error().Err(err).Msg("Failed to marshal hsr user")
		return nil
	}

	return &user
}
