package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type EvolveBuildMonsteCollection struct {
	ID          float64 `json:"ID"`
	UnlockQuest float64 `json:"UnlockQuest"`
}
type EvolveBuildMonsteCollectionAccessor struct {
	_data            []EvolveBuildMonsteCollection
	_dataID          map[float64]EvolveBuildMonsteCollection
	_dataUnlockQuest map[float64]EvolveBuildMonsteCollection
}

// LoadData retrieves the data. Must be called before EvolveBuildMonsteCollection.GroupData
func (a *EvolveBuildMonsteCollectionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildMonsteCollection.json")
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
func (a *EvolveBuildMonsteCollectionAccessor) Raw() ([]EvolveBuildMonsteCollection, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildMonsteCollection{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildMonsteCollectionAccessor.LoadData to preload everything
func (a *EvolveBuildMonsteCollectionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataUnlockQuest[d.UnlockQuest] = d
	}
}

// ByID returns the EvolveBuildMonsteCollection uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildMonsteCollectionAccessor) ByID(identifier float64) (EvolveBuildMonsteCollection, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildMonsteCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByUnlockQuest returns the EvolveBuildMonsteCollection uniquely identified by UnlockQuest
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildMonsteCollectionAccessor) ByUnlockQuest(identifier float64) (EvolveBuildMonsteCollection, error) {
	if a._dataUnlockQuest == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildMonsteCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockQuest[identifier], nil
}
