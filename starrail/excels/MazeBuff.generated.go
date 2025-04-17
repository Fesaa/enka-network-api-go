package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MazeBuff struct {
	BuffDesc            map[string]json.Number   `json:"BuffDesc"`
	BuffDescBattle      map[string]json.Number   `json:"BuffDescBattle"`
	BuffEffect          string                   `json:"BuffEffect"`
	BuffIcon            string                   `json:"BuffIcon"`
	BuffName            map[string]json.Number   `json:"BuffName"`
	BuffRarity          json.Number              `json:"BuffRarity"`
	BuffSeries          json.Number              `json:"BuffSeries"`
	BuffSimpleDesc      map[string]json.Number   `json:"BuffSimpleDesc"`
	ID                  json.Number              `json:"ID"`
	InBattleBindingKey  string                   `json:"InBattleBindingKey"`
	InBattleBindingType string                   `json:"InBattleBindingType"`
	IsDisplay           bool                     `json:"IsDisplay"`
	IsDisplayEnvInLevel bool                     `json:"IsDisplayEnvInLevel"`
	Lv                  json.Number              `json:"Lv"`
	LvMax               json.Number              `json:"LvMax"`
	MazeBuffIconType    string                   `json:"MazeBuffIconType"`
	MazeBuffPool        json.Number              `json:"MazeBuffPool"`
	MazeBuffType        string                   `json:"MazeBuffType"`
	ModifierName        string                   `json:"ModifierName"`
	ParamList           []map[string]json.Number `json:"ParamList"`
}
type MazeBuffAccessor struct {
	_data []MazeBuff
}

// LoadData retrieves the data. Must be called before MazeBuff.GroupData
func (a *MazeBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MazeBuff.json")
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
func (a *MazeBuffAccessor) Raw() ([]MazeBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MazeBuff{}, err
		}
	}
	return a._data, nil
}
