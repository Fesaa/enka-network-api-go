package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type BattleTargetConfig struct {
	AbilityName        string    `json:"AbilityName"`
	HintStep           []float64 `json:"HintStep"`
	ID                 float64   `json:"ID"`
	IconType           string    `json:"IconType"`
	IsFixableHeight    bool      `json:"IsFixableHeight"`
	IsShowProgress     float64   `json:"IsShowProgress"`
	MultiTarget        []float64 `json:"MultiTarget"`
	ParamType          string    `json:"ParamType"`
	ShowInScoreCounter bool      `json:"ShowInScoreCounter"`
	TargetName         hash.Hash `json:"TargetName"`
	TargetNameSimple   hash.Hash `json:"TargetNameSimple"`
	TargetParam        float64   `json:"TargetParam"`
	Type               string    `json:"Type"`
}
type BattleTargetConfigAccessor struct {
	_data   []BattleTargetConfig
	_dataID map[float64]BattleTargetConfig
}

// LoadData retrieves the data. Must be called before BattleTargetConfig.GroupData
func (a *BattleTargetConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/BattleTargetConfig.json")
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
func (a *BattleTargetConfigAccessor) Raw() ([]BattleTargetConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []BattleTargetConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with BattleTargetConfigAccessor.LoadData to preload everything
func (a *BattleTargetConfigAccessor) GroupData() {
	a._dataID = map[float64]BattleTargetConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the BattleTargetConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *BattleTargetConfigAccessor) ByID(identifier float64) (BattleTargetConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return BattleTargetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
