package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PerformanceA struct {
	EndBlack        string      `json:"EndBlack"`
	EndWithCrack    bool        `json:"EndWithCrack"`
	FloorID         json.Number `json:"FloorID"`
	IsSkip          string      `json:"IsSkip"`
	PerformanceID   json.Number `json:"PerformanceID"`
	PerformancePath string      `json:"PerformancePath"`
	PlaneID         json.Number `json:"PlaneID"`
	StartBlack      string      `json:"StartBlack"`
}
type PerformanceAAccessor struct {
	_data []PerformanceA
}

// LoadData retrieves the data. Must be called before PerformanceA.GroupData
func (a *PerformanceAAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PerformanceA.json")
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
func (a *PerformanceAAccessor) Raw() ([]PerformanceA, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PerformanceA{}, err
		}
	}
	return a._data, nil
}
