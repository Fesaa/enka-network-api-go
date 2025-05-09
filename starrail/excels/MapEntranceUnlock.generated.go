package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MapEntranceUnlock struct {
	EntranceID                float64 `json:"EntranceID"`
	UnlockConditionExpression string  `json:"UnlockConditionExpression"`
}
type MapEntranceUnlockAccessor struct {
	_data           []MapEntranceUnlock
	_dataEntranceID map[float64]MapEntranceUnlock
}

// LoadData retrieves the data. Must be called before MapEntranceUnlock.GroupData
func (a *MapEntranceUnlockAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MapEntranceUnlock.json")
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
func (a *MapEntranceUnlockAccessor) Raw() ([]MapEntranceUnlock, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MapEntranceUnlock{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MapEntranceUnlockAccessor.LoadData to preload everything
func (a *MapEntranceUnlockAccessor) GroupData() {
	a._dataEntranceID = map[float64]MapEntranceUnlock{}
	for _, d := range a._data {
		a._dataEntranceID[d.EntranceID] = d
	}
}

// ByEntranceID returns the MapEntranceUnlock uniquely identified by EntranceID
//
// Error is only non-nil if the source errors out
func (a *MapEntranceUnlockAccessor) ByEntranceID(identifier float64) (MapEntranceUnlock, error) {
	if a._dataEntranceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MapEntranceUnlock{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEntranceID[identifier], nil
}
