package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type BattlePassReward struct {
	ID         json.Number `json:"ID"`
	IsNew      bool        `json:"IsNew"`
	NumShow    bool        `json:"NumShow"`
	RewardIcon string      `json:"RewardIcon"`
	RewardItem json.Number `json:"RewardItem"`
}
type BattlePassRewardAccessor struct {
	_data []BattlePassReward
}

// LoadData retrieves the data. Must be called before BattlePassReward.GroupData
func (a *BattlePassRewardAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattlePassReward.json")
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
func (a *BattlePassRewardAccessor) Raw() ([]BattlePassReward, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattlePassReward{}, err
		}
	}
	return a._data, nil
}
