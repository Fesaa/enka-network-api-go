package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActivityQuestRewardConfig struct {
	ActivityModule    float64   `json:"ActivityModule"`
	ActivityRewardID  float64   `json:"ActivityRewardID"`
	FinalRewardQuest  float64   `json:"FinalRewardQuest"`
	QuestTabGroupList []float64 `json:"QuestTabGroupList"`
}
type ActivityQuestRewardConfigAccessor struct {
	_data                 []ActivityQuestRewardConfig
	_dataActivityModule   map[float64]ActivityQuestRewardConfig
	_dataActivityRewardID map[float64]ActivityQuestRewardConfig
}

// LoadData retrieves the data. Must be called before ActivityQuestRewardConfig.GroupData
func (a *ActivityQuestRewardConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityQuestRewardConfig.json")
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
func (a *ActivityQuestRewardConfigAccessor) Raw() ([]ActivityQuestRewardConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityQuestRewardConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityQuestRewardConfigAccessor.LoadData to preload everything
func (a *ActivityQuestRewardConfigAccessor) GroupData() {
	a._dataActivityModule = map[float64]ActivityQuestRewardConfig{}
	a._dataActivityRewardID = map[float64]ActivityQuestRewardConfig{}
	for _, d := range a._data {
		a._dataActivityModule[d.ActivityModule] = d
		a._dataActivityRewardID[d.ActivityRewardID] = d
	}
}

// ByActivityModule returns the ActivityQuestRewardConfig uniquely identified by ActivityModule
//
// Error is only non-nil if the source errors out
func (a *ActivityQuestRewardConfigAccessor) ByActivityModule(identifier float64) (ActivityQuestRewardConfig, error) {
	if a._dataActivityModule == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityQuestRewardConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityModule[identifier], nil
}

// ByActivityRewardID returns the ActivityQuestRewardConfig uniquely identified by ActivityRewardID
//
// Error is only non-nil if the source errors out
func (a *ActivityQuestRewardConfigAccessor) ByActivityRewardID(identifier float64) (ActivityQuestRewardConfig, error) {
	if a._dataActivityRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityQuestRewardConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActivityRewardID[identifier], nil
}
