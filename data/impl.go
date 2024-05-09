package data

import (
	"log/slog"
)

func New(log *slog.Logger) (EnkaData, error) {
	return NewMemoryCache(log)
}

type memoryCache struct {
	starRailData StarRailData
	genshinData  GenshinData
	log          *slog.Logger
}

func (m *memoryCache) StarRailData() StarRailData {
	return m.starRailData
}

func (m *memoryCache) GenshinData() GenshinData {
	return m.genshinData
}

func NewMemoryCache(logger *slog.Logger) (*memoryCache, error) {
	sr, err := newStarRail()
	if err != nil {
		return nil, err
	}

	g, err := newGenshin()
	if err != nil {
		return nil, err
	}

	c := &memoryCache{
		starRailData: sr,
		genshinData:  g,
		log:          logger,
	}
	return c, nil
}
