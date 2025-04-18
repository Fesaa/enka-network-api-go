package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeBossRewardLine struct {
	GroupID   float64 `json:"GroupID"`
	RewardID  float64 `json:"RewardID"`
	StarCount float64 `json:"StarCount"`
}
type ChallengeBossRewardLineAccessor struct {
	_data          []ChallengeBossRewardLine
	_dataRewardID  map[float64]ChallengeBossRewardLine
	_dataStarCount map[float64]ChallengeBossRewardLine
}

// LoadData retrieves the data. Must be called before ChallengeBossRewardLine.GroupData
func (a *ChallengeBossRewardLineAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeBossRewardLine.json")
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
func (a *ChallengeBossRewardLineAccessor) Raw() ([]ChallengeBossRewardLine, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeBossRewardLine{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeBossRewardLineAccessor.LoadData to preload everything
func (a *ChallengeBossRewardLineAccessor) GroupData() {
	a._dataRewardID = map[float64]ChallengeBossRewardLine{}
	a._dataStarCount = map[float64]ChallengeBossRewardLine{}
	for _, d := range a._data {
		a._dataRewardID[d.RewardID] = d
		a._dataStarCount[d.StarCount] = d
	}
}

// ByRewardID returns the ChallengeBossRewardLine uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *ChallengeBossRewardLineAccessor) ByRewardID(identifier float64) (ChallengeBossRewardLine, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeBossRewardLine{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}

// ByStarCount returns the ChallengeBossRewardLine uniquely identified by StarCount
//
// Error is only non-nil if the source errors out
func (a *ChallengeBossRewardLineAccessor) ByStarCount(identifier float64) (ChallengeBossRewardLine, error) {
	if a._dataStarCount == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeBossRewardLine{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStarCount[identifier], nil
}
