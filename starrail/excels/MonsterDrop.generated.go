package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MonsterDrop struct {
	AvatarExpReward   json.Number              `json:"AvatarExpReward"`
	DisplayItemList   []map[string]json.Number `json:"DisplayItemList"`
	MonsterTemplateID json.Number              `json:"MonsterTemplateID"`
	WorldLevel        json.Number              `json:"WorldLevel"`
}
type MonsterDropAccessor struct {
	_data []MonsterDrop
}

// LoadData retrieves the data. Must be called before MonsterDrop.GroupData
func (a *MonsterDropAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterDrop.json")
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
func (a *MonsterDropAccessor) Raw() ([]MonsterDrop, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterDrop{}, err
		}
	}
	return a._data, nil
}
