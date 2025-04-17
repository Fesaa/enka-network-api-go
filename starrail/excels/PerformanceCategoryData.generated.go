package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type PerformanceCategoryData struct {
	Category      map[string]json.Number `json:"Category"`
	CategoryID    json.Number            `json:"CategoryID"`
	IconPath      string                 `json:"IconPath"`
	IsSubCategory bool                   `json:"isSubCategory"`
}
type PerformanceCategoryDataAccessor struct {
	_data           []PerformanceCategoryData
	_dataCategoryID map[json.Number]PerformanceCategoryData
}

// LoadData retrieves the data. Must be called before PerformanceCategoryData.GroupData
func (a *PerformanceCategoryDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PerformanceCategoryData.json")
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
func (a *PerformanceCategoryDataAccessor) Raw() ([]PerformanceCategoryData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PerformanceCategoryData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PerformanceCategoryDataAccessor.LoadData to preload everything
func (a *PerformanceCategoryDataAccessor) GroupData() {
	a._dataCategoryID = map[json.Number]PerformanceCategoryData{}
	for _, d := range a._data {
		a._dataCategoryID[d.CategoryID] = d
	}
}

// ByCategoryID returns the PerformanceCategoryData uniquely identified by CategoryID
//
// Error is only non-nil if the source errors out
func (a *PerformanceCategoryDataAccessor) ByCategoryID(identifier json.Number) (PerformanceCategoryData, error) {
	if a._dataCategoryID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PerformanceCategoryData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCategoryID[identifier], nil
}
