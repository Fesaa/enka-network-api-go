package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MuseumPhase struct {
	MuseumPhaseID   float64   `json:"MuseumPhaseID"`
	PhaseFund       float64   `json:"PhaseFund"`
	PhaseIconPath   string    `json:"PhaseIconPath"`
	PhaseName       hash.Hash `json:"PhaseName"`
	PhaseQuestID    float64   `json:"PhaseQuestID"`
	PhaseTextID     hash.Hash `json:"PhaseTextID"`
	RenewPointCost  float64   `json:"RenewPointCost"`
	UnlockAreaID    float64   `json:"UnlockAreaID"`
	UnlockMissionID float64   `json:"UnlockMissionID"`
}
type MuseumPhaseAccessor struct {
	_data               []MuseumPhase
	_dataMuseumPhaseID  map[float64]MuseumPhase
	_dataPhaseIconPath  map[string]MuseumPhase
	_dataPhaseQuestID   map[float64]MuseumPhase
	_dataRenewPointCost map[float64]MuseumPhase
}

// LoadData retrieves the data. Must be called before MuseumPhase.GroupData
func (a *MuseumPhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MuseumPhase.json")
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
func (a *MuseumPhaseAccessor) Raw() ([]MuseumPhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MuseumPhase{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MuseumPhaseAccessor.LoadData to preload everything
func (a *MuseumPhaseAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMuseumPhaseID[d.MuseumPhaseID] = d
		a._dataPhaseIconPath[d.PhaseIconPath] = d
		a._dataPhaseQuestID[d.PhaseQuestID] = d
		a._dataRenewPointCost[d.RenewPointCost] = d
	}
}

// ByMuseumPhaseID returns the MuseumPhase uniquely identified by MuseumPhaseID
//
// Error is only non-nil if the source errors out
func (a *MuseumPhaseAccessor) ByMuseumPhaseID(identifier float64) (MuseumPhase, error) {
	if a._dataMuseumPhaseID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMuseumPhaseID[identifier], nil
}

// ByPhaseIconPath returns the MuseumPhase uniquely identified by PhaseIconPath
//
// Error is only non-nil if the source errors out
func (a *MuseumPhaseAccessor) ByPhaseIconPath(identifier string) (MuseumPhase, error) {
	if a._dataPhaseIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhaseIconPath[identifier], nil
}

// ByPhaseQuestID returns the MuseumPhase uniquely identified by PhaseQuestID
//
// Error is only non-nil if the source errors out
func (a *MuseumPhaseAccessor) ByPhaseQuestID(identifier float64) (MuseumPhase, error) {
	if a._dataPhaseQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhaseQuestID[identifier], nil
}

// ByRenewPointCost returns the MuseumPhase uniquely identified by RenewPointCost
//
// Error is only non-nil if the source errors out
func (a *MuseumPhaseAccessor) ByRenewPointCost(identifier float64) (MuseumPhase, error) {
	if a._dataRenewPointCost == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MuseumPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRenewPointCost[identifier], nil
}
