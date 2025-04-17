package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type StageInfiniteMonsterGroup struct {
	EliteGroup             json.Number   `json:"EliteGroup"`
	InfiniteMonsterGroupID json.Number   `json:"InfiniteMonsterGroupID"`
	MonsterList            []json.Number `json:"MonsterList"`
}
type StageInfiniteMonsterGroupAccessor struct {
	_data []StageInfiniteMonsterGroup
}

// LoadData retrieves the data. Must be called before StageInfiniteMonsterGroup.GroupData
func (a *StageInfiniteMonsterGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StageInfiniteMonsterGroup.json")
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
func (a *StageInfiniteMonsterGroupAccessor) Raw() ([]StageInfiniteMonsterGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StageInfiniteMonsterGroup{}, err
		}
	}
	return a._data, nil
}
