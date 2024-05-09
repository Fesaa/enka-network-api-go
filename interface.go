package enkanetworkapigo

import (
	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

type EnkaNetworkAPI interface {
	Cache() cache.EnkaCache
	StarRail() StarRailAPI
	Genshin() GenshinAPI

	SetUserAgent(userAgent string)
}

type StarRailAPI interface {
	Fetch(uid string, success utils.Consumer[*starrail.RawHonkaiUser], failure utils.Consumer[error])
	FetchAndReturn(uid string) (*starrail.RawHonkaiUser, error)

	CharacterData(userCharacter *starrail.UserCharacter) *starrail.CharacterData
	CharacterDataById(uuid string) *starrail.CharacterData

	RelicData(relic *starrail.Relic) *starrail.RelicData
	RelicDataById(relicId string) *starrail.RelicData

	LightConeData(lightCone *starrail.LightCone) *starrail.LightConeData
	LightConeDataById(lightConeId string) *starrail.LightConeData

	Icon(key string) string
	AvatarKey(avatarId string) string
}

type GenshinAPI interface {
	Fetch(uid string, showCaseInfo bool, success utils.Consumer[*genshin.RawGenshinUser], failure utils.Consumer[error])
	FetchAndReturn(uid string, showCaseInfo bool) (*genshin.RawGenshinUser, error)

	CharacterData(userCharacter *genshin.UserCharacter) *genshin.CharacterData
	CharacterDataById(uuid string) *genshin.CharacterData

	Icon(key string) string

	ProfileId(id int) string
	NameCard(nameCardId int) *genshin.NameCard
	Material(id int) *genshin.RawMaterial
}
