package cache

import (
	"encoding/json"
	"os"

	"github.com/Fesaa/enka-network-api-go/starrail"
)

// In memory cache with maps
// This is not thread save. Be careful
type MemoryCache struct {
	HonkaiUsers   map[string]CachedData[*starrail.RawHonkaiUser]
	CharacterData map[string]*starrail.CharacterData
}

func newMemoryCache() (*MemoryCache, error) {
	c := &MemoryCache{
		HonkaiUsers:   map[string]CachedData[*starrail.RawHonkaiUser]{},
		CharacterData: map[string]*starrail.CharacterData{},
	}
	e := c.loadResources()
	if e != nil {
		return nil, e
	}

	return c, nil
}

func (m *MemoryCache) loadResources() error {
	file, err := os.ReadFile("resources/honkai_characters.json")
	if err != nil {
		return err
	}

	var starRailCharacterData map[string]*starrail.CharacterData
	err = json.Unmarshal(file, &starRailCharacterData)
	if err != nil {
		return err
	}

	m.CharacterData = starRailCharacterData
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
	if data, ok := m.CharacterData[uid]; ok {
		return data
	}

	return nil
}

func (m *MemoryCache) GetAllStarRailCharacters() []*starrail.CharacterData {
	s := make([]*starrail.CharacterData, 0, len(m.CharacterData))

	for _, c := range m.CharacterData {
		s = append(s, c)
	}
	return s
}
