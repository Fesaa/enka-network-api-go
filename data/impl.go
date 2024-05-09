package data

import (
	"log/slog"
	"sync"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

func New(log *slog.Logger) (EnkaData, error) {
	return NewMemoryCache(log)
}

type memoryCache struct {
	StarRailCharacterData *utils.Map[string, *starrail.CharacterData]
	StarRailAvatars       *utils.Map[string, string]
	StarRailRelics        *utils.Map[string, *starrail.RelicData]
	StarRailLightCones    *utils.Map[string, *starrail.LightConeData]

	GenshinNameCards     *utils.Map[int, string]
	GenshinProfileIcons  *utils.Map[string, *genshin.ProfilePicture]
	GenshinCharacterData *utils.Map[string, *genshin.CharacterData]
	GenshinMaterials     *utils.Map[int, *genshin.RawMaterial]

	log *slog.Logger
}

func NewMemoryCache(logger *slog.Logger) (*memoryCache, error) {
	c := &memoryCache{
		StarRailCharacterData: utils.NewMap[string, *starrail.CharacterData](),
		GenshinNameCards:      utils.NewMap[int, string](),
		log:                   logger,
	}
	e := c.loadResources()
	if e != nil {
		return nil, e
	}
	return c, nil
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
