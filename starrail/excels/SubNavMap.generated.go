package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SubNavMap struct {
	FloorID        float64 `json:"FloorID"`
	ID             float64 `json:"ID"`
	NavMapSubTabID float64 `json:"NavMapSubTabID"`
	Type           string  `json:"Type"`
}
type SubNavMapAccessor struct {
	_data        []SubNavMap
	_dataFloorID map[float64]SubNavMap
	_dataID      map[float64]SubNavMap
}

// LoadData retrieves the data. Must be called before SubNavMap.GroupData
func (a *SubNavMapAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SubNavMap.json")
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
func (a *SubNavMapAccessor) Raw() ([]SubNavMap, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SubNavMap{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SubNavMapAccessor.LoadData to preload everything
func (a *SubNavMapAccessor) GroupData() {
	for _, d := range a._data {
		a._dataFloorID[d.FloorID] = d
		a._dataID[d.ID] = d
	}
}

// ByFloorID returns the SubNavMap uniquely identified by FloorID
//
// Error is only non-nil if the source errors out
func (a *SubNavMapAccessor) ByFloorID(identifier float64) (SubNavMap, error) {
	if a._dataFloorID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SubNavMap{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFloorID[identifier], nil
}

// ByID returns the SubNavMap uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *SubNavMapAccessor) ByID(identifier float64) (SubNavMap, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SubNavMap{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
