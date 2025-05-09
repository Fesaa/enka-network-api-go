package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartySkillEffect struct {
	EffectID   float64   `json:"EffectID"`
	EffectType string    `json:"EffectType"`
	ParamList  []float64 `json:"ParamList"`
}
type TrainPartySkillEffectAccessor struct {
	_data         []TrainPartySkillEffect
	_dataEffectID map[float64]TrainPartySkillEffect
}

// LoadData retrieves the data. Must be called before TrainPartySkillEffect.GroupData
func (a *TrainPartySkillEffectAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartySkillEffect.json")
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
func (a *TrainPartySkillEffectAccessor) Raw() ([]TrainPartySkillEffect, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartySkillEffect{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartySkillEffectAccessor.LoadData to preload everything
func (a *TrainPartySkillEffectAccessor) GroupData() {
	a._dataEffectID = map[float64]TrainPartySkillEffect{}
	for _, d := range a._data {
		a._dataEffectID[d.EffectID] = d
	}
}

// ByEffectID returns the TrainPartySkillEffect uniquely identified by EffectID
//
// Error is only non-nil if the source errors out
func (a *TrainPartySkillEffectAccessor) ByEffectID(identifier float64) (TrainPartySkillEffect, error) {
	if a._dataEffectID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartySkillEffect{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEffectID[identifier], nil
}
