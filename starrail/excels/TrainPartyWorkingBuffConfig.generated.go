package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyWorkingBuffConfig struct {
	Description   hash.Hash `json:"Description"`
	IconPath      string    `json:"IconPath"`
	Name          hash.Hash `json:"Name"`
	WorkingBuffID float64   `json:"WorkingBuffID"`
}
type TrainPartyWorkingBuffConfigAccessor struct {
	_data              []TrainPartyWorkingBuffConfig
	_dataIconPath      map[string]TrainPartyWorkingBuffConfig
	_dataWorkingBuffID map[float64]TrainPartyWorkingBuffConfig
}

// LoadData retrieves the data. Must be called before TrainPartyWorkingBuffConfig.GroupData
func (a *TrainPartyWorkingBuffConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyWorkingBuffConfig.json")
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
func (a *TrainPartyWorkingBuffConfigAccessor) Raw() ([]TrainPartyWorkingBuffConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyWorkingBuffConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyWorkingBuffConfigAccessor.LoadData to preload everything
func (a *TrainPartyWorkingBuffConfigAccessor) GroupData() {
	a._dataIconPath = map[string]TrainPartyWorkingBuffConfig{}
	a._dataWorkingBuffID = map[float64]TrainPartyWorkingBuffConfig{}
	for _, d := range a._data {
		a._dataIconPath[d.IconPath] = d
		a._dataWorkingBuffID[d.WorkingBuffID] = d
	}
}

// ByIconPath returns the TrainPartyWorkingBuffConfig uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *TrainPartyWorkingBuffConfigAccessor) ByIconPath(identifier string) (TrainPartyWorkingBuffConfig, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyWorkingBuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByWorkingBuffID returns the TrainPartyWorkingBuffConfig uniquely identified by WorkingBuffID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyWorkingBuffConfigAccessor) ByWorkingBuffID(identifier float64) (TrainPartyWorkingBuffConfig, error) {
	if a._dataWorkingBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyWorkingBuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWorkingBuffID[identifier], nil
}
