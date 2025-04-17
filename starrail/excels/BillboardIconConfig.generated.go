package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type BillboardIconConfig struct {
	BillboardIconPath string      `json:"BillboardIconPath"`
	ID                json.Number `json:"ID"`
	Priority          json.Number `json:"Priority"`
}
type BillboardIconConfigAccessor struct {
	_data   []BillboardIconConfig
	_dataID map[json.Number]BillboardIconConfig
}

// LoadData retrieves the data. Must be called before BillboardIconConfig.GroupData
func (a *BillboardIconConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BillboardIconConfig.json")
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
func (a *BillboardIconConfigAccessor) Raw() ([]BillboardIconConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BillboardIconConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BillboardIconConfigAccessor.LoadData to preload everything
func (a *BillboardIconConfigAccessor) GroupData() {
	a._dataID = map[json.Number]BillboardIconConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the BillboardIconConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BillboardIconConfigAccessor) ByID(identifier json.Number) (BillboardIconConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BillboardIconConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
