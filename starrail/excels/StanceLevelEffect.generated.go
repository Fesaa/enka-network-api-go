package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type StanceLevelEffect struct {
	ID                float64       `json:"ID"`
	LevelDifference   float64       `json:"LevelDifference"`
	StanceLevelEffect hash.IntValue `json:"StanceLevelEffect"`
}
type StanceLevelEffectAccessor struct {
	_data   []StanceLevelEffect
	_dataID map[float64]StanceLevelEffect
}

// LoadData retrieves the data. Must be called before StanceLevelEffect.GroupData
func (a *StanceLevelEffectAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StanceLevelEffect.json")
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
func (a *StanceLevelEffectAccessor) Raw() ([]StanceLevelEffect, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StanceLevelEffect{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StanceLevelEffectAccessor.LoadData to preload everything
func (a *StanceLevelEffectAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the StanceLevelEffect uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *StanceLevelEffectAccessor) ByID(identifier float64) (StanceLevelEffect, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return StanceLevelEffect{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
