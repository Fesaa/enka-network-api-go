package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BattleArea struct {
	BattleAreaGroupID  float64 `json:"BattleAreaGroupID"`
	BattleAreaID       float64 `json:"BattleAreaID"`
	FloorBattleAreaID  float64 `json:"FloorBattleAreaID"`
	FloorID            float64 `json:"FloorID"`
	ID                 float64 `json:"ID"`
	IsLegacy           bool    `json:"IsLegacy"`
	IsUseUnifiedConfig bool    `json:"IsUseUnifiedConfig"`
	PlaneID            float64 `json:"PlaneID"`
	UnifiedConfigID    float64 `json:"UnifiedConfigID"`
}
type BattleAreaAccessor struct {
	_data   []BattleArea
	_dataID map[float64]BattleArea
}

// LoadData retrieves the data. Must be called before BattleArea.GroupData
func (a *BattleAreaAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleArea.json")
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
func (a *BattleAreaAccessor) Raw() ([]BattleArea, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleArea{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleAreaAccessor.LoadData to preload everything
func (a *BattleAreaAccessor) GroupData() {
	a._dataID = map[float64]BattleArea{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the BattleArea uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BattleAreaAccessor) ByID(identifier float64) (BattleArea, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleArea{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
