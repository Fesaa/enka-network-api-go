package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type TrainPartySkillEffect struct {
	EffectID   json.Number   `json:"EffectID"`
	EffectType string        `json:"EffectType"`
	ParamList  []json.Number `json:"ParamList"`
}
type TrainPartySkillEffectAccessor struct {
	_data []TrainPartySkillEffect
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
