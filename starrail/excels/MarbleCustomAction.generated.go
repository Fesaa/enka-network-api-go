package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MarbleCustomAction struct {
	ID              float64   `json:"ID"`
	LaunchParamList []float64 `json:"LaunchParamList"`
	SealInsID       float64   `json:"SealInsID"`
}
type MarbleCustomActionAccessor struct {
	_data          []MarbleCustomAction
	_dataID        map[float64]MarbleCustomAction
	_dataSealInsID map[float64]MarbleCustomAction
}

// LoadData retrieves the data. Must be called before MarbleCustomAction.GroupData
func (a *MarbleCustomActionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleCustomAction.json")
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
func (a *MarbleCustomActionAccessor) Raw() ([]MarbleCustomAction, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleCustomAction{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MarbleCustomActionAccessor.LoadData to preload everything
func (a *MarbleCustomActionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataSealInsID[d.SealInsID] = d
	}
}

// ByID returns the MarbleCustomAction uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MarbleCustomActionAccessor) ByID(identifier float64) (MarbleCustomAction, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return MarbleCustomAction{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// BySealInsID returns the MarbleCustomAction uniquely identified by SealInsID
//
// Error is only non-nil if the source errors out
func (a *MarbleCustomActionAccessor) BySealInsID(identifier float64) (MarbleCustomAction, error) {
	if a._dataSealInsID == nil {
		err := a.LoadData()
		if err != nil {
			return MarbleCustomAction{}, err
		}
		a.GroupData()
	}
	return a._dataSealInsID[identifier], nil
}
