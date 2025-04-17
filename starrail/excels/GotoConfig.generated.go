package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type GotoConfig struct {
	GotoType          json.Number   `json:"GotoType"`
	ID                json.Number   `json:"ID"`
	ParamIntList      []json.Number `json:"ParamIntList"`
	ParamStringList   []string      `json:"ParamStringList"`
	UnlockID          json.Number   `json:"UnlockID"`
	UnlockMainMission json.Number   `json:"UnlockMainMission"`
}
type GotoConfigAccessor struct {
	_data   []GotoConfig
	_dataID map[json.Number]GotoConfig
}

// LoadData retrieves the data. Must be called before GotoConfig.GroupData
func (a *GotoConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GotoConfig.json")
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
func (a *GotoConfigAccessor) Raw() ([]GotoConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GotoConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GotoConfigAccessor.LoadData to preload everything
func (a *GotoConfigAccessor) GroupData() {
	a._dataID = map[json.Number]GotoConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the GotoConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *GotoConfigAccessor) ByID(identifier json.Number) (GotoConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GotoConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
