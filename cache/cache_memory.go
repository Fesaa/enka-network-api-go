package cache

import (
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

// In memory cache with maps
// This is not thread save. Be careful
type MemoryCache struct {
	HonkaiUsers           map[string]CachedData[*starrail.RawHonkaiUser]
	StarRailCharacterData map[string]*starrail.CharacterData

	GenshinUsers     map[string]CachedData[*genshin.RawGenshinUser]
	GenshinNameCards map[int]string

	GensshinProfileIcons map[int]string
	MaxGenshinProfileId  int

	GenshinCharacterData map[string]*genshin.CharacterData

	GenshinMaterials map[int]*genshin.Material
}

func newMemoryCache() (*MemoryCache, error) {
	c := &MemoryCache{
		HonkaiUsers:           map[string]CachedData[*starrail.RawHonkaiUser]{},
		StarRailCharacterData: map[string]*starrail.CharacterData{},

		GenshinUsers:     map[string]CachedData[*genshin.RawGenshinUser]{},
		GenshinNameCards: map[int]string{},
	}
	e := c.loadResources()
	if e != nil {
		return nil, e
	}

	return c, nil
}

func (m *MemoryCache) loadResources() error {
	err := m.loadStarRailResources()
	if err != nil {
		return err
	}

	err = m.loadGenshinResources()
	if err != nil {
		return err
	}

	return nil
}
