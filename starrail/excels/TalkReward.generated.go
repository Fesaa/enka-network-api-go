package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TalkReward struct {
	FloorID        float64 `json:"FloorID"`
	GroupID        float64 `json:"GroupID"`
	ID             float64 `json:"ID"`
	NPCConfigID    float64 `json:"NPCConfigID"`
	PlaneID        float64 `json:"PlaneID"`
	PropConfigID   float64 `json:"PropConfigID"`
	RewardID       float64 `json:"RewardID"`
	VerificationID float64 `json:"VerificationID"`
}
type TalkRewardAccessor struct {
	_data         []TalkReward
	_dataID       map[float64]TalkReward
	_dataRewardID map[float64]TalkReward
}

// LoadData retrieves the data. Must be called before TalkReward.GroupData
func (a *TalkRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TalkReward.json")
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
func (a *TalkRewardAccessor) Raw() ([]TalkReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TalkReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TalkRewardAccessor.LoadData to preload everything
func (a *TalkRewardAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataRewardID[d.RewardID] = d
	}
}

// ByID returns the TalkReward uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TalkRewardAccessor) ByID(identifier float64) (TalkReward, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TalkReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByRewardID returns the TalkReward uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *TalkRewardAccessor) ByRewardID(identifier float64) (TalkReward, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TalkReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}
