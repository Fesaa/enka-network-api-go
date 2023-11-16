package starrail

import "github.com/Fesaa/enka-network-api-go/localization"

type CharacterData struct {
	// 4 or 5
	Star int `json:"Rarity"`
	Path string
	// Internal name for the Path
	RawPath                 string                    `json:"AvatarBaseType"`
	AvatarName              localization.HashNameAble `json:"AvatarName"`
	AvatarFullName          localization.HashNameAble `json:"AvatarFullname"`
	Element                 string                    `json:"Element"`
	AvatarSideIconPath      string                    `json:"AvatarSideIconPath"`
	AvatarCutinFrontImgPath string                    `json:"AvatarCutinFrontImgPath"`
	RankIdList              []int                     `json:"RankIDList"`
	SkillList               []int                     `json:"SkillList"`
}

// Returns the path used in game, from the raw path.
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
	}
	return rawPath
}

// Returns the localized AvatarName or Hash if not found
func (d *CharacterData) Name() string {
	return localization.GetHonkaiLocaleOrHash(&d.AvatarName)
}

// Returns the localized AvatarFullName or Hash if not found
//
// For some reason always returns the Hash. FullName isn't present in the GitHub file.
func (d *CharacterData) FullName() string {
	return localization.GetHonkaiLocaleOrHash(&d.AvatarFullName)
}
