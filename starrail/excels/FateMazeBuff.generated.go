package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type FateMazeBuff struct {
	BuffDesc            hash.Hash    `json:"BuffDesc"`
	BuffEffect          string       `json:"BuffEffect"`
	BuffIcon            string       `json:"BuffIcon"`
	BuffName            hash.Hash    `json:"BuffName"`
	BuffRarity          float64      `json:"BuffRarity"`
	BuffSeries          float64      `json:"BuffSeries"`
	BuffSimpleDesc      hash.Hash    `json:"BuffSimpleDesc"`
	ID                  float64      `json:"ID"`
	InBattleBindingKey  string       `json:"InBattleBindingKey"`
	InBattleBindingType string       `json:"InBattleBindingType"`
	IsDisplay           bool         `json:"IsDisplay"`
	IsDisplayEnvInLevel bool         `json:"IsDisplayEnvInLevel"`
	Lv                  float64      `json:"Lv"`
	LvMax               float64      `json:"LvMax"`
	MazeBuffIconType    string       `json:"MazeBuffIconType"`
	MazeBuffType        string       `json:"MazeBuffType"`
	ModifierName        string       `json:"ModifierName"`
	ParamList           []hash.Value `json:"ParamList"`
}
type FateMazeBuffAccessor struct {
	_data []FateMazeBuff
}

// LoadData retrieves the data. Must be called before FateMazeBuff.GroupData
func (a *FateMazeBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FateMazeBuff.json")
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
func (a *FateMazeBuffAccessor) Raw() ([]FateMazeBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FateMazeBuff{}, err
		}
	}
	return a._data, nil
}
