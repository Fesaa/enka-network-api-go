package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournCollectionConfig struct {
	ConfigID float64 `json:"ConfigID"`
	GroupID  float64 `json:"GroupID"`
	PillarID float64 `json:"PillarID"`
}
type RogueTournCollectionConfigAccessor struct {
	_data         []RogueTournCollectionConfig
	_dataConfigID map[float64]RogueTournCollectionConfig
	_dataPillarID map[float64]RogueTournCollectionConfig
}

// LoadData retrieves the data. Must be called before RogueTournCollectionConfig.GroupData
func (a *RogueTournCollectionConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournCollectionConfig.json")
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
func (a *RogueTournCollectionConfigAccessor) Raw() ([]RogueTournCollectionConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournCollectionConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournCollectionConfigAccessor.LoadData to preload everything
func (a *RogueTournCollectionConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataConfigID[d.ConfigID] = d
		a._dataPillarID[d.PillarID] = d
	}
}

// ByConfigID returns the RogueTournCollectionConfig uniquely identified by ConfigID
//
// Error is only non-nil if the source errors out
func (a *RogueTournCollectionConfigAccessor) ByConfigID(identifier float64) (RogueTournCollectionConfig, error) {
	if a._dataConfigID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournCollectionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataConfigID[identifier], nil
}

// ByPillarID returns the RogueTournCollectionConfig uniquely identified by PillarID
//
// Error is only non-nil if the source errors out
func (a *RogueTournCollectionConfigAccessor) ByPillarID(identifier float64) (RogueTournCollectionConfig, error) {
	if a._dataPillarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournCollectionConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPillarID[identifier], nil
}
