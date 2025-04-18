package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PetConfig struct {
	PetID        float64 `json:"PetID"`
	PetItemID    float64 `json:"PetItemID"`
	SummonUnitID float64 `json:"SummonUnitID"`
}
type PetConfigAccessor struct {
	_data             []PetConfig
	_dataPetID        map[float64]PetConfig
	_dataPetItemID    map[float64]PetConfig
	_dataSummonUnitID map[float64]PetConfig
}

// LoadData retrieves the data. Must be called before PetConfig.GroupData
func (a *PetConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PetConfig.json")
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
func (a *PetConfigAccessor) Raw() ([]PetConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PetConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PetConfigAccessor.LoadData to preload everything
func (a *PetConfigAccessor) GroupData() {
	a._dataPetID = map[float64]PetConfig{}
	a._dataPetItemID = map[float64]PetConfig{}
	a._dataSummonUnitID = map[float64]PetConfig{}
	for _, d := range a._data {
		a._dataPetID[d.PetID] = d
		a._dataPetItemID[d.PetItemID] = d
		a._dataSummonUnitID[d.SummonUnitID] = d
	}
}

// ByPetID returns the PetConfig uniquely identified by PetID
//
// Error is only non-nil if the source errors out
func (a *PetConfigAccessor) ByPetID(identifier float64) (PetConfig, error) {
	if a._dataPetID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPetID[identifier], nil
}

// ByPetItemID returns the PetConfig uniquely identified by PetItemID
//
// Error is only non-nil if the source errors out
func (a *PetConfigAccessor) ByPetItemID(identifier float64) (PetConfig, error) {
	if a._dataPetItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPetItemID[identifier], nil
}

// BySummonUnitID returns the PetConfig uniquely identified by SummonUnitID
//
// Error is only non-nil if the source errors out
func (a *PetConfigAccessor) BySummonUnitID(identifier float64) (PetConfig, error) {
	if a._dataSummonUnitID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PetConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSummonUnitID[identifier], nil
}
