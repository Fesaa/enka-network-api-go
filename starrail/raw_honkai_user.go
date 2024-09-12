package starrail

type RawHonkaiUser struct {
	DetailInfo RawDetailInfo `json:"detailInfo"`
	Ttl        int           `json:"ttl"`
	Uid        string        `json:"uid"`
}

func (user *RawHonkaiUser) GetTtl() int {
	return user.Ttl
}

func (user *RawHonkaiUser) ToUser() *User {
	return UserFromRaw(user)
}

type RawDetailInfo struct {
	HeadIcon         int               `json:"headIcon"`
	AvatarDetailList []RawAvatarDetail `json:"avatarDetailList"`
	NickName         string            `json:"nickname"`
	Level            int               `json:"level"`
	Signature        string            `json:"signature"`
	Uid              int               `json:"uid"`
	FriendCount      int               `json:"friendCount"`
	RecordInfo       RawRecordInfo     `json:"recordInfo"`
	WordlLevel       int               `json:"worldLevel"`
	IsDisplayAvatar  *bool             `json:"isDisplayAvatar,omitempty"`
	Platform         string            `json:"platform"`
}

type RawAvatarDetail struct {
	AvatarId int `json:"avatarId"`
	Rank     int `json:"rank"`
	Level    int `json:"level"`
	// Ascension, 0 if level < 20
	Promotion     int               `json:"promotion,omitempty"`
	Equipment     *RawEquipmentInfo `json:"equipment,ompitempty"`
	SkillTreeList []RawSkillTree    `json:"skillTreeList"`
	RelicList     []RawRelic        `json:"relicList"`
}

type RawRecordInfo struct {
	ChallengeInfo          RawChallengeInfo `json:"challengeInfo"`
	EquipmentCount         int              `json:"equipmentCount"`
	MaxRogueChallengeScore int              `json:"maxRogueChallengeScore"`
	AchievementCount       int              `json:"achievementCount"`
	AvatarCount            int              `json:"avatarCount"`
	BookCount              int              `json:"bookCount"`
	RelicCount             int              `json:"relicCount"`
	MusicCount             int              `json:"musicCount"`
}

type RawChallengeInfo struct {
	ScheduleMaxLevel     int  `json:"scheduleMaxLevel"`
	ScheduleMaxGroupId   *int `json:"scheduleMaxGroupId,omitempty"`
	NoneScheduleMaxLevel *int `json:"noneScheduleMaxLevel,omitempty"`
	AbyssLevel           int  `json:"abyssLevel"`
	AbyssStarCount       int  `json:"abyssStarCount"`
}

type RawEquipmentInfo struct {
	Tid   int `json:"tid"`
	Rank  int `json:"rank"`
	Level int `json:"level"`
	// Ascension
	Promotion         int                  `json:"promotion"`
	EquipmentFlatData RawEquipmentFlatData `json:"_flat"`
}

type RawEquipmentFlatData struct {
	Props []RawEquipmentFlatProp `json:"props"`
	Name  int64                  `json:"name"`
}

type RawSkillTree struct {
	PointId int `json:"pointId"`
	Level   int `json:"level"`
}

type RawEquipmentFlatProp struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type RawRelic struct {
	Level         int              `json:"level"`
	MainAffixId   int              `json:"mainAffixId"`
	Type          int              `json:"type"`
	Tid           int              `json:"tid"`
	RelicFlatData RawRelicFlatData `json:"_flat"`
}

type RawRelicFlatData struct {
	Props   []RawRelicFlatProp `json:"props"`
	SetName int64              `json:"setName"`
	SetId   int                `json:"setId"`
}

type RawRelicFlatProp struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}
