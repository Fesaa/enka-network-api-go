package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueCaptureMonster struct {
	GameTime     json.Number   `json:"GameTime"`
	MonsterNum   json.Number   `json:"MonsterNum"`
	ParamGroupID json.Number   `json:"ParamGroupID"`
	PrepareTime  json.Number   `json:"PrepareTime"`
	ScoreRange   []json.Number `json:"ScoreRange"`
}
type RogueCaptureMonsterAccessor struct {
	_data             []RogueCaptureMonster
	_dataParamGroupID map[json.Number]RogueCaptureMonster
}

// LoadData retrieves the data. Must be called before RogueCaptureMonster.GroupData
func (a *RogueCaptureMonsterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueCaptureMonster.json")
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
func (a *RogueCaptureMonsterAccessor) Raw() ([]RogueCaptureMonster, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueCaptureMonster{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueCaptureMonsterAccessor.LoadData to preload everything
func (a *RogueCaptureMonsterAccessor) GroupData() {
	a._dataParamGroupID = map[json.Number]RogueCaptureMonster{}
	for _, d := range a._data {
		a._dataParamGroupID[d.ParamGroupID] = d
	}
}

// ByParamGroupID returns the RogueCaptureMonster uniquely identified by ParamGroupID
//
// Error is only non-nil if the source errors out
func (a *RogueCaptureMonsterAccessor) ByParamGroupID(identifier json.Number) (RogueCaptureMonster, error) {
	if a._dataParamGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueCaptureMonster{}, err
			}
		}
		a.GroupData()
	}
	return a._dataParamGroupID[identifier], nil
}
