package cache

import (
	"encoding/json"
	"os"

	"github.com/Fesaa/enka-network-api-go/starrail"
)

func (m *MemoryCache) loadStarRailResources() error {
	file, err := os.ReadFile("resources/honkai_characters.json")
	if err != nil {
		return err
	}

	var starRailCharacterData map[string]*starrail.CharacterData
	err = json.Unmarshal(file, &starRailCharacterData)
	if err != nil {
		return err
	}

	m.StarRailCharacterData = starRailCharacterData
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
