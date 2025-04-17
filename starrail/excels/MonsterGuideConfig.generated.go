package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MonsterGuideConfig struct {
	Difficulty          json.Number   `json:"Difficulty"`
	DifficultyGuideList []json.Number `json:"DifficultyGuideList"`
	DifficultyList      []json.Number `json:"DifficultyList"`
	MonsterID           json.Number   `json:"MonsterID"`
	PhaseList           []json.Number `json:"PhaseList"`
	TagList             []json.Number `json:"TagList"`
	TextGuideList       []json.Number `json:"TextGuideList"`
}
type MonsterGuideConfigAccessor struct {
	_data []MonsterGuideConfig
}

// LoadData retrieves the data. Must be called before MonsterGuideConfig.GroupData
func (a *MonsterGuideConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterGuideConfig.json")
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
func (a *MonsterGuideConfigAccessor) Raw() ([]MonsterGuideConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterGuideConfig{}, err
		}
	}
	return a._data, nil
}
