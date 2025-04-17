package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type BattleEventData struct {
	BEActionBarPrefab string        `json:"BEActionBarPrefab"`
	BasePoint         string        `json:"BasePoint"`
	BattleEventID     json.Number   `json:"BattleEventID"`
	Config            string        `json:"Config"`
	IsSPReserved      bool          `json:"IsSPReserved"`
	LevelAreaPrefab   string        `json:"LevelAreaPrefab"`
	Prefab            string        `json:"Prefab"`
	SkillIDList       []json.Number `json:"SkillIDList"`
}
type BattleEventDataAccessor struct {
	_data []BattleEventData
}

// LoadData retrieves the data. Must be called before BattleEventData.GroupData
func (a *BattleEventDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleEventData.json")
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
func (a *BattleEventDataAccessor) Raw() ([]BattleEventData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleEventData{}, err
		}
	}
	return a._data, nil
}
