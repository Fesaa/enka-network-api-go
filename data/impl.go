package data

import (
	"github.com/Fesaa/enka-network-api-go/localization"
	"github.com/rs/zerolog"
	"sync"
)

func New(log zerolog.Logger) (EnkaData, error) {
	return newMemoryCache(log)
}

type memoryCache struct {
	starRailData StarRailData
	genshinData  GenshinData
}

func (m *memoryCache) StarRailData() StarRailData {
	return m.starRailData
}

func (m *memoryCache) GenshinData() GenshinData {
	return m.genshinData
}

func newMemoryCache(log zerolog.Logger) (*memoryCache, error) {
	wg := sync.WaitGroup{}

	var (
		sr  StarRailData
		g   GenshinData
		err error
	)

	if localization.LoadStarRail {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sr, err = newStarRail(log)
		}()
	}

	if localization.LoadGenshin {
		wg.Add(1)
		go func() {
			defer wg.Done()
			g, err = newGenshin(log)
		}()
	}

	wg.Wait()

	if err != nil {
		return nil, err
	}

	c := &memoryCache{
		starRailData: sr,
		genshinData:  g,
	}
	return c, nil
}
