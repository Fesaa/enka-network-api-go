package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityQuestRewardData struct {
	ActivityModuleID float64   `json:"ActivityModuleID"`
	QuestList        []float64 `json:"QuestList"`
	QuestTabID       float64   `json:"QuestTabID"`
	QuestTabName     hash.Hash `json:"QuestTabName"`
}
type ActivityQuestRewardDataAccessor struct {
	_data           []ActivityQuestRewardData
	_dataQuestTabID map[float64]ActivityQuestRewardData
}

// LoadData retrieves the data. Must be called before ActivityQuestRewardData.GroupData
func (a *ActivityQuestRewardDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityQuestRewardData.json")
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
func (a *ActivityQuestRewardDataAccessor) Raw() ([]ActivityQuestRewardData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityQuestRewardData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityQuestRewardDataAccessor.LoadData to preload everything
func (a *ActivityQuestRewardDataAccessor) GroupData() {
	a._dataQuestTabID = map[float64]ActivityQuestRewardData{}
	for _, d := range a._data {
		a._dataQuestTabID[d.QuestTabID] = d
	}
}

// ByQuestTabID returns the ActivityQuestRewardData uniquely identified by QuestTabID
//
// Error is only non-nil if the source errors out
func (a *ActivityQuestRewardDataAccessor) ByQuestTabID(identifier float64) (ActivityQuestRewardData, error) {
	if a._dataQuestTabID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityQuestRewardData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataQuestTabID[identifier], nil
}
