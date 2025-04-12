package data

import "github.com/rs/zerolog"

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
	sr, err := newStarRail(log)
	if err != nil {
		return nil, err
	}

	g, err := newGenshin(log)
	if err != nil {
		return nil, err
	}

	c := &memoryCache{
		starRailData: sr,
		genshinData:  g,
	}
	return c, nil
}
