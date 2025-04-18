package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type BattleEventButtonTypeConfig struct {
	ButtonPath         string  `json:"ButtonPath"`
	ButtonReadyPath    string  `json:"ButtonReadyPath"`
	CutinPath          string  `json:"CutinPath"`
	ID                 float64 `json:"ID"`
	SkillButtonEffPath string  `json:"SkillButtonEffPath"`
}
type BattleEventButtonTypeConfigAccessor struct {
	_data           []BattleEventButtonTypeConfig
	_dataButtonPath map[string]BattleEventButtonTypeConfig
	_dataID         map[float64]BattleEventButtonTypeConfig
}

// LoadData retrieves the data. Must be called before BattleEventButtonTypeConfig.GroupData
func (a *BattleEventButtonTypeConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleEventButtonTypeConfig.json")
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
func (a *BattleEventButtonTypeConfigAccessor) Raw() ([]BattleEventButtonTypeConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleEventButtonTypeConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleEventButtonTypeConfigAccessor.LoadData to preload everything
func (a *BattleEventButtonTypeConfigAccessor) GroupData() {
	a._dataButtonPath = map[string]BattleEventButtonTypeConfig{}
	a._dataID = map[float64]BattleEventButtonTypeConfig{}
	for _, d := range a._data {
		a._dataButtonPath[d.ButtonPath] = d
		a._dataID[d.ID] = d
	}
}

// ByButtonPath returns the BattleEventButtonTypeConfig uniquely identified by ButtonPath
//
// Error is only non-nil if the source errors out
func (a *BattleEventButtonTypeConfigAccessor) ByButtonPath(identifier string) (BattleEventButtonTypeConfig, error) {
	if a._dataButtonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleEventButtonTypeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataButtonPath[identifier], nil
}

// ByID returns the BattleEventButtonTypeConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BattleEventButtonTypeConfigAccessor) ByID(identifier float64) (BattleEventButtonTypeConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleEventButtonTypeConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
