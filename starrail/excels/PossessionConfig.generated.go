package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PossessionConfig struct {
	AttachPoint          string    `json:"AttachPoint"`
	IsEffect             bool      `json:"IsEffect"`
	LocalPosition        []float64 `json:"LocalPosition"`
	LocalRotation        []float64 `json:"LocalRotation"`
	LocalScale           []float64 `json:"LocalScale"`
	PossessionName       string    `json:"PossessionName"`
	PossessionPrefabPath string    `json:"PossessionPrefabPath"`
}
type PossessionConfigAccessor struct {
	_data               []PossessionConfig
	_dataPossessionName map[string]PossessionConfig
}

// LoadData retrieves the data. Must be called before PossessionConfig.GroupData
func (a *PossessionConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PossessionConfig.json")
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
func (a *PossessionConfigAccessor) Raw() ([]PossessionConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PossessionConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PossessionConfigAccessor.LoadData to preload everything
func (a *PossessionConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPossessionName[d.PossessionName] = d
	}
}

// ByPossessionName returns the PossessionConfig uniquely identified by PossessionName
//
// Error is only non-nil if the source errors out
func (a *PossessionConfigAccessor) ByPossessionName(identifier string) (PossessionConfig, error) {
	if a._dataPossessionName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PossessionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPossessionName[identifier], nil
}
