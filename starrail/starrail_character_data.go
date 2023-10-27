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
