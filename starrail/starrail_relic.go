package starrail

import (
	"fmt"

	"github.com/Fesaa/enka-network-api-go/localization"
)

type RelicType int

const (
	HEAD RelicType = iota + 1
	HAND
	BODY
	FEET
	PLANAR
	ROPE
	UNKNOWN
)

func starRailRelicTypeFromId(id int) RelicType {
	switch id {
	case 1:
		return HEAD
	case 2:
		return HAND
	case 3:
		return BODY
	case 4:
		return FEET
	case 5:
		return PLANAR
	case 6:
		return ROPE
	default:
		return UNKNOWN
	}
}

func StarRailRelicTypeFromString(t string) RelicType {
	switch t {
	case "Head":
		return HEAD
	case "Hand":
		return HAND
	case "Body":
		return BODY
	case "Feet":
		return FEET
	case "Planar":
		return PLANAR
	case "Rope":
		return ROPE
	default:
		return UNKNOWN
	}
}

type Relic struct {
	RelicID  int
	Level    int
	Type     RelicType
	MainStat RelicStat
	SubStats []RelicStat
	Hash     int64
}

func (relic *Relic) GetHash() string {
	return fmt.Sprint(relic.Hash)
}

// Returns the localized name or hash
func (relic *Relic) Name() string {
	return localization.GetHonkaiLocaleOrHash(relic)
}

type RelicStat struct {
	Stat  string
	Value float64
}

type RelicData struct {
	Rarity         int    `json:"Rarity"`
	Type           string `json:"Type"`
	MainAffixGroup int    `json:"MainAffixGroup"`
	SubAffixGroup  int    `json:"SubAffixGroup"`
	Icon           string `json:"Icon"`
	SetId          int    `json:"SetID"`
}
