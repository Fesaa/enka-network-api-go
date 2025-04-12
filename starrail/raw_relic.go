package starrail

import (
	"github.com/Fesaa/enka-network-api-go/hash"
)

type (
	RelicMainAffixConfigProperty string
	RelicRarity                  string
	RelicTypeString              string
	RelicMode                    string
)

const (
	HeadRelicType         RelicTypeString = "HEAD"
	HandsRelicType        RelicTypeString = "HAND"
	BodyRelicType         RelicTypeString = "BODY"
	FeetRelicType         RelicTypeString = "FOOT"
	PlanarSphereRelicType RelicTypeString = "NECK"
	LinkRopeRelicType     RelicTypeString = "OBJECT"

	HPDelta                   RelicMainAffixConfigProperty = "HPDelta"
	AttackDelta               RelicMainAffixConfigProperty = "AttackDelta"
	HPAddedRatio              RelicMainAffixConfigProperty = "HPAddedRatio"
	AttackAddedRatio          RelicMainAffixConfigProperty = "AttackAddedRatio"
	DefenceAddedRatio         RelicMainAffixConfigProperty = "DefenceAddedRatio"
	CriticalChanceBase        RelicMainAffixConfigProperty = "CriticalChanceBase"
	CriticalDamageBase        RelicMainAffixConfigProperty = "CriticalDamageBase"
	HealRatioBase             RelicMainAffixConfigProperty = "HealRatioBase"
	SpeedDelta                RelicMainAffixConfigProperty = "SpeedDelta"
	StatusProbabilityBase     RelicMainAffixConfigProperty = "StatusProbabilityBase"
	PhysicalAddedRatio        RelicMainAffixConfigProperty = "PhysicalAddedRatio"
	FireAddedRatio            RelicMainAffixConfigProperty = "FireAddedRatio"
	IceAddedRatio             RelicMainAffixConfigProperty = "IceAddedRatio"
	ThunderAddedRatio         RelicMainAffixConfigProperty = "ThunderAddedRatio"
	WindAddedRatio            RelicMainAffixConfigProperty = "WindAddedRatio"
	QuantumAddedRatio         RelicMainAffixConfigProperty = "QuantumAddedRatio"
	ImaginaryAddedRatio       RelicMainAffixConfigProperty = "ImaginaryAddedRatio"
	BreakDamageAddedRatioBase RelicMainAffixConfigProperty = "BreakDamageAddedRatioBase"
	SPRatioBase               RelicMainAffixConfigProperty = "SPRatioBase"
)

type RelicMainAffixConfig struct {
	GroupID   int                          `json:"GroupID"`
	AffixID   int                          `json:"AffixID"`
	Property  RelicMainAffixConfigProperty `json:"Property"`
	BaseValue Value                        `json:"BaseValue"`
	LevelAdd  Value                        `json:"LevelAdd"`
}

type RelicSubAffixConfig struct {
	GroupID   int                          `json:"GroupID"`
	AffixID   int                          `json:"AffixID"`
	Property  RelicMainAffixConfigProperty `json:"Property"`
	BaseValue Value                        `json:"BaseValue"`
	StepValue Value                        `json:"StepValue"`
	StepNum   int                          `json:"StepNum"`
}

type RelicSetConfig struct {
	SetID                int         `json:"SetID"`
	SetSkillList         []int       `json:"SetSkillList"`
	SetIconPath          string      `json:"SetIconPath"`
	SetIconFigurePath    string      `json:"SetIconFigurePath"`
	SetName              hash.UInt64 `json:"SetName"`
	DisplayItemID        int         `json:"DisplayItemID"`
	DisplayItemIDRarity4 int         `json:"DisplayItemIDRarity4"`
	Release              bool        `json:"Release"`
	ReleaseVersion       string      `json:"ReleaseVersion"`
}

type RelicConfig struct {
	ID             int             `json:"ID"`
	SetID          int             `json:"SetID"`
	Type           RelicTypeString `json:"Type"`
	Rarity         RelicRarity     `json:"Rarity"`
	MainAffixGroup int             `json:"MainAffixGroup"`
	SubAffixGroup  int             `json:"SubAffixGroup"`
	MaxLevel       int             `json:"MaxLevel"`
	ExpType        int             `json:"ExpType"`
	ExpProvide     int             `json:"ExpProvide"`
	CoinCost       int             `json:"CoinCost"`
	Mode           RelicMode       `json:"Mode"`
}
