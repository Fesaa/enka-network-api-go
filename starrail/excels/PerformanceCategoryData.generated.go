package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PerformanceCategoryData struct {
	Category      map[string]json.Number `json:"Category"`
	CategoryID    json.Number            `json:"CategoryID"`
	IconPath      string                 `json:"IconPath"`
	IsSubCategory bool                   `json:"isSubCategory"`
}
type PerformanceCategoryDataAccessor struct {
	_data []PerformanceCategoryData
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
