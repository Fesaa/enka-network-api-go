package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MultiFloorConflictGroup struct {
	FloorIDList []float64 `json:"FloorIDList"`
	GroupID     float64   `json:"GroupID"`
	PlaneID     float64   `json:"PlaneID"`
}
type MultiFloorConflictGroupAccessor struct {
	_data        []MultiFloorConflictGroup
	_dataGroupID map[float64]MultiFloorConflictGroup
	_dataPlaneID map[float64]MultiFloorConflictGroup
}

// LoadData retrieves the data. Must be called before MultiFloorConflictGroup.GroupData
func (a *MultiFloorConflictGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MultiFloorConflictGroup.json")
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
func (a *MultiFloorConflictGroupAccessor) Raw() ([]MultiFloorConflictGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MultiFloorConflictGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MultiFloorConflictGroupAccessor.LoadData to preload everything
func (a *MultiFloorConflictGroupAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGroupID[d.GroupID] = d
		a._dataPlaneID[d.PlaneID] = d
	}
}

// ByGroupID returns the MultiFloorConflictGroup uniquely identified by GroupID
//
// Error is only non-nil if the source errors out
func (a *MultiFloorConflictGroupAccessor) ByGroupID(identifier float64) (MultiFloorConflictGroup, error) {
	if a._dataGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MultiFloorConflictGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupID[identifier], nil
}

// ByPlaneID returns the MultiFloorConflictGroup uniquely identified by PlaneID
//
// Error is only non-nil if the source errors out
func (a *MultiFloorConflictGroupAccessor) ByPlaneID(identifier float64) (MultiFloorConflictGroup, error) {
	if a._dataPlaneID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MultiFloorConflictGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPlaneID[identifier], nil
}
