package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueTalent struct {
	Cost                []RogueTalentCost `json:"Cost"`
	EffectDesc          hash.Hash         `json:"EffectDesc"`
	EffectDescParamList []hash.Value      `json:"EffectDescParamList"`
	EffectTag           hash.Hash         `json:"EffectTag"`
	EffectTitle         hash.Hash         `json:"EffectTitle"`
	Icon                string            `json:"Icon"`
	IsImportant         bool              `json:"IsImportant"`
	NextTalentIDList    []float64         `json:"NextTalentIDList"`
	TalentID            float64           `json:"TalentID"`
	UnlockIDList        []float64         `json:"UnlockIDList"`
}
type RogueTalentCost struct {
	ItemID  float64 `json:"ItemID"`
	ItemNum float64 `json:"ItemNum"`
}
type RogueTalentAccessor struct {
	_data         []RogueTalent
	_dataTalentID map[float64]RogueTalent
}

// LoadData retrieves the data. Must be called before RogueTalent.GroupData
func (a *RogueTalentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTalent.json")
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
func (a *RogueTalentAccessor) Raw() ([]RogueTalent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTalent{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTalentAccessor.LoadData to preload everything
func (a *RogueTalentAccessor) GroupData() {
	a._dataTalentID = map[float64]RogueTalent{}
	for _, d := range a._data {
		a._dataTalentID[d.TalentID] = d
	}
}

// ByTalentID returns the RogueTalent uniquely identified by TalentID
//
// Error is only non-nil if the source errors out
func (a *RogueTalentAccessor) ByTalentID(identifier float64) (RogueTalent, error) {
	if a._dataTalentID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTalent{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTalentID[identifier], nil
}
