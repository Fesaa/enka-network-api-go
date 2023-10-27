package enkanetworkapigo

import "fmt"

type StarRailLightCone struct {
	SuperImposion int
	Level         int
	// Also ascension
	Promotion int
	Stats     []StarRailLightConeStat
	Hash      int64
}

func (lightCone *StarRailLightCone) GetHash() string {
	return fmt.Sprint(lightCone.Hash)
}

type StarRailLightConeStat struct {
	Stat  string
	Value float64
}
