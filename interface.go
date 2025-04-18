package enkanetworkapigo

import (
	"github.com/Fesaa/enka-network-api-go/cache"
	"github.com/Fesaa/enka-network-api-go/data"
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
	"net/http"
)

type EnkaNetworkAPI interface {
	Data() data.EnkaData
	Cache() cache.EnkaHttpCache
	StarRail() StarRailAPI
	Genshin() GenshinAPI

	SetUserAgent(userAgent string)
	GetUserAgent() string
	HttpClient() *http.Client
}

type StarRailAPI interface {
	Fetch(uid string, success utils.Consumer[*starrail.RawUser], failure utils.Consumer[error])
	FetchAndReturn(uid string) (*starrail.RawUser, error)

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
	Fetch(uid string, showCaseInfo bool, success utils.Consumer[*genshin.RawUser], failure utils.Consumer[error])
	FetchAndReturn(uid string, showCaseInfo bool) (*genshin.RawUser, error)

	CharacterData(userCharacter *genshin.UserCharacter) *genshin.CharacterData
	CharacterDataById(uuid string) *genshin.CharacterData

	Icon(key string) string

	ProfileId(id int) string
	NameCard(nameCardId int) *genshin.NameCard
	Material(id int) *genshin.RawMaterial
}
