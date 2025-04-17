package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicMazeBuff struct {
	BuffDesc            hash.Hash    `json:"BuffDesc"`
	BuffEffect          string       `json:"BuffEffect"`
	BuffIcon            string       `json:"BuffIcon"`
	BuffName            hash.Hash    `json:"BuffName"`
	BuffRarity          float64      `json:"BuffRarity"`
	BuffSeries          float64      `json:"BuffSeries"`
	ID                  float64      `json:"ID"`
	InBattleBindingKey  string       `json:"InBattleBindingKey"`
	InBattleBindingType string       `json:"InBattleBindingType"`
	Lv                  float64      `json:"Lv"`
	LvMax               float64      `json:"LvMax"`
	MazeBuffType        string       `json:"MazeBuffType"`
	ModifierName        string       `json:"ModifierName"`
	ParamList           []hash.Value `json:"ParamList"`
}
type RogueMagicMazeBuffAccessor struct {
	_data []RogueMagicMazeBuff
}

// LoadData retrieves the data. Must be called before RogueMagicMazeBuff.GroupData
func (a *RogueMagicMazeBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicMazeBuff.json")
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
func (a *RogueMagicMazeBuffAccessor) Raw() ([]RogueMagicMazeBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicMazeBuff{}, err
		}
	}
	return a._data, nil
}
