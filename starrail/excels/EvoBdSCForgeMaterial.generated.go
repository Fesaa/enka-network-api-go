package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type EvoBdSCForgeMaterial struct {
	CostGearList     []float64          `json:"CostGearList"`
	ForgeGearID      float64            `json:"ForgeGearID"`
	MaterialGearList map[string]float64 `json:"MaterialGearList"`
}
type EvoBdSCForgeMaterialAccessor struct {
	_data            []EvoBdSCForgeMaterial
	_dataForgeGearID map[float64]EvoBdSCForgeMaterial
}

// LoadData retrieves the data. Must be called before EvoBdSCForgeMaterial.GroupData
func (a *EvoBdSCForgeMaterialAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvoBdSCForgeMaterial.json")
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
func (a *EvoBdSCForgeMaterialAccessor) Raw() ([]EvoBdSCForgeMaterial, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvoBdSCForgeMaterial{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvoBdSCForgeMaterialAccessor.LoadData to preload everything
func (a *EvoBdSCForgeMaterialAccessor) GroupData() {
	a._dataForgeGearID = map[float64]EvoBdSCForgeMaterial{}
	for _, d := range a._data {
		a._dataForgeGearID[d.ForgeGearID] = d
	}
}

// ByForgeGearID returns the EvoBdSCForgeMaterial uniquely identified by ForgeGearID
//
// Error is only non-nil if the source errors out
func (a *EvoBdSCForgeMaterialAccessor) ByForgeGearID(identifier float64) (EvoBdSCForgeMaterial, error) {
	if a._dataForgeGearID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvoBdSCForgeMaterial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataForgeGearID[identifier], nil
}
