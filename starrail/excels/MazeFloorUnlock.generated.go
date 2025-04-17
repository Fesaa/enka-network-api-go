package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MazeFloorUnlock struct {
	FloorID                   json.Number `json:"FloorID"`
	UnlockConditionExpression string      `json:"UnlockConditionExpression"`
}
type MazeFloorUnlockAccessor struct {
	_data        []MazeFloorUnlock
	_dataFloorID map[json.Number]MazeFloorUnlock
}

// LoadData retrieves the data. Must be called before MazeFloorUnlock.GroupData
func (a *MazeFloorUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazeFloorUnlock.json")
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
func (a *MazeFloorUnlockAccessor) Raw() ([]MazeFloorUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazeFloorUnlock{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MazeFloorUnlockAccessor.LoadData to preload everything
func (a *MazeFloorUnlockAccessor) GroupData() {
	a._dataFloorID = map[json.Number]MazeFloorUnlock{}
	for _, d := range a._data {
		a._dataFloorID[d.FloorID] = d
	}
}

// ByFloorID returns the MazeFloorUnlock uniquely identified by FloorID
//
// Error is only non-nil if the source errors out
func (a *MazeFloorUnlockAccessor) ByFloorID(identifier json.Number) (MazeFloorUnlock, error) {
	if a._dataFloorID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MazeFloorUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFloorID[identifier], nil
}
