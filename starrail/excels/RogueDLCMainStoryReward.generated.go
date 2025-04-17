package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCMainStoryReward struct {
	IsImportant     float64 `json:"IsImportant"`
	MainStoryID     float64 `json:"MainStoryID"`
	MainStoryReward float64 `json:"MainStoryReward"`
	QuestID         float64 `json:"QuestID"`
	Sort            float64 `json:"Sort"`
}
type RogueDLCMainStoryRewardAccessor struct {
	_data                []RogueDLCMainStoryReward
	_dataMainStoryReward map[float64]RogueDLCMainStoryReward
	_dataSort            map[float64]RogueDLCMainStoryReward
	_dataQuestID         map[float64]RogueDLCMainStoryReward
}

// LoadData retrieves the data. Must be called before RogueDLCMainStoryReward.GroupData
func (a *RogueDLCMainStoryRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCMainStoryReward.json")
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
func (a *RogueDLCMainStoryRewardAccessor) Raw() ([]RogueDLCMainStoryReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCMainStoryReward{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCMainStoryRewardAccessor.LoadData to preload everything
func (a *RogueDLCMainStoryRewardAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMainStoryReward[d.MainStoryReward] = d
		a._dataSort[d.Sort] = d
		a._dataQuestID[d.QuestID] = d
	}
}

// ByMainStoryReward returns the RogueDLCMainStoryReward uniquely identified by MainStoryReward
//
// Error is only non-nil if the source errors out
func (a *RogueDLCMainStoryRewardAccessor) ByMainStoryReward(identifier float64) (RogueDLCMainStoryReward, error) {
	if a._dataMainStoryReward == nil {
		err := a.LoadData()
		if err != nil {
			return RogueDLCMainStoryReward{}, err
		}
		a.GroupData()
	}
	return a._dataMainStoryReward[identifier], nil
}

// BySort returns the RogueDLCMainStoryReward uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *RogueDLCMainStoryRewardAccessor) BySort(identifier float64) (RogueDLCMainStoryReward, error) {
	if a._dataSort == nil {
		err := a.LoadData()
		if err != nil {
			return RogueDLCMainStoryReward{}, err
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}

// ByQuestID returns the RogueDLCMainStoryReward uniquely identified by QuestID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCMainStoryRewardAccessor) ByQuestID(identifier float64) (RogueDLCMainStoryReward, error) {
	if a._dataQuestID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueDLCMainStoryReward{}, err
		}
		a.GroupData()
	}
	return a._dataQuestID[identifier], nil
}
