package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ChallengeMazeConfig struct {
	ChallengeCountDown json.Number            `json:"ChallengeCountDown"`
	ChallengeTargetID  []json.Number          `json:"ChallengeTargetID"`
	ConfigList1        []json.Number          `json:"ConfigList1"`
	ConfigList2        []json.Number          `json:"ConfigList2"`
	DamageType1        []string               `json:"DamageType1"`
	DamageType2        []string               `json:"DamageType2"`
	EventIDList1       []json.Number          `json:"EventIDList1"`
	EventIDList2       []json.Number          `json:"EventIDList2"`
	Floor              json.Number            `json:"Floor"`
	GroupID            json.Number            `json:"GroupID"`
	ID                 json.Number            `json:"ID"`
	MapEntranceID      json.Number            `json:"MapEntranceID"`
	MapEntranceID2     json.Number            `json:"MapEntranceID2"`
	MazeBuffID         json.Number            `json:"MazeBuffID"`
	MazeGroupID1       json.Number            `json:"MazeGroupID1"`
	MazeGroupID2       json.Number            `json:"MazeGroupID2"`
	MonsterID1         []interface{}          `json:"MonsterID1"`
	MonsterID2         []json.Number          `json:"MonsterID2"`
	Name               map[string]json.Number `json:"Name"`
	NpcMonsterIDList1  []json.Number          `json:"NpcMonsterIDList1"`
	NpcMonsterIDList2  []json.Number          `json:"NpcMonsterIDList2"`
	PreChallengeMazeID json.Number            `json:"PreChallengeMazeID"`
	PreLevel           json.Number            `json:"PreLevel"`
	RewardID           json.Number            `json:"RewardID"`
	StageNum           json.Number            `json:"StageNum"`
}
type ChallengeMazeConfigAccessor struct {
	_data []ChallengeMazeConfig
}

// LoadData retrieves the data. Must be called before ChallengeMazeConfig.GroupData
func (a *ChallengeMazeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeMazeConfig.json")
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
func (a *ChallengeMazeConfigAccessor) Raw() ([]ChallengeMazeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeMazeConfig{}, err
		}
	}
	return a._data, nil
}
