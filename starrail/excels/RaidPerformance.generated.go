package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RaidPerformance struct {
	PerformanceID   float64 `json:"PerformanceID"`
	PerformanceType string  `json:"PerformanceType"`
	RaidID          float64 `json:"RaidID"`
}
type RaidPerformanceAccessor struct {
	_data              []RaidPerformance
	_dataPerformanceID map[float64]RaidPerformance
	_dataRaidID        map[float64]RaidPerformance
}

// LoadData retrieves the data. Must be called before RaidPerformance.GroupData
func (a *RaidPerformanceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RaidPerformance.json")
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
func (a *RaidPerformanceAccessor) Raw() ([]RaidPerformance, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RaidPerformance{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RaidPerformanceAccessor.LoadData to preload everything
func (a *RaidPerformanceAccessor) GroupData() {
	a._dataPerformanceID = map[float64]RaidPerformance{}
	a._dataRaidID = map[float64]RaidPerformance{}
	for _, d := range a._data {
		a._dataPerformanceID[d.PerformanceID] = d
		a._dataRaidID[d.RaidID] = d
	}
}

// ByPerformanceID returns the RaidPerformance uniquely identified by PerformanceID
//
// Error is only non-nil if the source errors out
func (a *RaidPerformanceAccessor) ByPerformanceID(identifier float64) (RaidPerformance, error) {
	if a._dataPerformanceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RaidPerformance{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPerformanceID[identifier], nil
}

// ByRaidID returns the RaidPerformance uniquely identified by RaidID
//
// Error is only non-nil if the source errors out
func (a *RaidPerformanceAccessor) ByRaidID(identifier float64) (RaidPerformance, error) {
	if a._dataRaidID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RaidPerformance{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidID[identifier], nil
}
