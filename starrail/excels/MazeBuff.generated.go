package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MazeBuff struct {
	BuffDesc            hash.Hash       `json:"BuffDesc"`
	BuffDescBattle      hash.Hash       `json:"BuffDescBattle"`
	BuffEffect          string          `json:"BuffEffect"`
	BuffIcon            string          `json:"BuffIcon"`
	BuffName            hash.Hash       `json:"BuffName"`
	BuffRarity          float64         `json:"BuffRarity"`
	BuffSeries          float64         `json:"BuffSeries"`
	BuffSimpleDesc      hash.Hash       `json:"BuffSimpleDesc"`
	ID                  float64         `json:"ID"`
	InBattleBindingKey  string          `json:"InBattleBindingKey"`
	InBattleBindingType string          `json:"InBattleBindingType"`
	IsDisplay           bool            `json:"IsDisplay"`
	IsDisplayEnvInLevel bool            `json:"IsDisplayEnvInLevel"`
	Lv                  float64         `json:"Lv"`
	LvMax               float64         `json:"LvMax"`
	MazeBuffIconType    string          `json:"MazeBuffIconType"`
	MazeBuffPool        float64         `json:"MazeBuffPool"`
	MazeBuffType        string          `json:"MazeBuffType"`
	ModifierName        string          `json:"ModifierName"`
	ParamList           []hash.IntValue `json:"ParamList"`
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
		a.GroupData()
	}
	return a._data, nil
}
