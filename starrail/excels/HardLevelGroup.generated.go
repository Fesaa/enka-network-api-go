package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type HardLevelGroup struct {
	AttackRatio       hash.Value   `json:"AttackRatio"`
	CombatPowerList   []hash.Value `json:"CombatPowerList"`
	DefenceRatio      hash.Value   `json:"DefenceRatio"`
	HPRatio           hash.Value   `json:"HPRatio"`
	HardLevelGroup    float64      `json:"HardLevelGroup"`
	Level             float64      `json:"Level"`
	SpeedRatio        hash.Value   `json:"SpeedRatio"`
	StanceRatio       hash.Value   `json:"StanceRatio"`
	StatusProbability hash.Value   `json:"StatusProbability"`
	StatusResistance  hash.Value   `json:"StatusResistance"`
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
