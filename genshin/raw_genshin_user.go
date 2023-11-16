package genshin

type RawGenshinUser struct {
	PlayerInfo     RawPlayerInfo   `json:"playerInfo"`
	AvatarInfoList []RawAvatarInfo `json:"avatarInfoList"`
	Ttl            int             `json:"ttl"`
	Uid            string          `json:"uid"`
}

func (user *RawGenshinUser) GetTtl() int {
	return user.Ttl
}

type RawPlayerInfo struct {
	Nickname                  string              `json:"nickname"`
	Level                     int                 `json:"level"`
	Signature                 *string             `json:"signature"`
	WorldLevel                int                 `json:"worldLevel"`
	NameCardId                int                 `json:"nameCardId"`
	FinishedAchievementsCount int                 `json:"finishedAchievementNum"`
	TowerFloorIndex           int                 `json:"towerFloorIndex"`
	TowerLevelIndex           int                 `json:"towerLevelIndex"`
	ShowAvatarInfoList        []RawShowAvatarInfo `json:"showAvatarInfoList"`
	ShowNameCardIdList        []int               `json:"showNameCardIdList"`
	ProfilePicture            RawProfilePicture   `json:"profilePicture"`
}

type RawAvatarInfo struct {
	AvatarId               int                    `json:"avatarId"`
	PropMap                map[string]interface{} `json:"propMap"`
	TalentIdList           []int                  `json:"talentIdList"`
	FightPropMap           map[string]float64     `json:"fightPropMap"`
	SkillDepotId           int                    `json:"skillDepotId"`
	InherentProudSkillList []int                  `json:"inherentProudSkillList"`
	EquipList              []RawEquipData         `json:"equipList"`
	FetterInfo             RawFetterInfo          `json:"fetterInfo"`
}

type RawShowAvatarInfo struct {
	AvatarId  int `json:"avatarId"`
	Level     int `json:"level"`
	CostumeId int `json:"costumeId"`
}

type RawFetterInfo struct {
	ExpLevel int `json:"expLevel"`
}

type RawEquipData struct {
	ItemId int64 `json:"itemId"`
	// Empty if weapon
	ArtifactData *RawArtifactData `json:"reliquary,omitempty"`
	// Empty if artifact
	WeaponData *RawWeaponData `json:"weapon,omitempty"`
	Flat       RawFlatData    `json:"flat"`
}

type RawArtifactData struct {
	Level            int   `json:"level"`
	MainPropId       int   `json:"mainPropId"`
	AppendPropIdList []int `json:"appendPropIdList"`
}

type RawWeaponData struct {
	Level        int            `json:"level"`
	PromoteLevel int            `json:"promoteLevel"`
	AffixMap     map[string]int `json:"affixMap"`
}

type RawFlatData struct {
	NameTextMapHash    string              `json:"nameTextMapHash"`
	SetNameTextMapHash string              `json:"setNameTextMapHash"`
	RankLevel          int                 `json:"rankLevel"`
	WeaponStats        []RawSubData        `json:"weaponStats"`
	ArtifactMainData   RawArtifactMainData `json:"artifactMainData"`
	ReliquarySubStats  []RawSubData        `json:"reliquarySubStats"`
	Icon               string              `json:"icon"`
	EquipType          string              `json:"equipType"`
}

type RawArtifactMainData struct {
	MainPropId string `json:"mainPropId"`
	StatValue  int64  `json:"statValue"`
}

type RawSubData struct {
	AppendPropId string  `json:"appendPropId"`
	StatValue    float64 `json:"statValue"`
}

// As of version 4.1, HoYo is not storing an avatar ID anymore, but rather a profile picture ID.
// Both fields will continue to exist, as some people MAY not have migrated yet (changed their avatar after 4.1).
type RawProfilePicture struct {
	AvatarId int `json:"avatarId,omitempty"`
	Id       int `json:"id,omitempty"`
	// Fallback if above are empty
	CostumeId int `json:"costume_id,omitempty"`
}
