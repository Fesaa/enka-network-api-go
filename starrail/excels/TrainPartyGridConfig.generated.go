package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TrainPartyGridConfig struct {
	GridID       json.Number   `json:"GridID"`
	GridIconPath string        `json:"GridIconPath"`
	GridType     string        `json:"GridType"`
	ParamList    []json.Number `json:"ParamList"`
}
type TrainPartyGridConfigAccessor struct {
	_data       []TrainPartyGridConfig
	_dataGridID map[json.Number]TrainPartyGridConfig
}

// LoadData retrieves the data. Must be called before TrainPartyGridConfig.GroupData
func (a *TrainPartyGridConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyGridConfig.json")
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
func (a *TrainPartyGridConfigAccessor) Raw() ([]TrainPartyGridConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyGridConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyGridConfigAccessor.LoadData to preload everything
func (a *TrainPartyGridConfigAccessor) GroupData() {
	a._dataGridID = map[json.Number]TrainPartyGridConfig{}
	for _, d := range a._data {
		a._dataGridID[d.GridID] = d
	}
}

// ByGridID returns the TrainPartyGridConfig uniquely identified by GridID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyGridConfigAccessor) ByGridID(identifier json.Number) (TrainPartyGridConfig, error) {
	if a._dataGridID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyGridConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGridID[identifier], nil
}
