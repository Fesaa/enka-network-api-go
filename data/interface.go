package data

import (
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

type EnkaData interface {
	GetStarRailCharacterData(string) *starrail.CharacterData
	GetAllStarRailCharacters() []*starrail.CharacterData

	GetStarRailAvatarKey(string) string

	GetStarRailRelicData(string) *starrail.RelicData

	GetStarRailLightConeData(string) *starrail.LightConeData

	HasNameCard(int) bool
	GetNameCardName(int) *string

	GetProfileIcon(string) string

	GetGenshinCharacterData(string) *genshin.CharacterData
	GetAllGenshinCharacterData() []*genshin.CharacterData

	GetGenshinMaterial(int) *genshin.RawMaterial
}
