package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueMazeBuff struct {
	BuffDesc                     map[string]json.Number   `json:"BuffDesc"`
	BuffDescBattle               map[string]json.Number   `json:"BuffDescBattle"`
	BuffDescParamByAvatarSkillID json.Number              `json:"BuffDescParamByAvatarSkillID"`
	BuffEffect                   string                   `json:"BuffEffect"`
	BuffIcon                     string                   `json:"BuffIcon"`
	BuffName                     map[string]json.Number   `json:"BuffName"`
	BuffRarity                   json.Number              `json:"BuffRarity"`
	BuffSeries                   json.Number              `json:"BuffSeries"`
	BuffSimpleDesc               map[string]json.Number   `json:"BuffSimpleDesc"`
	ID                           json.Number              `json:"ID"`
	InBattleBindingKey           string                   `json:"InBattleBindingKey"`
	InBattleBindingType          string                   `json:"InBattleBindingType"`
	Lv                           json.Number              `json:"Lv"`
	LvMax                        json.Number              `json:"LvMax"`
	MazeBuffType                 string                   `json:"MazeBuffType"`
	ModifierName                 string                   `json:"ModifierName"`
	ParamList                    []map[string]json.Number `json:"ParamList"`
}
type RogueMazeBuffAccessor struct {
	_data []RogueMazeBuff
}

// LoadData retrieves the data. Must be called before RogueMazeBuff.GroupData
func (a *RogueMazeBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMazeBuff.json")
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
func (a *RogueMazeBuffAccessor) Raw() ([]RogueMazeBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMazeBuff{}, err
		}
	}
	return a._data, nil
}
