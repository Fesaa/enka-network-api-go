package cache

import (
	"log/slog"
	"sync"
	"time"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

// In memory cache with maps
// This is not thread save. Be careful
type memoryCache struct {
	HonkaiUsers           *utils.Map[string, CachedData[*starrail.RawHonkaiUser]]
	StarRailCharacterData *utils.Map[string, *starrail.CharacterData]
	StarRailAvatars       *utils.Map[string, string]

	GenshinUsers     *utils.Map[string, CachedData[*genshin.RawGenshinUser]]
	GenshinNameCards *utils.Map[int, string]

	GenshinProfileIcons *utils.Map[int, *genshin.ProfilePicture]
	MaxGenshinProfileId int

	GenshinCharacterData *utils.Map[string, *genshin.CharacterData]

	GenshinMaterials *utils.Map[int, *genshin.RawMaterial]

	log *slog.Logger
}

func NewMemoryCache(logger *slog.Logger, ds ...time.Duration) (*memoryCache, error) {
	c := &memoryCache{
		HonkaiUsers:           utils.NewMap[string, CachedData[*starrail.RawHonkaiUser]](),
		StarRailCharacterData: utils.NewMap[string, *starrail.CharacterData](),

		GenshinUsers:     utils.NewMap[string, CachedData[*genshin.RawGenshinUser]](),
		GenshinNameCards: utils.NewMap[int, string](),
		log:              logger,
	}
	e := c.loadResources()
	if e != nil {
		return nil, e
	}

	d := time.Duration(5) * time.Minute
	if len(ds) > 0 {
		d = ds[0]
	}

	go c.cleaner(d)

	return c, nil
}

func (m *memoryCache) cleaner(d time.Duration) {
	for range time.Tick(d) {
		m.log.Info(" (Honkai) Cleaning cache")
		m.HonkaiUsers.ForEachModifySafe(func(key string, value CachedData[*starrail.RawHonkaiUser]) {
			if value.IsExpired() {
				m.HonkaiUsers.Delete(key)
			}
		})
		m.log.Info(" (Genshin) Cleaning cache")
		m.GenshinUsers.ForEachModifySafe(func(key string, value CachedData[*genshin.RawGenshinUser]) {
			if value.IsExpired() {
				m.GenshinUsers.Delete(key)
			}
		})
	}
}

func (m *memoryCache) loadResources() error {
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
