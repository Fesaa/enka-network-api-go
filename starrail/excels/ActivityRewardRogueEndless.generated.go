package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityRewardRogueEndless struct {
	RewardID        float64   `json:"RewardID"`
	RewardLevel     float64   `json:"RewardLevel"`
	RewardLevelName hash.Hash `json:"RewardLevelName"`
	RewardPoint     float64   `json:"RewardPoint"`
}
type ActivityRewardRogueEndlessAccessor struct {
	_data            []ActivityRewardRogueEndless
	_dataRewardID    map[float64]ActivityRewardRogueEndless
	_dataRewardLevel map[float64]ActivityRewardRogueEndless
	_dataRewardPoint map[float64]ActivityRewardRogueEndless
}

// LoadData retrieves the data. Must be called before ActivityRewardRogueEndless.GroupData
func (a *ActivityRewardRogueEndlessAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRewardRogueEndless.json")
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
func (a *ActivityRewardRogueEndlessAccessor) Raw() ([]ActivityRewardRogueEndless, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRewardRogueEndless{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityRewardRogueEndlessAccessor.LoadData to preload everything
func (a *ActivityRewardRogueEndlessAccessor) GroupData() {
	a._dataRewardID = map[float64]ActivityRewardRogueEndless{}
	a._dataRewardLevel = map[float64]ActivityRewardRogueEndless{}
	a._dataRewardPoint = map[float64]ActivityRewardRogueEndless{}
	for _, d := range a._data {
		a._dataRewardID[d.RewardID] = d
		a._dataRewardLevel[d.RewardLevel] = d
		a._dataRewardPoint[d.RewardPoint] = d
	}
}

// ByRewardID returns the ActivityRewardRogueEndless uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *ActivityRewardRogueEndlessAccessor) ByRewardID(identifier float64) (ActivityRewardRogueEndless, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRewardRogueEndless{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}

// ByRewardLevel returns the ActivityRewardRogueEndless uniquely identified by RewardLevel
//
// Error is only non-nil if the source errors out
func (a *ActivityRewardRogueEndlessAccessor) ByRewardLevel(identifier float64) (ActivityRewardRogueEndless, error) {
	if a._dataRewardLevel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRewardRogueEndless{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardLevel[identifier], nil
}

// ByRewardPoint returns the ActivityRewardRogueEndless uniquely identified by RewardPoint
//
// Error is only non-nil if the source errors out
func (a *ActivityRewardRogueEndlessAccessor) ByRewardPoint(identifier float64) (ActivityRewardRogueEndless, error) {
	if a._dataRewardPoint == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRewardRogueEndless{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardPoint[identifier], nil
}
