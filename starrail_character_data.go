package enkanetworkapigo

type StarRailCharacterData struct {
	// 4 or 5
	Star int `json:"Rarity"`
	Path string
	// Internal name for the Path
	RawPath                 string       `json:"AvatarBaseType"`
	AvatarName              HashNameAble `json:"AvatarName"`
	AvatarFullName          HashNameAble `json:"AvatarFullname"`
	Element                 string       `json:"Element"`
	AvatarSideIconPath      string       `json:"AvatarSideIconPath"`
	AvatarCutinFrontImgPath string       `json:"AvatarCutinFrontImgPath"`
	RankIdList              []int        `json:"RankIDList"`
	SkillList               []int        `json:"SkillList"`
}
