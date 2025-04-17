package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ItemUseData struct {
	IsAutoUse      bool      `json:"IsAutoUse"`
	UseDataID      float64   `json:"UseDataID"`
	UseMultipleMax float64   `json:"UseMultipleMax"`
	UseParam       []float64 `json:"UseParam"`
}
type ItemUseDataAccessor struct {
	_data          []ItemUseData
	_dataUseDataID map[float64]ItemUseData
}

// LoadData retrieves the data. Must be called before ItemUseData.GroupData
func (a *ItemUseDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ItemUseData.json")
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
func (a *ItemUseDataAccessor) Raw() ([]ItemUseData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ItemUseData{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ItemUseDataAccessor.LoadData to preload everything
func (a *ItemUseDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataUseDataID[d.UseDataID] = d
	}
}

// ByUseDataID returns the ItemUseData uniquely identified by UseDataID
//
// Error is only non-nil if the source errors out
func (a *ItemUseDataAccessor) ByUseDataID(identifier float64) (ItemUseData, error) {
	if a._dataUseDataID == nil {
		err := a.LoadData()
		if err != nil {
			return ItemUseData{}, err
		}
		a.GroupData()
	}
	return a._dataUseDataID[identifier], nil
}
