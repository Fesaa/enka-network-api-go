package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RndOptionsData struct {
	DialogShowOrder float64 `json:"DialogShowOrder"`
	GroupID         string  `json:"GroupID"`
	ID              float64 `json:"ID"`
	JsonPath        string  `json:"JsonPath"`
	MenuItemID      float64 `json:"MenuItemID"`
	MenuItemType    string  `json:"MenuItemType"`
	Weight          float64 `json:"Weight"`
}
type RndOptionsDataAccessor struct {
	_data           []RndOptionsData
	_dataID         map[float64]RndOptionsData
	_dataMenuItemID map[float64]RndOptionsData
	_dataJsonPath   map[string]RndOptionsData
}

// LoadData retrieves the data. Must be called before RndOptionsData.GroupData
func (a *RndOptionsDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RndOptionsData.json")
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
func (a *RndOptionsDataAccessor) Raw() ([]RndOptionsData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RndOptionsData{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RndOptionsDataAccessor.LoadData to preload everything
func (a *RndOptionsDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataMenuItemID[d.MenuItemID] = d
		a._dataJsonPath[d.JsonPath] = d
	}
}

// ByID returns the RndOptionsData uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *RndOptionsDataAccessor) ByID(identifier float64) (RndOptionsData, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return RndOptionsData{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByMenuItemID returns the RndOptionsData uniquely identified by MenuItemID
//
// Error is only non-nil if the source errors out
func (a *RndOptionsDataAccessor) ByMenuItemID(identifier float64) (RndOptionsData, error) {
	if a._dataMenuItemID == nil {
		err := a.LoadData()
		if err != nil {
			return RndOptionsData{}, err
		}
		a.GroupData()
	}
	return a._dataMenuItemID[identifier], nil
}

// ByJsonPath returns the RndOptionsData uniquely identified by JsonPath
//
// Error is only non-nil if the source errors out
func (a *RndOptionsDataAccessor) ByJsonPath(identifier string) (RndOptionsData, error) {
	if a._dataJsonPath == nil {
		err := a.LoadData()
		if err != nil {
			return RndOptionsData{}, err
		}
		a.GroupData()
	}
	return a._dataJsonPath[identifier], nil
}
