package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BattlePassWeekConfig struct {
	BPLevelExp   float64 `json:"BPLevelExp"`
	BPWeekMaxExp float64 `json:"BPWeekMaxExp"`
	ID           float64 `json:"ID"`
}
type BattlePassWeekConfigAccessor struct {
	_data           []BattlePassWeekConfig
	_dataBPLevelExp map[float64]BattlePassWeekConfig
	_dataID         map[float64]BattlePassWeekConfig
}

// LoadData retrieves the data. Must be called before BattlePassWeekConfig.GroupData
func (a *BattlePassWeekConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattlePassWeekConfig.json")
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
func (a *BattlePassWeekConfigAccessor) Raw() ([]BattlePassWeekConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattlePassWeekConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattlePassWeekConfigAccessor.LoadData to preload everything
func (a *BattlePassWeekConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBPLevelExp[d.BPLevelExp] = d
		a._dataID[d.ID] = d
	}
}

// ByBPLevelExp returns the BattlePassWeekConfig uniquely identified by BPLevelExp
//
// Error is only non-nil if the source errors out
func (a *BattlePassWeekConfigAccessor) ByBPLevelExp(identifier float64) (BattlePassWeekConfig, error) {
	if a._dataBPLevelExp == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattlePassWeekConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBPLevelExp[identifier], nil
}

// ByID returns the BattlePassWeekConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BattlePassWeekConfigAccessor) ByID(identifier float64) (BattlePassWeekConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattlePassWeekConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
