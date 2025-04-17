package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type EvolveBuildForgeMaterial struct {
	CostGearList     []float64          `json:"CostGearList"`
	ForgeGearID      float64            `json:"ForgeGearID"`
	MaterialGearList map[string]float64 `json:"MaterialGearList"`
}
type EvolveBuildForgeMaterialAccessor struct {
	_data            []EvolveBuildForgeMaterial
	_dataForgeGearID map[float64]EvolveBuildForgeMaterial
}

// LoadData retrieves the data. Must be called before EvolveBuildForgeMaterial.GroupData
func (a *EvolveBuildForgeMaterialAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildForgeMaterial.json")
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
func (a *EvolveBuildForgeMaterialAccessor) Raw() ([]EvolveBuildForgeMaterial, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildForgeMaterial{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildForgeMaterialAccessor.LoadData to preload everything
func (a *EvolveBuildForgeMaterialAccessor) GroupData() {
	a._dataForgeGearID = map[float64]EvolveBuildForgeMaterial{}
	for _, d := range a._data {
		a._dataForgeGearID[d.ForgeGearID] = d
	}
}

// ByForgeGearID returns the EvolveBuildForgeMaterial uniquely identified by ForgeGearID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildForgeMaterialAccessor) ByForgeGearID(identifier float64) (EvolveBuildForgeMaterial, error) {
	if a._dataForgeGearID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildForgeMaterial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataForgeGearID[identifier], nil
}
