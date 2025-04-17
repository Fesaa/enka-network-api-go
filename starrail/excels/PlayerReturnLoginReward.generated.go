package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type PlayerReturnLoginReward struct {
	FirstWordText string      `json:"FirstWordText"`
	ID            json.Number `json:"ID"`
	LoginReward   json.Number `json:"LoginReward"`
}
type PlayerReturnLoginRewardAccessor struct {
	_data            []PlayerReturnLoginReward
	_dataID          map[json.Number]PlayerReturnLoginReward
	_dataLoginReward map[json.Number]PlayerReturnLoginReward
}

// LoadData retrieves the data. Must be called before PlayerReturnLoginReward.GroupData
func (a *PlayerReturnLoginRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnLoginReward.json")
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
func (a *PlayerReturnLoginRewardAccessor) Raw() ([]PlayerReturnLoginReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnLoginReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerReturnLoginRewardAccessor.LoadData to preload everything
func (a *PlayerReturnLoginRewardAccessor) GroupData() {
	a._dataID = map[json.Number]PlayerReturnLoginReward{}
	a._dataLoginReward = map[json.Number]PlayerReturnLoginReward{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataLoginReward[d.LoginReward] = d
	}
}

// ByID returns the PlayerReturnLoginReward uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnLoginRewardAccessor) ByID(identifier json.Number) (PlayerReturnLoginReward, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnLoginReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByLoginReward returns the PlayerReturnLoginReward uniquely identified by LoginReward
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnLoginRewardAccessor) ByLoginReward(identifier json.Number) (PlayerReturnLoginReward, error) {
	if a._dataLoginReward == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnLoginReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLoginReward[identifier], nil
}
