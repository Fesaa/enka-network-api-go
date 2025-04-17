package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RndOptionGroup struct {
	ID          string      `json:"ID"`
	OptionCount json.Number `json:"OptionCount"`
}
type RndOptionGroupAccessor struct {
	_data            []RndOptionGroup
	_dataID          map[string]RndOptionGroup
	_dataOptionCount map[json.Number]RndOptionGroup
}

// LoadData retrieves the data. Must be called before RndOptionGroup.GroupData
func (a *RndOptionGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RndOptionGroup.json")
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
func (a *RndOptionGroupAccessor) Raw() ([]RndOptionGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RndOptionGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RndOptionGroupAccessor.LoadData to preload everything
func (a *RndOptionGroupAccessor) GroupData() {
	a._dataID = map[string]RndOptionGroup{}
	a._dataOptionCount = map[json.Number]RndOptionGroup{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataOptionCount[d.OptionCount] = d
	}
}

// ByID returns the RndOptionGroup uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *RndOptionGroupAccessor) ByID(identifier string) (RndOptionGroup, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RndOptionGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByOptionCount returns the RndOptionGroup uniquely identified by OptionCount
//
// Error is only non-nil if the source errors out
func (a *RndOptionGroupAccessor) ByOptionCount(identifier json.Number) (RndOptionGroup, error) {
	if a._dataOptionCount == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RndOptionGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOptionCount[identifier], nil
}
