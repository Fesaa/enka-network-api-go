package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesRegionPhase struct {
	BuffID        float64   `json:"BuffID"`
	Description   hash.Hash `json:"Description"`
	EffectDesc    hash.Hash `json:"EffectDesc"`
	Name          hash.Hash `json:"Name"`
	PhaseID       float64   `json:"PhaseID"`
	PicPath       string    `json:"PicPath"`
	ProgressValue float64   `json:"ProgressValue"`
	RewardID      float64   `json:"RewardID"`
}
type PlanetFesRegionPhaseAccessor struct {
	_data              []PlanetFesRegionPhase
	_dataBuffID        map[float64]PlanetFesRegionPhase
	_dataPhaseID       map[float64]PlanetFesRegionPhase
	_dataPicPath       map[string]PlanetFesRegionPhase
	_dataProgressValue map[float64]PlanetFesRegionPhase
}

// LoadData retrieves the data. Must be called before PlanetFesRegionPhase.GroupData
func (a *PlanetFesRegionPhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesRegionPhase.json")
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
func (a *PlanetFesRegionPhaseAccessor) Raw() ([]PlanetFesRegionPhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesRegionPhase{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesRegionPhaseAccessor.LoadData to preload everything
func (a *PlanetFesRegionPhaseAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBuffID[d.BuffID] = d
		a._dataPhaseID[d.PhaseID] = d
		a._dataPicPath[d.PicPath] = d
		a._dataProgressValue[d.ProgressValue] = d
	}
}

// ByBuffID returns the PlanetFesRegionPhase uniquely identified by BuffID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesRegionPhaseAccessor) ByBuffID(identifier float64) (PlanetFesRegionPhase, error) {
	if a._dataBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesRegionPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBuffID[identifier], nil
}

// ByPhaseID returns the PlanetFesRegionPhase uniquely identified by PhaseID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesRegionPhaseAccessor) ByPhaseID(identifier float64) (PlanetFesRegionPhase, error) {
	if a._dataPhaseID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesRegionPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhaseID[identifier], nil
}

// ByPicPath returns the PlanetFesRegionPhase uniquely identified by PicPath
//
// Error is only non-nil if the source errors out
func (a *PlanetFesRegionPhaseAccessor) ByPicPath(identifier string) (PlanetFesRegionPhase, error) {
	if a._dataPicPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesRegionPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPicPath[identifier], nil
}

// ByProgressValue returns the PlanetFesRegionPhase uniquely identified by ProgressValue
//
// Error is only non-nil if the source errors out
func (a *PlanetFesRegionPhaseAccessor) ByProgressValue(identifier float64) (PlanetFesRegionPhase, error) {
	if a._dataProgressValue == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesRegionPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataProgressValue[identifier], nil
}
