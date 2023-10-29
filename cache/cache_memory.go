package cache

import (
	"sync"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

// In memory cache with maps
// This is not thread save. Be careful
type MemoryCache struct {
	HonkaiUsers           map[string]CachedData[*starrail.RawHonkaiUser]
	StarRailCharacterData map[string]*starrail.CharacterData
	StarRailAvatars       map[string]string

	GenshinUsers     map[string]CachedData[*genshin.RawGenshinUser]
	GenshinNameCards map[int]string

	GensshinProfileIcons map[int]string
	MaxGenshinProfileId  int

	GenshinCharacterData map[string]*genshin.CharacterData

	GenshinMaterials map[int]genshin.RawMaterial
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
	var SRError error
	var GError error
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		SRError = m.loadStarRailResources()
		wg.Add(-1)
	}()

	go func() {
		GError = m.loadGenshinResources()
		wg.Add(-1)
	}()

	wg.Wait()
	if SRError != nil {
		return SRError
	}

	if GError != nil {
		return GError
	}

	return nil
}
