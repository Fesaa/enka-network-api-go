package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type LoadingStratageConfig struct {
	AvailableEntranceIDList []json.Number                         `json:"AvailableEntranceIDList"`
	FloorOperation          []LoadingStratageConfigFloorOperation `json:"FloorOperation"`
	LoadingID               json.Number                           `json:"LoadingID"`
	MissionIDList           []json.Number                         `json:"MissionIDList"`
	Priority                json.Number                           `json:"Priority"`
	StratageType            string                                `json:"StratageType"`
}
type LoadingStratageConfigFloorOperation struct {
	CAOAPDCCPCA string `json:"CAOAPDCCPCA"`
	PICHIHHCOCB string `json:"PICHIHHCOCB"`
}
type LoadingStratageConfigAccessor struct {
	_data          []LoadingStratageConfig
	_dataLoadingID map[json.Number]LoadingStratageConfig
}

// LoadData retrieves the data. Must be called before LoadingStratageConfig.GroupData
func (a *LoadingStratageConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/LoadingStratageConfig.json")
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
func (a *LoadingStratageConfigAccessor) Raw() ([]LoadingStratageConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []LoadingStratageConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with LoadingStratageConfigAccessor.LoadData to preload everything
func (a *LoadingStratageConfigAccessor) GroupData() {
	a._dataLoadingID = map[json.Number]LoadingStratageConfig{}
	for _, d := range a._data {
		a._dataLoadingID[d.LoadingID] = d
	}
}

// ByLoadingID returns the LoadingStratageConfig uniquely identified by LoadingID
//
// Error is only non-nil if the source errors out
func (a *LoadingStratageConfigAccessor) ByLoadingID(identifier json.Number) (LoadingStratageConfig, error) {
	if a._dataLoadingID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return LoadingStratageConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLoadingID[identifier], nil
}
