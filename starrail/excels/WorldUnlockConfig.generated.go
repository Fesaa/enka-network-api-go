package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type WorldUnlockConfig struct {
	DirectUnlockCondition string    `json:"DirectUnlockCondition"`
	ID                    float64   `json:"ID"`
	InitMainMissionList   []float64 `json:"InitMainMissionList"`
	MainMissionIDList     []float64 `json:"MainMissionIDList"`
}
type WorldUnlockConfigAccessor struct {
	_data                      []WorldUnlockConfig
	_dataDirectUnlockCondition map[string]WorldUnlockConfig
	_dataID                    map[float64]WorldUnlockConfig
}

// LoadData retrieves the data. Must be called before WorldUnlockConfig.GroupData
func (a *WorldUnlockConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/WorldUnlockConfig.json")
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
func (a *WorldUnlockConfigAccessor) Raw() ([]WorldUnlockConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []WorldUnlockConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with WorldUnlockConfigAccessor.LoadData to preload everything
func (a *WorldUnlockConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataDirectUnlockCondition[d.DirectUnlockCondition] = d
		a._dataID[d.ID] = d
	}
}

// ByDirectUnlockCondition returns the WorldUnlockConfig uniquely identified by DirectUnlockCondition
//
// Error is only non-nil if the source errors out
func (a *WorldUnlockConfigAccessor) ByDirectUnlockCondition(identifier string) (WorldUnlockConfig, error) {
	if a._dataDirectUnlockCondition == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return WorldUnlockConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDirectUnlockCondition[identifier], nil
}

// ByID returns the WorldUnlockConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *WorldUnlockConfigAccessor) ByID(identifier float64) (WorldUnlockConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return WorldUnlockConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
