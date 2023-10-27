package enkanetworkapigo

type HonkaiUser struct {
	DetailInfo DetailInfo `json:"detailInfo"`
	Ttl        int        `json:"ttl"`
	Uid        string     `json:"uid"`
}

type DetailInfo struct {
	HeadIcon         int            `json:"headIcon"`
	AvatarDetailList []AvatarDetail `json:"avatarDetailList"`
	NickName         string         `json:"nickname"`
	Level            int            `json:"level"`
	Signature        string         `json:"signature"`
	Uid              int            `json:"uid"`
	FriendCount      int            `json:"friendCount"`
	RecordInfo       RecordInfo     `json:"recordInfo"`
	WordlLevel       int            `json:"worldLevel"`
	IsDisplayAvatar  bool           `json:"isDisplayAvatar"`
	Platform         string         `json:"platform"`
}

type AvatarDetail struct {
	AvatarId int `json:"avatarId"`
	Rank     int `json:"rank"`
	Level    int `json:"level"`
	// Ascension, nill if level < 20
	Promotion     int           `json:"promotion,omitempty"`
	Equipment     EquipmentInfo `json:"equipment"`
	SkillTreeList []SkillTree   `json:"skillTreeList"`
	RelicList     []Relic       `json:"relicList"`
}

type RecordInfo struct {
	ChallengeInfo          ChallengeInfo `json:"challengeInfo"`
	EquipmentCount         int           `json:"equipmentCount"`
	MaxRogueChallengeScore int           `json:"maxRogueChallengeScore"`
	AchievementCount       int           `json:"achievementCount"`
	AvatarCount            int           `json:"avatarCount"`
}

type ChallengeInfo struct {
	ScheduleMaxLevel     int `json:"scheduleMaxLevel"`
	ScheduleMaxGroupId   int `json:"scheduleMaxGroupId,omitempty"`
	NoneScheduleMaxLevel int `json:"noneScheduleMaxLevel,omitempty"`
}

type EquipmentInfo struct {
	Tid   int `json:"tid"`
	Rank  int `json:"rank"`
	Level int `json:"level"`
	// Ascension
	Promotion         int               `json:"promotion"`
	EquipmentFlatData EquipmentFlatData `json:"_flat"`
}

type EquipmentFlatData struct {
	Props []EquipmentFlatProp `json:"props"`
	Name  int64               `json:"name"`
}

type SkillTree struct {
	PointId int `json:"pointId"`
	Level   int `json:"level"`
}

type EquipmentFlatProp struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type Relic struct {
	Level         int           `json:"level"`
	MainAffixId   int           `json:"mainAffixId"`
	Type          int           `json:"type"`
	Tid           int           `json:"tid"`
	RelicFlatData RelicFlatData `json:"_flat"`
}

type RelicFlatData struct {
	Props   []RelicFlatProp `json:"props"`
	SetName int             `json:"setName"`
	SetId   int             `json:"setId"`
}

type RelicFlatProp struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}
