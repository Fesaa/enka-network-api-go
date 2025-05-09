package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityRaidCollection struct {
	GuideID          float64 `json:"GuideID"`
	PrepareType      string  `json:"PrepareType"`
	RaidCollectionID float64 `json:"RaidCollectionID"`
	RaidID           float64 `json:"RaidID"`
	SubMissionID     float64 `json:"SubMissionID"`
}
type ActivityRaidCollectionAccessor struct {
	_data                 []ActivityRaidCollection
	_dataRaidCollectionID map[float64]ActivityRaidCollection
	_dataRaidID           map[float64]ActivityRaidCollection
	_dataSubMissionID     map[float64]ActivityRaidCollection
}

// LoadData retrieves the data. Must be called before ActivityRaidCollection.GroupData
func (a *ActivityRaidCollectionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRaidCollection.json")
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
func (a *ActivityRaidCollectionAccessor) Raw() ([]ActivityRaidCollection, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRaidCollection{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityRaidCollectionAccessor.LoadData to preload everything
func (a *ActivityRaidCollectionAccessor) GroupData() {
	a._dataRaidCollectionID = map[float64]ActivityRaidCollection{}
	a._dataRaidID = map[float64]ActivityRaidCollection{}
	a._dataSubMissionID = map[float64]ActivityRaidCollection{}
	for _, d := range a._data {
		a._dataRaidCollectionID[d.RaidCollectionID] = d
		a._dataRaidID[d.RaidID] = d
		a._dataSubMissionID[d.SubMissionID] = d
	}
}

// ByRaidCollectionID returns the ActivityRaidCollection uniquely identified by RaidCollectionID
//
// Error is only non-nil if the source errors out
func (a *ActivityRaidCollectionAccessor) ByRaidCollectionID(identifier float64) (ActivityRaidCollection, error) {
	if a._dataRaidCollectionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRaidCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidCollectionID[identifier], nil
}

// ByRaidID returns the ActivityRaidCollection uniquely identified by RaidID
//
// Error is only non-nil if the source errors out
func (a *ActivityRaidCollectionAccessor) ByRaidID(identifier float64) (ActivityRaidCollection, error) {
	if a._dataRaidID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRaidCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidID[identifier], nil
}

// BySubMissionID returns the ActivityRaidCollection uniquely identified by SubMissionID
//
// Error is only non-nil if the source errors out
func (a *ActivityRaidCollectionAccessor) BySubMissionID(identifier float64) (ActivityRaidCollection, error) {
	if a._dataSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRaidCollection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSubMissionID[identifier], nil
}
