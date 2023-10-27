package enkanetworkapigo

import "fmt"

type StarRailRelicType int

const (
	HEAD StarRailRelicType = iota + 1
	HAND
	BODY
	FEET
	PLANAR
	ROPE
	UNKNOWN
)

func starRailRelicTypeFromId(id int) StarRailRelicType {
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

type StarRailRelic struct {
	Level    int
	Type     StarRailRelicType
	MainStat StarRailRelicStat
	SubStats []StarRailRelicStat
	Hash     int64
}

func (relic *StarRailRelic) GetHash() string {
	return fmt.Sprint(relic.Hash)
}

type StarRailRelicStat struct {
	Stat  string
	Value float64
}
