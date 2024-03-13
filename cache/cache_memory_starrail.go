package cache

import (
	_ "embed"
	"encoding/json"

	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

func (m *memoryCache) loadStarRailResources() error {
	var starRailCharacterData map[string]*starrail.CharacterData
	err := json.Unmarshal(starRailCharacterJson, &starRailCharacterData)
	if err != nil {
		return err
	}

	type Icon struct {
		Icon string `json:"icon"`
	}
	var starRailAvatars map[string]*Icon
	err = json.Unmarshal(starRailAvataJson, &starRailAvatars)
	if err != nil {
		return err
	}

	var avatars map[string]string = make(map[string]string, len(starRailAvatars))
	for k, v := range starRailAvatars {
		avatars[k] = v.Icon
	}

	m.StarRailCharacterData = utils.FromMap(starRailCharacterData)
	m.StarRailAvatars = utils.FromMap(avatars)
	return nil
}

func (m *memoryCache) GetHonkaiUser(uid string) *starrail.RawHonkaiUser {
	if cache, ok := m.HonkaiUsers.Get(uid); ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		m.HonkaiUsers.Delete(uid)
		return nil
	}
	return nil
}

func (m *memoryCache) AddHonkaiUser(user *starrail.RawHonkaiUser) {
	m.HonkaiUsers.Set(user.Uid, NewCachedData[*starrail.RawHonkaiUser](user))
}

func (m *memoryCache) GetStarRailCharacterData(uid string) *starrail.CharacterData {
	if data, ok := m.StarRailCharacterData.Get(uid); ok {
		data.Path = starrail.PathFromRaw(data.RawPath)
		return data
	}

	return nil
}

func (m *memoryCache) GetAllStarRailCharacters() []*starrail.CharacterData {
	s := make([]*starrail.CharacterData, 0, m.StarRailCharacterData.Len())
	m.StarRailCharacterData.ForEach(func(_ string, c *starrail.CharacterData) {
		s = append(s, c)
	})
	return s
}

func (m *memoryCache) GetStarRailAvatarKey(id string) string {
	if avatar, ok := m.StarRailAvatars.Get(id); ok {
		return avatar
	}
	return id
}
