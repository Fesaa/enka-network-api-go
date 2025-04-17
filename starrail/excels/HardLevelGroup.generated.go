package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type HardLevelGroup struct {
	AttackRatio       map[string]json.Number   `json:"AttackRatio"`
	CombatPowerList   []map[string]json.Number `json:"CombatPowerList"`
	DefenceRatio      map[string]json.Number   `json:"DefenceRatio"`
	HPRatio           map[string]json.Number   `json:"HPRatio"`
	HardLevelGroup    json.Number              `json:"HardLevelGroup"`
	Level             json.Number              `json:"Level"`
	SpeedRatio        map[string]json.Number   `json:"SpeedRatio"`
	StanceRatio       map[string]json.Number   `json:"StanceRatio"`
	StatusProbability map[string]json.Number   `json:"StatusProbability"`
	StatusResistance  map[string]json.Number   `json:"StatusResistance"`
}
type HardLevelGroupAccessor struct {
	_data []HardLevelGroup
}

// LoadData retrieves the data. Must be called before HardLevelGroup.GroupData
func (a *HardLevelGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HardLevelGroup.json")
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
func (a *HardLevelGroupAccessor) Raw() ([]HardLevelGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HardLevelGroup{}, err
		}
	}
	return a._data, nil
}
