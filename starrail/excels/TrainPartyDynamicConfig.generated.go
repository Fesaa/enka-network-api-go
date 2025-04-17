package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyDynamicConfig struct {
	ID          float64   `json:"ID"`
	PrefabPath  string    `json:"PrefabPath"`
	Taglist     []float64 `json:"Taglist"`
	UseLowLight bool      `json:"UseLowLight"`
}
type TrainPartyDynamicConfigAccessor struct {
	_data           []TrainPartyDynamicConfig
	_dataID         map[float64]TrainPartyDynamicConfig
	_dataPrefabPath map[string]TrainPartyDynamicConfig
}

// LoadData retrieves the data. Must be called before TrainPartyDynamicConfig.GroupData
func (a *TrainPartyDynamicConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyDynamicConfig.json")
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
func (a *TrainPartyDynamicConfigAccessor) Raw() ([]TrainPartyDynamicConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyDynamicConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyDynamicConfigAccessor.LoadData to preload everything
func (a *TrainPartyDynamicConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataPrefabPath[d.PrefabPath] = d
	}
}

// ByID returns the TrainPartyDynamicConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyDynamicConfigAccessor) ByID(identifier float64) (TrainPartyDynamicConfig, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return TrainPartyDynamicConfig{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByPrefabPath returns the TrainPartyDynamicConfig uniquely identified by PrefabPath
//
// Error is only non-nil if the source errors out
func (a *TrainPartyDynamicConfigAccessor) ByPrefabPath(identifier string) (TrainPartyDynamicConfig, error) {
	if a._dataPrefabPath == nil {
		err := a.LoadData()
		if err != nil {
			return TrainPartyDynamicConfig{}, err
		}
		a.GroupData()
	}
	return a._dataPrefabPath[identifier], nil
}
