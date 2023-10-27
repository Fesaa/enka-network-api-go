package enkanetworkapigo

import "time"

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
	AddHonkaiUser(*RawHonkaiUser)
	GetHonkaiUser(string) *RawHonkaiUser

	GetStarRailCharacterData(string) *StarRailCharacterData
	GetAllStarRailCharacters() []*StarRailCharacterData
}
