package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type QuestKeyPointReward struct {
	ID                  float64 `json:"ID"`
	QuestKeyPoint       float64 `json:"QuestKeyPoint"`
	QuestKeyPointItem   float64 `json:"QuestKeyPointItem"`
	QuestKeyPointReward float64 `json:"QuestKeyPointReward"`
}
type QuestKeyPointRewardAccessor struct {
	_data   []QuestKeyPointReward
	_dataID map[float64]QuestKeyPointReward
}

// LoadData retrieves the data. Must be called before QuestKeyPointReward.GroupData
func (a *QuestKeyPointRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/QuestKeyPointReward.json")
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
func (a *QuestKeyPointRewardAccessor) Raw() ([]QuestKeyPointReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []QuestKeyPointReward{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with QuestKeyPointRewardAccessor.LoadData to preload everything
func (a *QuestKeyPointRewardAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the QuestKeyPointReward uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *QuestKeyPointRewardAccessor) ByID(identifier float64) (QuestKeyPointReward, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return QuestKeyPointReward{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
