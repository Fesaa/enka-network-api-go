package starrail

import (
	"fmt"

	"github.com/Fesaa/enka-network-api-go/localization"
)

type LightCone struct {
	LightConeID   int
	SuperImposion int
	Level         int
	// Also ascension
	Promotion int
	Stats     []LightConeStat
	Hash      string
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

type LightConeData struct {
	Rarity int `json:"Rarity"`
	// Convert to Path with starrail.PathFromRaw
	RawPath       string                    `json:"AvatarBaseType"`
	EquipmentName localization.HashNameAble `json:"EquipmentName"`
	ImagePath     string                    `json:"ImagePath"`
}
