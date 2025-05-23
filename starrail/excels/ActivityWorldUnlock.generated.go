package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityWorldUnlock struct {
	ActivityID float64 `json:"ActivityID"`
	WorldID    float64 `json:"WorldID"`
}
type ActivityWorldUnlockAccessor struct {
	_data           []ActivityWorldUnlock
	_dataActivityID map[float64]ActivityWorldUnlock
}

// LoadData retrieves the data. Must be called before ActivityWorldUnlock.GroupData
func (a *ActivityWorldUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityWorldUnlock.json")
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
func (a *ActivityWorldUnlockAccessor) Raw() ([]ActivityWorldUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityWorldUnlock{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityWorldUnlockAccessor.LoadData to preload everything
func (a *ActivityWorldUnlockAccessor) GroupData() {
	a._dataActivityID = map[float64]ActivityWorldUnlock{}
	for _, d := range a._data {
		a._dataActivityID[d.ActivityID] = d
	}
}

// ByActivityID returns the ActivityWorldUnlock uniquely identified by ActivityID
//
// Error is only non-nil if the source errors out
func (a *ActivityWorldUnlockAccessor) ByActivityID(identifier float64) (ActivityWorldUnlock, error) {
	if a._dataActivityID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityWorldUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityID[identifier], nil
}
