package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type PerformanceC struct {
	EndBlack        string      `json:"EndBlack"`
	EndWithCrack    bool        `json:"EndWithCrack"`
	FloorID         json.Number `json:"FloorID"`
	IsSkip          string      `json:"IsSkip"`
	PerformanceID   json.Number `json:"PerformanceID"`
	PerformancePath string      `json:"PerformancePath"`
	PlaneID         json.Number `json:"PlaneID"`
	StartBlack      string      `json:"StartBlack"`
}
type PerformanceCAccessor struct {
	_data              []PerformanceC
	_dataPerformanceID map[json.Number]PerformanceC
}

// LoadData retrieves the data. Must be called before PerformanceC.GroupData
func (a *PerformanceCAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PerformanceC.json")
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
func (a *PerformanceCAccessor) Raw() ([]PerformanceC, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PerformanceC{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PerformanceCAccessor.LoadData to preload everything
func (a *PerformanceCAccessor) GroupData() {
	a._dataPerformanceID = map[json.Number]PerformanceC{}
	for _, d := range a._data {
		a._dataPerformanceID[d.PerformanceID] = d
	}
}

// ByPerformanceID returns the PerformanceC uniquely identified by PerformanceID
//
// Error is only non-nil if the source errors out
func (a *PerformanceCAccessor) ByPerformanceID(identifier json.Number) (PerformanceC, error) {
	if a._dataPerformanceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PerformanceC{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPerformanceID[identifier], nil
}
