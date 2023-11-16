package genshin

import "github.com/Fesaa/enka-network-api-go/localization"

type Weapon struct {
	NameHash   string
	Level      int
	Ascension  int
	Refinement int
	Icon       string
	Stats      []WeaponStat
}

func (w *Weapon) GetHash() string {
	return w.NameHash
}

func (w *Weapon) Name() string {
	return localization.GetGenshinLocaleOrHash(w)
}

type WeaponStat struct {
	Stat  string
	Value float64
}

func fromRawData(raw RawEquipData, flat RawFlatData) Weapon {
	rawWeaponData := raw.WeaponData
	weaponStats := make([]WeaponStat, len(flat.WeaponStats))

	for i, stat := range flat.WeaponStats {
		weaponStats[i] = WeaponStat{
			Stat:  stat.AppendPropId,
			Value: float64(stat.StatValue),
		}
	}

	return Weapon{
		NameHash:   flat.NameTextMapHash,
		Level:      rawWeaponData.Level,
		Ascension:  rawWeaponData.PromoteLevel,
		Refinement: resolveFirst(rawWeaponData.AffixMap),
		Icon:       flat.Icon,
		Stats:      weaponStats,
	}
}

func resolveFirst(affixMap map[string]int) int {
	if affixMap == nil {
		return 1
	}
	for _, value := range affixMap {
		return value + 1
	}
	return -1
}
