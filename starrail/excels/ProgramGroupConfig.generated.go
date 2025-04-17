package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ProgramGroupConfig struct {
	Asset          string  `json:"Asset"`
	Duration       float64 `json:"Duration"`
	ID             float64 `json:"ID"`
	IfAnAsset      string  `json:"IfAnAsset"`
	Order          float64 `json:"Order"`
	PlayType       float64 `json:"PlayType"`
	ProgramGroupID float64 `json:"ProgramGroupID"`
	SoundEvent     string  `json:"SoundEvent"`
}
type ProgramGroupConfigAccessor struct {
	_data   []ProgramGroupConfig
	_dataID map[float64]ProgramGroupConfig
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
	a._dataID = map[float64]ProgramGroupConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ProgramGroupConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ProgramGroupConfigAccessor) ByID(identifier float64) (ProgramGroupConfig, error) {
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
