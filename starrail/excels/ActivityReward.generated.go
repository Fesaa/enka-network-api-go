package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityReward struct {
	ActivityRewardID float64 `json:"ActivityRewardID"`
	Count            float64 `json:"Count"`
	Reward           float64 `json:"Reward"`
	RewardIconPath   string  `json:"RewardIconPath"`
}
type ActivityRewardAccessor struct {
	_data                 []ActivityReward
	_dataActivityRewardID map[float64]ActivityReward
	_dataCount            map[float64]ActivityReward
}

// LoadData retrieves the data. Must be called before ActivityReward.GroupData
func (a *ActivityRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityReward.json")
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
func (a *ActivityRewardAccessor) Raw() ([]ActivityReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityRewardAccessor.LoadData to preload everything
func (a *ActivityRewardAccessor) GroupData() {
	a._dataActivityRewardID = map[float64]ActivityReward{}
	a._dataCount = map[float64]ActivityReward{}
	for _, d := range a._data {
		a._dataActivityRewardID[d.ActivityRewardID] = d
		a._dataCount[d.Count] = d
	}
}

// ByActivityRewardID returns the ActivityReward uniquely identified by ActivityRewardID
//
// Error is only non-nil if the source errors out
func (a *ActivityRewardAccessor) ByActivityRewardID(identifier float64) (ActivityReward, error) {
	if a._dataActivityRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityRewardID[identifier], nil
}

// ByCount returns the ActivityReward uniquely identified by Count
//
// Error is only non-nil if the source errors out
func (a *ActivityRewardAccessor) ByCount(identifier float64) (ActivityReward, error) {
	if a._dataCount == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCount[identifier], nil
}
