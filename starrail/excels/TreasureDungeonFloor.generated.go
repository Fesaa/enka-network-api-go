package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TreasureDungeonFloor struct {
	AddExploreValue  float64   `json:"AddExploreValue"`
	DungeonBuffID    []float64 `json:"DungeonBuffID"`
	DungeonID        float64   `json:"DungeonID"`
	EliteGroup2      float64   `json:"EliteGroup2"`
	FloorID          float64   `json:"FloorID"`
	HardLevelGroupID float64   `json:"HardLevelGroupID"`
	HardLevelList    []float64 `json:"HardLevelList"`
	MapID            []float64 `json:"MapID"`
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
