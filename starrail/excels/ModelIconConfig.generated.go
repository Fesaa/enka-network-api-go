package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ModelIconConfig struct {
	ID         float64 `json:"ID"`
	PrefabPath string  `json:"PrefabPath"`
}
type ModelIconConfigAccessor struct {
	_data           []ModelIconConfig
	_dataID         map[float64]ModelIconConfig
	_dataPrefabPath map[string]ModelIconConfig
}

// LoadData retrieves the data. Must be called before ModelIconConfig.GroupData
func (a *ModelIconConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ModelIconConfig.json")
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
func (a *ModelIconConfigAccessor) Raw() ([]ModelIconConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ModelIconConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ModelIconConfigAccessor.LoadData to preload everything
func (a *ModelIconConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataPrefabPath[d.PrefabPath] = d
	}
}

// ByID returns the ModelIconConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ModelIconConfigAccessor) ByID(identifier float64) (ModelIconConfig, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return ModelIconConfig{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByPrefabPath returns the ModelIconConfig uniquely identified by PrefabPath
//
// Error is only non-nil if the source errors out
func (a *ModelIconConfigAccessor) ByPrefabPath(identifier string) (ModelIconConfig, error) {
	if a._dataPrefabPath == nil {
		err := a.LoadData()
		if err != nil {
			return ModelIconConfig{}, err
		}
		a.GroupData()
	}
	return a._dataPrefabPath[identifier], nil
}
