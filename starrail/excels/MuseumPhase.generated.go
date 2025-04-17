package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MuseumPhase struct {
	MuseumPhaseID   json.Number            `json:"MuseumPhaseID"`
	PhaseFund       json.Number            `json:"PhaseFund"`
	PhaseIconPath   string                 `json:"PhaseIconPath"`
	PhaseName       map[string]json.Number `json:"PhaseName"`
	PhaseQuestID    json.Number            `json:"PhaseQuestID"`
	PhaseTextID     map[string]json.Number `json:"PhaseTextID"`
	RenewPointCost  json.Number            `json:"RenewPointCost"`
	UnlockAreaID    json.Number            `json:"UnlockAreaID"`
	UnlockMissionID json.Number            `json:"UnlockMissionID"`
}
type MuseumPhaseAccessor struct {
	_data              []MuseumPhase
	_dataPhaseIconPath map[string]MuseumPhase
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
	a._dataPhaseIconPath = map[string]MuseumPhase{}
	for _, d := range a._data {
		a._dataPhaseIconPath[d.PhaseIconPath] = d
	}
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
