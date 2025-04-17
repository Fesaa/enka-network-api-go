package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ShareRewardData struct {
	ID        float64 `json:"ID"`
	RewardID  float64 `json:"RewardID"`
	RewardNum float64 `json:"RewardNum"`
}
type ShareRewardDataAccessor struct {
	_data          []ShareRewardData
	_dataID        map[float64]ShareRewardData
	_dataRewardID  map[float64]ShareRewardData
	_dataRewardNum map[float64]ShareRewardData
}

// LoadData retrieves the data. Must be called before ShareRewardData.GroupData
func (a *ShareRewardDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ShareRewardData.json")
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
func (a *ShareRewardDataAccessor) Raw() ([]ShareRewardData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ShareRewardData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ShareRewardDataAccessor.LoadData to preload everything
func (a *ShareRewardDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataRewardID[d.RewardID] = d
		a._dataRewardNum[d.RewardNum] = d
	}
}

// ByID returns the ShareRewardData uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ShareRewardDataAccessor) ByID(identifier float64) (ShareRewardData, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ShareRewardData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByRewardID returns the ShareRewardData uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *ShareRewardDataAccessor) ByRewardID(identifier float64) (ShareRewardData, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ShareRewardData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}

// ByRewardNum returns the ShareRewardData uniquely identified by RewardNum
//
// Error is only non-nil if the source errors out
func (a *ShareRewardDataAccessor) ByRewardNum(identifier float64) (ShareRewardData, error) {
	if a._dataRewardNum == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ShareRewardData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardNum[identifier], nil
}
