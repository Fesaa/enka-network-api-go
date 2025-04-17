package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PerformanceD struct {
	ChangePlayerType     string      `json:"ChangePlayerType"`
	EndBlack             string      `json:"EndBlack"`
	EndWithCrack         bool        `json:"EndWithCrack"`
	FloorID              json.Number `json:"FloorID"`
	GroupID              json.Number `json:"GroupID"`
	IsSkip               string      `json:"IsSkip"`
	PerformanceCharacter string      `json:"PerformanceCharacter"`
	PerformanceID        json.Number `json:"PerformanceID"`
	PerformancePath      string      `json:"PerformancePath"`
	PlaneID              json.Number `json:"PlaneID"`
	StartBlack           string      `json:"StartBlack"`
}
type PerformanceDAccessor struct {
	_data                []PerformanceD
	_dataPerformancePath map[string]PerformanceD
}

// LoadData retrieves the data. Must be called before PerformanceD.GroupData
func (a *PerformanceDAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PerformanceD.json")
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
func (a *PerformanceDAccessor) Raw() ([]PerformanceD, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PerformanceD{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PerformanceDAccessor.LoadData to preload everything
func (a *PerformanceDAccessor) GroupData() {
	a._dataPerformancePath = map[string]PerformanceD{}
	for _, d := range a._data {
		a._dataPerformancePath[d.PerformancePath] = d
	}
}

// ByPerformancePath returns the PerformanceD uniquely identified by PerformancePath
//
// Error is only non-nil if the source errors out
func (a *PerformanceDAccessor) ByPerformancePath(identifier string) (PerformanceD, error) {
	if a._dataPerformancePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PerformanceD{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPerformancePath[identifier], nil
}
