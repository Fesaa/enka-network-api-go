package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type TrainPartyRewardConfig struct {
	Level       json.Number            `json:"Level"`
	Name        map[string]json.Number `json:"Name"`
	RequireStar json.Number            `json:"RequireStar"`
	RewardID    json.Number            `json:"RewardID"`
}
type TrainPartyRewardConfigAccessor struct {
	_data []TrainPartyRewardConfig
}

// LoadData retrieves the data. Must be called before TrainPartyRewardConfig.GroupData
func (a *TrainPartyRewardConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyRewardConfig.json")
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
func (a *TrainPartyRewardConfigAccessor) Raw() ([]TrainPartyRewardConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyRewardConfig{}, err
		}
	}
	return a._data, nil
}
