package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournExpReward struct {
	Exp         float64 `json:"Exp"`
	Level       float64 `json:"Level"`
	MainTournID float64 `json:"MainTournID"`
	RewardID    float64 `json:"RewardID"`
}
type RogueTournExpRewardAccessor struct {
	_data         []RogueTournExpReward
	_dataRewardID map[float64]RogueTournExpReward
}

// LoadData retrieves the data. Must be called before RogueTournExpReward.GroupData
func (a *RogueTournExpRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournExpReward.json")
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
func (a *RogueTournExpRewardAccessor) Raw() ([]RogueTournExpReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournExpReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournExpRewardAccessor.LoadData to preload everything
func (a *RogueTournExpRewardAccessor) GroupData() {
	for _, d := range a._data {
		a._dataRewardID[d.RewardID] = d
	}
}

// ByRewardID returns the RogueTournExpReward uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *RogueTournExpRewardAccessor) ByRewardID(identifier float64) (RogueTournExpReward, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournExpReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}
