package enkanetworkapigo

import (
	"encoding/json"
	"os"
)

// In memory cache with maps
// This is not thread save. Be careful
type MemoryCache struct {
	HonkaiUsers           map[string]CachedData[*RawHonkaiUser]
	StarRailCharacterData map[string]*StarRailCharacterData
}

func NewMemoryCache() (*MemoryCache, error) {
	c := &MemoryCache{
		HonkaiUsers:           map[string]CachedData[*RawHonkaiUser]{},
		StarRailCharacterData: map[string]*StarRailCharacterData{},
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

	var starRailCharacterData map[string]*StarRailCharacterData
	err = json.Unmarshal(file, &starRailCharacterData)
	if err != nil {
		return err
	}

	m.StarRailCharacterData = starRailCharacterData
	return nil
}

func (m *MemoryCache) GetHonkaiUser(uid string) *RawHonkaiUser {
	if cache, ok := m.HonkaiUsers[uid]; ok {
		if !cache.IsExpired() {
			return cache.GetData()
		}
		delete(m.HonkaiUsers, uid)
		return nil
	}
	return nil
}

func (m *MemoryCache) AddHonkaiUser(user *RawHonkaiUser) {
	m.HonkaiUsers[user.Uid] = NewCachedData[*RawHonkaiUser](user)
}

func (m *MemoryCache) GetStarRailCharacterData(uid string) *StarRailCharacterData {
	if data, ok := m.StarRailCharacterData[uid]; ok {
		return data
	}

	return nil
}

func (m *MemoryCache) GetAllStarRailCharacters() []*StarRailCharacterData {
	s := make([]*StarRailCharacterData, 0, len(m.StarRailCharacterData))

	for _, c := range m.StarRailCharacterData {
		s = append(s, c)
	}
	return s
}
