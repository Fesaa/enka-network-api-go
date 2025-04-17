package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueTalent struct {
	Cost                []map[string]json.Number `json:"Cost"`
	EffectDesc          map[string]json.Number   `json:"EffectDesc"`
	EffectDescParamList []map[string]json.Number `json:"EffectDescParamList"`
	EffectTag           map[string]json.Number   `json:"EffectTag"`
	EffectTitle         map[string]json.Number   `json:"EffectTitle"`
	Icon                string                   `json:"Icon"`
	IsImportant         bool                     `json:"IsImportant"`
	NextTalentIDList    []json.Number            `json:"NextTalentIDList"`
	TalentID            json.Number              `json:"TalentID"`
	UnlockIDList        []json.Number            `json:"UnlockIDList"`
}
type RogueTalentAccessor struct {
	_data         []RogueTalent
	_dataTalentID map[json.Number]RogueTalent
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
	a._dataTalentID = map[json.Number]RogueTalent{}
	for _, d := range a._data {
		a._dataTalentID[d.TalentID] = d
	}
}

// ByTalentID returns the RogueTalent uniquely identified by TalentID
//
// Error is only non-nil if the source errors out
func (a *RogueTalentAccessor) ByTalentID(identifier json.Number) (RogueTalent, error) {
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
