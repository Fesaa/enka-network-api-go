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
	Excels() HSRExcels
}

// HSRExcels load data almost directly from the ExcelOutput, these are NOT loaded at startup, rather the request is made
// only when a method is called. If this fails, it'll be called each time as the underlying data stays nil.
// We're not returning an error as this isn't expected to error, this make working with them nicer
type HSRExcels interface {
	RelicMainAffix() []starrail.RelicMainAffixConfig
	RelicSubAffixConfig() []starrail.RelicSubAffixConfig
	RelicSetConfig(string) (*starrail.RelicSetConfig, bool)
	RelicConfig(string) (*starrail.RelicConfig, bool)
	SkillTree(string) map[starrail.SkillTreeAnchor]starrail.SkillTreeNode
	MultiplePathAvatarConfig(string) (*starrail.MultiplePathAvatarConfig, bool)
}

type GenshinData interface {
	HasNameCard(int) bool
	NameCardName(int) *string
	ProfileIcon(string) string
	CharacterData(string) *genshin.CharacterData
	Characters() []*genshin.CharacterData
	Material(int) *genshin.RawMaterial
}
