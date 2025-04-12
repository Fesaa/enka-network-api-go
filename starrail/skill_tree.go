package starrail

type SkillTreeAnchor string

const (
	AnchorBasic      SkillTreeAnchor = "Point01"
	AnchorSkill                      = "Point02"
	AnchorUlt                        = "Point03"
	AnchorTalent                     = "Point04"
	AnchorTechnique                  = "Point05"
	AnchorMajor1                     = "Point06"
	AnchorMajor2                     = "Point07"
	AnchorMajor3                     = "Point08"
	AnchorMinor1                     = "Point09"
	AnchorMinor2                     = "Point10"
	AnchorMinor3                     = "Point11"
	AnchorMinor4                     = "Point12"
	AnchorMinor5                     = "Point13"
	AnchorMinor6                     = "Point14"
	AnchorMinor7                     = "Point15"
	AnchorMinor8                     = "Point16"
	AnchorMinor9                     = "Point17"
	AnchorMinor10                    = "Point18"
	AnchorMemoSkill                  = "Point19"
	AnchorMemoTalent                 = "Point20"
)

type SkillTreeNode struct {
	PointID              int             `json:"PointID"`
	Level                int             `json:"Level"`
	AvatarID             int             `json:"AvatarID"`
	PointType            int             `json:"PointType"`
	Anchor               SkillTreeAnchor `json:"Anchor"`
	MaxLevel             int             `json:"MaxLevel"`
	DefaultUnlock        bool            `json:"DefaultUnlock"`
	PrePoint             []int           `json:"PrePoint"`
	StatusAddList        []StatusAdd     `json:"StatusAddList"`
	MaterialList         []Material      `json:"MaterialList"`
	AvatarPromotionLimit int             `json:"AvatarPromotionLimit"`
}

type StatusAdd struct {
	PropertyType string `json:"PropertyType"`
	Value        Value  `json:"Value"`
}

type Value struct {
	Value float64 `json:"Value"`
}

type Material struct {
	ItemID  int `json:"ItemID"`
	ItemNum int `json:"ItemNum"`
}
