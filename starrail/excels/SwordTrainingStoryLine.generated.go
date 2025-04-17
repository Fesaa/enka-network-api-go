package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type SwordTrainingStoryLine struct {
	AvatarIDList      []interface{}          `json:"AvatarIDList"`
	EndingOptionKey   string                 `json:"EndingOptionKey"`
	EndingStoryIDList []json.Number          `json:"EndingStoryIDList"`
	RewardID          json.Number            `json:"RewardID"`
	StartTalkImage    string                 `json:"StartTalkImage"`
	StoryHardDesc     map[string]json.Number `json:"StoryHardDesc"`
	StoryLine         json.Number            `json:"StoryLine"`
	StoryLineDesc     map[string]json.Number `json:"StoryLineDesc"`
	StoryLineImage    string                 `json:"StoryLineImage"`
	StoryLineName     map[string]json.Number `json:"StoryLineName"`
	TurnIDList        []json.Number          `json:"TurnIDList"`
	UnlockID          json.Number            `json:"UnlockID"`
}
type SwordTrainingStoryLineAccessor struct {
	_data                []SwordTrainingStoryLine
	_dataEndingOptionKey map[string]SwordTrainingStoryLine
}

// LoadData retrieves the data. Must be called before SwordTrainingStoryLine.GroupData
func (a *SwordTrainingStoryLineAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingStoryLine.json")
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
func (a *SwordTrainingStoryLineAccessor) Raw() ([]SwordTrainingStoryLine, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingStoryLine{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingStoryLineAccessor.LoadData to preload everything
func (a *SwordTrainingStoryLineAccessor) GroupData() {
	a._dataEndingOptionKey = map[string]SwordTrainingStoryLine{}
	for _, d := range a._data {
		a._dataEndingOptionKey[d.EndingOptionKey] = d
	}
}

// ByEndingOptionKey returns the SwordTrainingStoryLine uniquely identified by EndingOptionKey
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingStoryLineAccessor) ByEndingOptionKey(identifier string) (SwordTrainingStoryLine, error) {
	if a._dataEndingOptionKey == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingStoryLine{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEndingOptionKey[identifier], nil
}
