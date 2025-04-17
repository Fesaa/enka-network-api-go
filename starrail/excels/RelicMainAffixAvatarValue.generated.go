package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RelicMainAffixAvatarValue struct {
	Attack            float64 `json:"Attack"`
	AvatarID          float64 `json:"AvatarID"`
	BreakDamage       float64 `json:"BreakDamage"`
	CriticalChance    float64 `json:"CriticalChance"`
	CriticalDamage    float64 `json:"CriticalDamage"`
	DamageAddedRatio  float64 `json:"DamageAddedRatio"`
	Defence           float64 `json:"Defence"`
	HP                float64 `json:"HP"`
	HealRatio         float64 `json:"HealRatio"`
	SPRatio           float64 `json:"SPRatio"`
	Speed             float64 `json:"Speed"`
	StatusProbability float64 `json:"StatusProbability"`
}
type RelicMainAffixAvatarValueAccessor struct {
	_data         []RelicMainAffixAvatarValue
	_dataAvatarID map[float64]RelicMainAffixAvatarValue
}

// LoadData retrieves the data. Must be called before RelicMainAffixAvatarValue.GroupData
func (a *RelicMainAffixAvatarValueAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RelicMainAffixAvatarValue.json")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &a._data)
}

// Raw returns the raw data.
func (a *RelicMainAffixAvatarValueAccessor) Raw() ([]RelicMainAffixAvatarValue, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RelicMainAffixAvatarValue{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RelicMainAffixAvatarValueAccessor.LoadData to preload everything
func (a *RelicMainAffixAvatarValueAccessor) GroupData() {
	a._dataAvatarID = map[float64]RelicMainAffixAvatarValue{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
	}
}

// ByAvatarID returns the RelicMainAffixAvatarValue uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *RelicMainAffixAvatarValueAccessor) ByAvatarID(identifier float64) (RelicMainAffixAvatarValue, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicMainAffixAvatarValue{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}
