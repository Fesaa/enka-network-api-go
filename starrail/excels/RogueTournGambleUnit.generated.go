package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueTournGambleUnit struct {
	GambleUnitID    json.Number `json:"GambleUnitID"`
	GambleUnitIcon  string      `json:"GambleUnitIcon"`
	GambleUnitParam json.Number `json:"GambleUnitParam"`
	GambleUnitType  string      `json:"GambleUnitType"`
}
type RogueTournGambleUnitAccessor struct {
	_data                []RogueTournGambleUnit
	_dataGambleUnitID    map[json.Number]RogueTournGambleUnit
	_dataGambleUnitParam map[json.Number]RogueTournGambleUnit
}

// LoadData retrieves the data. Must be called before RogueTournGambleUnit.GroupData
func (a *RogueTournGambleUnitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournGambleUnit.json")
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
func (a *RogueTournGambleUnitAccessor) Raw() ([]RogueTournGambleUnit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournGambleUnit{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournGambleUnitAccessor.LoadData to preload everything
func (a *RogueTournGambleUnitAccessor) GroupData() {
	a._dataGambleUnitID = map[json.Number]RogueTournGambleUnit{}
	a._dataGambleUnitParam = map[json.Number]RogueTournGambleUnit{}
	for _, d := range a._data {
		a._dataGambleUnitID[d.GambleUnitID] = d
		a._dataGambleUnitParam[d.GambleUnitParam] = d
	}
}

// ByGambleUnitID returns the RogueTournGambleUnit uniquely identified by GambleUnitID
//
// Error is only non-nil if the source errors out
func (a *RogueTournGambleUnitAccessor) ByGambleUnitID(identifier json.Number) (RogueTournGambleUnit, error) {
	if a._dataGambleUnitID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournGambleUnit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGambleUnitID[identifier], nil
}

// ByGambleUnitParam returns the RogueTournGambleUnit uniquely identified by GambleUnitParam
//
// Error is only non-nil if the source errors out
func (a *RogueTournGambleUnitAccessor) ByGambleUnitParam(identifier json.Number) (RogueTournGambleUnit, error) {
	if a._dataGambleUnitParam == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournGambleUnit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGambleUnitParam[identifier], nil
}
