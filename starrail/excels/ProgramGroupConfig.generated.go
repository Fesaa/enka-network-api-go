package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ProgramGroupConfig struct {
	Asset          string      `json:"Asset"`
	Duration       json.Number `json:"Duration"`
	ID             json.Number `json:"ID"`
	IfAnAsset      string      `json:"IfAnAsset"`
	Order          json.Number `json:"Order"`
	PlayType       json.Number `json:"PlayType"`
	ProgramGroupID json.Number `json:"ProgramGroupID"`
	SoundEvent     string      `json:"SoundEvent"`
}
type ProgramGroupConfigAccessor struct {
	_data   []ProgramGroupConfig
	_dataID map[json.Number]ProgramGroupConfig
}

// LoadData retrieves the data. Must be called before ProgramGroupConfig.GroupData
func (a *ProgramGroupConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ProgramGroupConfig.json")
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
func (a *ProgramGroupConfigAccessor) Raw() ([]ProgramGroupConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ProgramGroupConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ProgramGroupConfigAccessor.LoadData to preload everything
func (a *ProgramGroupConfigAccessor) GroupData() {
	a._dataID = map[json.Number]ProgramGroupConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ProgramGroupConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ProgramGroupConfigAccessor) ByID(identifier json.Number) (ProgramGroupConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ProgramGroupConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
