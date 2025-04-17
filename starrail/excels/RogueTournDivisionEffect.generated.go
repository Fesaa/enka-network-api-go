package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueTournDivisionEffect struct {
	DescParamList []map[string]json.Number `json:"DescParamList"`
	DescText      map[string]json.Number   `json:"DescText"`
	DivisionLevel json.Number              `json:"DivisionLevel"`
}
type RogueTournDivisionEffectAccessor struct {
	_data              []RogueTournDivisionEffect
	_dataDivisionLevel map[json.Number]RogueTournDivisionEffect
}

// LoadData retrieves the data. Must be called before RogueTournDivisionEffect.GroupData
func (a *RogueTournDivisionEffectAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournDivisionEffect.json")
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
func (a *RogueTournDivisionEffectAccessor) Raw() ([]RogueTournDivisionEffect, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournDivisionEffect{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournDivisionEffectAccessor.LoadData to preload everything
func (a *RogueTournDivisionEffectAccessor) GroupData() {
	a._dataDivisionLevel = map[json.Number]RogueTournDivisionEffect{}
	for _, d := range a._data {
		a._dataDivisionLevel[d.DivisionLevel] = d
	}
}

// ByDivisionLevel returns the RogueTournDivisionEffect uniquely identified by DivisionLevel
//
// Error is only non-nil if the source errors out
func (a *RogueTournDivisionEffectAccessor) ByDivisionLevel(identifier json.Number) (RogueTournDivisionEffect, error) {
	if a._dataDivisionLevel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournDivisionEffect{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDivisionLevel[identifier], nil
}
