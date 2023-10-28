package cache

import (
	"time"

	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

type EnkaCacheAble interface {
	GetTtl() int
}

type CachedData[T EnkaCacheAble] struct {
	data           T
	expirationTime time.Time
}

func NewCachedData[T EnkaCacheAble](data T) CachedData[T] {
	return CachedData[T]{
		data:           data,
		expirationTime: time.Now().Add(time.Duration(data.GetTtl()) * time.Second),
	}
}

func (c *CachedData[T]) GetData() T {
	return c.data
}

func (c *CachedData[T]) IsExpired() bool {
	return time.Now().After(c.expirationTime)
}

type EnkaCache interface {
	AddHonkaiUser(*starrail.RawHonkaiUser)
	GetHonkaiUser(string) *starrail.RawHonkaiUser

	GetStarRailCharacterData(string) *starrail.CharacterData
	GetAllStarRailCharacters() []*starrail.CharacterData

	HasNameCard(int) bool
	GetNameCardName(int) *string

	AddGenshinUser(*genshin.RawGenshinUser)
	GetGenshinUser(string) *genshin.RawGenshinUser

	GetProfileIcon(int) string

	GetGenshinCharacterData(string) *genshin.CharacterData
	GetAllGenshinCharacterData() []*genshin.CharacterData

	GetGenshinMaterial(int) *genshin.Material
}
