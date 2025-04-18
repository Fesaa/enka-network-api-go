package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityRaidCollectionGroup struct {
	GroupEntrancePrefabPath       string    `json:"GroupEntrancePrefabPath"`
	RaidCollectionGroupID         float64   `json:"RaidCollectionGroupID"`
	RaidCollectionGroupName       hash.Hash `json:"RaidCollectionGroupName"`
	RaidCollectionGroupNextEnable bool      `json:"RaidCollectionGroupNextEnable"`
	RaidCollectionList            []float64 `json:"RaidCollectionList"`
	UnlockGroupID                 float64   `json:"UnlockGroupID"`
}
type ActivityRaidCollectionGroupAccessor struct {
	_data                        []ActivityRaidCollectionGroup
	_dataGroupEntrancePrefabPath map[string]ActivityRaidCollectionGroup
	_dataRaidCollectionGroupID   map[float64]ActivityRaidCollectionGroup
}

// LoadData retrieves the data. Must be called before ActivityRaidCollectionGroup.GroupData
func (a *ActivityRaidCollectionGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRaidCollectionGroup.json")
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
func (a *ActivityRaidCollectionGroupAccessor) Raw() ([]ActivityRaidCollectionGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRaidCollectionGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityRaidCollectionGroupAccessor.LoadData to preload everything
func (a *ActivityRaidCollectionGroupAccessor) GroupData() {
	a._dataGroupEntrancePrefabPath = map[string]ActivityRaidCollectionGroup{}
	a._dataRaidCollectionGroupID = map[float64]ActivityRaidCollectionGroup{}
	for _, d := range a._data {
		a._dataGroupEntrancePrefabPath[d.GroupEntrancePrefabPath] = d
		a._dataRaidCollectionGroupID[d.RaidCollectionGroupID] = d
	}
}

// ByGroupEntrancePrefabPath returns the ActivityRaidCollectionGroup uniquely identified by GroupEntrancePrefabPath
//
// Error is only non-nil if the source errors out
func (a *ActivityRaidCollectionGroupAccessor) ByGroupEntrancePrefabPath(identifier string) (ActivityRaidCollectionGroup, error) {
	if a._dataGroupEntrancePrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRaidCollectionGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupEntrancePrefabPath[identifier], nil
}

// ByRaidCollectionGroupID returns the ActivityRaidCollectionGroup uniquely identified by RaidCollectionGroupID
//
// Error is only non-nil if the source errors out
func (a *ActivityRaidCollectionGroupAccessor) ByRaidCollectionGroupID(identifier float64) (ActivityRaidCollectionGroup, error) {
	if a._dataRaidCollectionGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRaidCollectionGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidCollectionGroupID[identifier], nil
}
