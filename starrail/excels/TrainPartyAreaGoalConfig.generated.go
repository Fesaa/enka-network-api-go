package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyAreaGoalConfig struct {
	AreaID        float64   `json:"AreaID"`
	ID            float64   `json:"ID"`
	StepGroupList []float64 `json:"StepGroupList"`
}
type TrainPartyAreaGoalConfigAccessor struct {
	_data       []TrainPartyAreaGoalConfig
	_dataAreaID map[float64]TrainPartyAreaGoalConfig
	_dataID     map[float64]TrainPartyAreaGoalConfig
}

// LoadData retrieves the data. Must be called before TrainPartyAreaGoalConfig.GroupData
func (a *TrainPartyAreaGoalConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyAreaGoalConfig.json")
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
func (a *TrainPartyAreaGoalConfigAccessor) Raw() ([]TrainPartyAreaGoalConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyAreaGoalConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyAreaGoalConfigAccessor.LoadData to preload everything
func (a *TrainPartyAreaGoalConfigAccessor) GroupData() {
	a._dataAreaID = map[float64]TrainPartyAreaGoalConfig{}
	a._dataID = map[float64]TrainPartyAreaGoalConfig{}
	for _, d := range a._data {
		a._dataAreaID[d.AreaID] = d
		a._dataID[d.ID] = d
	}
}

// ByAreaID returns the TrainPartyAreaGoalConfig uniquely identified by AreaID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyAreaGoalConfigAccessor) ByAreaID(identifier float64) (TrainPartyAreaGoalConfig, error) {
	if a._dataAreaID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyAreaGoalConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAreaID[identifier], nil
}

// ByID returns the TrainPartyAreaGoalConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyAreaGoalConfigAccessor) ByID(identifier float64) (TrainPartyAreaGoalConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyAreaGoalConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
