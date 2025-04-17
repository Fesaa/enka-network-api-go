package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type BattleConditionConfig struct {
	AbilityName    string                 `json:"AbilityName"`
	ConditionDes   map[string]json.Number `json:"ConditionDes"`
	ID             json.Number            `json:"ID"`
	IsShowProgress json.Number            `json:"IsShowProgress"`
	TargetParam    json.Number            `json:"TargetParam"`
	WinOrLose      bool                   `json:"WinOrLose"`
}
type BattleConditionConfigAccessor struct {
	_data   []BattleConditionConfig
	_dataID map[json.Number]BattleConditionConfig
}

// LoadData retrieves the data. Must be called before BattleConditionConfig.GroupData
func (a *BattleConditionConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleConditionConfig.json")
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
func (a *BattleConditionConfigAccessor) Raw() ([]BattleConditionConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleConditionConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleConditionConfigAccessor.LoadData to preload everything
func (a *BattleConditionConfigAccessor) GroupData() {
	a._dataID = map[json.Number]BattleConditionConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the BattleConditionConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BattleConditionConfigAccessor) ByID(identifier json.Number) (BattleConditionConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleConditionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
