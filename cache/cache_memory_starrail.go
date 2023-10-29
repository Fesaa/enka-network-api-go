package cache

import (
	_ "embed"
	"encoding/json"

	"github.com/Fesaa/enka-network-api-go/starrail"
)

func (m *MemoryCache) loadStarRailResources() error {
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

	m.StarRailCharacterData = starRailCharacterData
	m.StarRailAvatars = avatars
	return nil
}

func (m *MemoryCache) GetHonkaiUser(uid string) *starrail.RawHonkaiUser {
	if cache, ok := m.HonkaiUsers[uid]; ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		delete(m.HonkaiUsers, uid)
		return nil
	}
	return nil
}

func (m *MemoryCache) AddHonkaiUser(user *starrail.RawHonkaiUser) {
	m.HonkaiUsers[user.Uid] = NewCachedData[*starrail.RawHonkaiUser](user)
}

func (m *MemoryCache) GetStarRailCharacterData(uid string) *starrail.CharacterData {
	if data, ok := m.StarRailCharacterData[uid]; ok {
		return data
	}

	return nil
}

func (m *MemoryCache) GetAllStarRailCharacters() []*starrail.CharacterData {
	s := make([]*starrail.CharacterData, 0, len(m.StarRailCharacterData))

	for _, c := range m.StarRailCharacterData {
		s = append(s, c)
	}
	return s
}

func (m *MemoryCache) GetStarRailAvatarKey(id string) string {
	if avatar, ok := m.StarRailAvatars[id]; ok {
		return avatar
	}
	return id
}
