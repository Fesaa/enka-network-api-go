package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TreasureDungeonFloor struct {
	AddExploreValue  json.Number   `json:"AddExploreValue"`
	DungeonBuffID    []json.Number `json:"DungeonBuffID"`
	DungeonID        json.Number   `json:"DungeonID"`
	EliteGroup2      json.Number   `json:"EliteGroup2"`
	FloorID          json.Number   `json:"FloorID"`
	HardLevelGroupID json.Number   `json:"HardLevelGroupID"`
	HardLevelList    []json.Number `json:"HardLevelList"`
	MapID            []json.Number `json:"MapID"`
}
type TreasureDungeonFloorAccessor struct {
	_data []TreasureDungeonFloor
}

// LoadData retrieves the data. Must be called before TreasureDungeonFloor.GroupData
func (a *TreasureDungeonFloorAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TreasureDungeonFloor.json")
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
func (a *TreasureDungeonFloorAccessor) Raw() ([]TreasureDungeonFloor, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TreasureDungeonFloor{}, err
		}
	}
	return a._data, nil
}
