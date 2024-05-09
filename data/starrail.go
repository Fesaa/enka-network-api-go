package data

import (
	"github.com/Fesaa/enka-network-api-go/starrail"
)

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

func (m *memoryCache) GetStarRailRelicData(id string) *starrail.RelicData {
	if relic, ok := m.StarRailRelics.Get(id); ok {
		return relic
	}
	return nil
}

func (m *memoryCache) GetStarRailLightConeData(id string) *starrail.LightConeData {
	if lightCone, ok := m.StarRailLightCones.Get(id); ok {
		return lightCone
	}
	return nil
}
