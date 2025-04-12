package starrail

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"github.com/Fesaa/enka-network-api-go/localization"
)

type CharacterData struct {
	// 4 or 5
	Star int `json:"Rarity"`
	Path string
	// Internal name for the Path
	RawPath                 string      `json:"AvatarBaseType"`
	AvatarName              hash.String `json:"AvatarName"`
	AvatarFullName          hash.String `json:"AvatarFullname"`
	Element                 string      `json:"Element"`
	AvatarSideIconPath      string      `json:"AvatarSideIconPath"`
	AvatarCutinFrontImgPath string      `json:"AvatarCutinFrontImgPath"`
	RankIdList              []int       `json:"RankIDList"`
	SkillList               []int       `json:"SkillList"`
}

// PathFromRaw returns the path used in game, from the raw path.
// Used internally, you can just use CharacterData.Path
func PathFromRaw(rawPath string) string {
	switch rawPath {
	case "Knight":
		return "Preservation"
	case "Mage":
		return "Erudition"
	case "Priest":
		return "Abundance"
	case "Rogue":
		return "The Hunt"
	case "Shaman":
		return "Harmony"
	case "Warlock":
		return "Nihility"
	case "Warrior":
		return "Destruction"
	case "Memory":
		return "Remembrance"
	}
	return rawPath
}

// Name returns the localized AvatarName or Hash if not found
func (d *CharacterData) Name() string {
	return localization.GetHsrLocaleOrHash(&d.AvatarName)
}

// FullName returns the localized AvatarFullName or Hash if not found
//
// For some reason always returns the Hash. FullName isn't present in the GitHub file.
func (d *CharacterData) FullName() string {
	return localization.GetHsrLocaleOrHash(&d.AvatarFullName)
}
