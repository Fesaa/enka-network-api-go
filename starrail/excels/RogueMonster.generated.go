package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueMonster struct {
	EventID         json.Number `json:"EventID"`
	MonsterDropType string      `json:"MonsterDropType"`
	NpcMonsterID    json.Number `json:"NpcMonsterID"`
	RogueMonsterID  json.Number `json:"RogueMonsterID"`
}
type RogueMonsterAccessor struct {
	_data []RogueMonster
}

// LoadData retrieves the data. Must be called before RogueMonster.GroupData
func (a *RogueMonsterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMonster.json")
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
func (a *RogueMonsterAccessor) Raw() ([]RogueMonster, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMonster{}, err
		}
	}
	return a._data, nil
}
