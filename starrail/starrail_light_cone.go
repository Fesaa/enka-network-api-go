package starrail

import (
	"fmt"

	"github.com/Fesaa/enka-network-api-go/localization"
)

type LightCone struct {
	SuperImposion int
	Level         int
	// Also ascension
	Promotion int
	Stats     []LightConeStat
	Hash      int64
}

func (lightCone *LightCone) GetHash() string {
	return fmt.Sprint(lightCone.Hash)
}

func (lightCone *LightCone) Name() string {
	return localization.GetHonkaiLocaleOrHash(lightCone)
}

type LightConeStat struct {
	Stat  string
	Value float64
}
