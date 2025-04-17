package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueNousStoryReward struct {
	IsImportant     bool    `json:"IsImportant"`
	MainStoryReward float64 `json:"MainStoryReward"`
	QuestID         float64 `json:"QuestID"`
}
type RogueNousStoryRewardAccessor struct {
	_data                []RogueNousStoryReward
	_dataMainStoryReward map[float64]RogueNousStoryReward
	_dataQuestID         map[float64]RogueNousStoryReward
}

// LoadData retrieves the data. Must be called before RogueNousStoryReward.GroupData
func (a *RogueNousStoryRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousStoryReward.json")
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
func (a *RogueNousStoryRewardAccessor) Raw() ([]RogueNousStoryReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousStoryReward{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueNousStoryRewardAccessor.LoadData to preload everything
func (a *RogueNousStoryRewardAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMainStoryReward[d.MainStoryReward] = d
		a._dataQuestID[d.QuestID] = d
	}
}

// ByMainStoryReward returns the RogueNousStoryReward uniquely identified by MainStoryReward
//
// Error is only non-nil if the source errors out
func (a *RogueNousStoryRewardAccessor) ByMainStoryReward(identifier float64) (RogueNousStoryReward, error) {
	if a._dataMainStoryReward == nil {
		err := a.LoadData()
		if err != nil {
			return RogueNousStoryReward{}, err
		}
		a.GroupData()
	}
	return a._dataMainStoryReward[identifier], nil
}

// ByQuestID returns the RogueNousStoryReward uniquely identified by QuestID
//
// Error is only non-nil if the source errors out
func (a *RogueNousStoryRewardAccessor) ByQuestID(identifier float64) (RogueNousStoryReward, error) {
	if a._dataQuestID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueNousStoryReward{}, err
		}
		a.GroupData()
	}
	return a._dataQuestID[identifier], nil
}
