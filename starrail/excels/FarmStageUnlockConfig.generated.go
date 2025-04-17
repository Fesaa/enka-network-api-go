package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type FarmStageUnlockConfig struct {
	FarmGachaIDList       []interface{} `json:"FarmGachaIDList"`
	FarmType              string        `json:"FarmType"`
	ID                    json.Number   `json:"ID"`
	UIEnterBattleArea     json.Number   `json:"UIEnterBattleArea"`
	UIEntranceBgPath      string        `json:"UIEntranceBgPath"`
	UIEnviromentConfig    string        `json:"UIEnviromentConfig"`
	UnlockWorldLevelEnd   json.Number   `json:"UnlockWorldLevelEnd"`
	UnlockWorldLevelStart json.Number   `json:"UnlockWorldLevelStart"`
}
type FarmStageUnlockConfigAccessor struct {
	_data                 []FarmStageUnlockConfig
	_dataUIEntranceBgPath map[string]FarmStageUnlockConfig
}

// LoadData retrieves the data. Must be called before FarmStageUnlockConfig.GroupData
func (a *FarmStageUnlockConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FarmStageUnlockConfig.json")
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
func (a *FarmStageUnlockConfigAccessor) Raw() ([]FarmStageUnlockConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FarmStageUnlockConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FarmStageUnlockConfigAccessor.LoadData to preload everything
func (a *FarmStageUnlockConfigAccessor) GroupData() {
	a._dataUIEntranceBgPath = map[string]FarmStageUnlockConfig{}
	for _, d := range a._data {
		a._dataUIEntranceBgPath[d.UIEntranceBgPath] = d
	}
}

// ByUIEntranceBgPath returns the FarmStageUnlockConfig uniquely identified by UIEntranceBgPath
//
// Error is only non-nil if the source errors out
func (a *FarmStageUnlockConfigAccessor) ByUIEntranceBgPath(identifier string) (FarmStageUnlockConfig, error) {
	if a._dataUIEntranceBgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FarmStageUnlockConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUIEntranceBgPath[identifier], nil
}
