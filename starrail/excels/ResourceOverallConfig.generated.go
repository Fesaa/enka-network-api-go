package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ResourceOverallConfig struct {
	CurrencyIDList []float64 `json:"CurrencyIDList"`
	IsDetailButton bool      `json:"IsDetailButton"`
	PageKey        string    `json:"PageKey"`
}
type ResourceOverallConfigAccessor struct {
	_data        []ResourceOverallConfig
	_dataPageKey map[string]ResourceOverallConfig
}

// LoadData retrieves the data. Must be called before ResourceOverallConfig.GroupData
func (a *ResourceOverallConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ResourceOverallConfig.json")
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
func (a *ResourceOverallConfigAccessor) Raw() ([]ResourceOverallConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ResourceOverallConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ResourceOverallConfigAccessor.LoadData to preload everything
func (a *ResourceOverallConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataPageKey[d.PageKey] = d
	}
}

// ByPageKey returns the ResourceOverallConfig uniquely identified by PageKey
//
// Error is only non-nil if the source errors out
func (a *ResourceOverallConfigAccessor) ByPageKey(identifier string) (ResourceOverallConfig, error) {
	if a._dataPageKey == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ResourceOverallConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPageKey[identifier], nil
}
