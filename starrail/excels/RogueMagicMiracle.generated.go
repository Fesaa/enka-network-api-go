package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicMiracle struct {
	MiracleDisplayID        float64 `json:"MiracleDisplayID"`
	MiracleEffectDisplayID  float64 `json:"MiracleEffectDisplayID"`
	MiracleID               float64 `json:"MiracleID"`
	UnlockHandbookMiracleID float64 `json:"UnlockHandbookMiracleID"`
}
type RogueMagicMiracleAccessor struct {
	_data                       []RogueMagicMiracle
	_dataMiracleEffectDisplayID map[float64]RogueMagicMiracle
	_dataMiracleID              map[float64]RogueMagicMiracle
}

// LoadData retrieves the data. Must be called before RogueMagicMiracle.GroupData
func (a *RogueMagicMiracleAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicMiracle.json")
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
func (a *RogueMagicMiracleAccessor) Raw() ([]RogueMagicMiracle, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicMiracle{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicMiracleAccessor.LoadData to preload everything
func (a *RogueMagicMiracleAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMiracleEffectDisplayID[d.MiracleEffectDisplayID] = d
		a._dataMiracleID[d.MiracleID] = d
	}
}

// ByMiracleEffectDisplayID returns the RogueMagicMiracle uniquely identified by MiracleEffectDisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicMiracleAccessor) ByMiracleEffectDisplayID(identifier float64) (RogueMagicMiracle, error) {
	if a._dataMiracleEffectDisplayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicMiracle{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiracleEffectDisplayID[identifier], nil
}

// ByMiracleID returns the RogueMagicMiracle uniquely identified by MiracleID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicMiracleAccessor) ByMiracleID(identifier float64) (RogueMagicMiracle, error) {
	if a._dataMiracleID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicMiracle{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiracleID[identifier], nil
}
