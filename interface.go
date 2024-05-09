package enkanetworkapigo

import (
	"github.com/Fesaa/enka-network-api-go/starrail"
	"github.com/Fesaa/enka-network-api-go/utils"
)

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
