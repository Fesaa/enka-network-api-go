package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueNousMissionReward struct {
	MissionRewardID json.Number            `json:"MissionRewardID"`
	QuestList       []json.Number          `json:"QuestList"`
	TabTitle        map[string]json.Number `json:"TabTitle"`
}
type RogueNousMissionRewardAccessor struct {
	_data                []RogueNousMissionReward
	_dataMissionRewardID map[json.Number]RogueNousMissionReward
}

// LoadData retrieves the data. Must be called before RogueNousMissionReward.GroupData
func (a *RogueNousMissionRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousMissionReward.json")
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
func (a *RogueNousMissionRewardAccessor) Raw() ([]RogueNousMissionReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousMissionReward{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueNousMissionRewardAccessor.LoadData to preload everything
func (a *RogueNousMissionRewardAccessor) GroupData() {
	a._dataMissionRewardID = map[json.Number]RogueNousMissionReward{}
	for _, d := range a._data {
		a._dataMissionRewardID[d.MissionRewardID] = d
	}
}

// ByMissionRewardID returns the RogueNousMissionReward uniquely identified by MissionRewardID
//
// Error is only non-nil if the source errors out
func (a *RogueNousMissionRewardAccessor) ByMissionRewardID(identifier json.Number) (RogueNousMissionReward, error) {
	if a._dataMissionRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueNousMissionReward{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionRewardID[identifier], nil
}
