package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TrainPartySkillConfig struct {
	IsRare           bool                   `json:"IsRare"`
	SKillID          json.Number            `json:"SKillID"`
	SkillDescription map[string]json.Number `json:"SkillDescription"`
	SkillFigurePath  string                 `json:"SkillFigurePath"`
	SkillIconPath    string                 `json:"SkillIconPath"`
	SkillName        map[string]json.Number `json:"SkillName"`
}
type TrainPartySkillConfigAccessor struct {
	_data        []TrainPartySkillConfig
	_dataSKillID map[json.Number]TrainPartySkillConfig
}

// LoadData retrieves the data. Must be called before TrainPartySkillConfig.GroupData
func (a *TrainPartySkillConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartySkillConfig.json")
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
func (a *TrainPartySkillConfigAccessor) Raw() ([]TrainPartySkillConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartySkillConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartySkillConfigAccessor.LoadData to preload everything
func (a *TrainPartySkillConfigAccessor) GroupData() {
	a._dataSKillID = map[json.Number]TrainPartySkillConfig{}
	for _, d := range a._data {
		a._dataSKillID[d.SKillID] = d
	}
}

// BySKillID returns the TrainPartySkillConfig uniquely identified by SKillID
//
// Error is only non-nil if the source errors out
func (a *TrainPartySkillConfigAccessor) BySKillID(identifier json.Number) (TrainPartySkillConfig, error) {
	if a._dataSKillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartySkillConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSKillID[identifier], nil
}
