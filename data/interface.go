package data

import (
	"github.com/Fesaa/enka-network-api-go/genshin"
	"github.com/Fesaa/enka-network-api-go/starrail"
)

type EnkaData interface {
	StarRailData() StarRailData
	GenshinData() GenshinData
}

type StarRailData interface {
	CharacterData(string) *starrail.CharacterData
	Characters() []*starrail.CharacterData
	AvatarKey(string) string
	RelicData(string) *starrail.RelicData
	LightConeData(string) *starrail.LightConeData
	SkillTree(string) map[starrail.SkillTreeAnchor]starrail.SkillTreeNode
}

type GenshinData interface {
	HasNameCard(int) bool
	NameCardName(int) *string
	ProfileIcon(string) string
	CharacterData(string) *genshin.CharacterData
	Characters() []*genshin.CharacterData
	Material(int) *genshin.RawMaterial
}
