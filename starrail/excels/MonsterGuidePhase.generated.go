package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MonsterGuidePhase struct {
	Difficulty       json.Number            `json:"Difficulty"`
	PhaseAnswer      map[string]json.Number `json:"PhaseAnswer"`
	PhaseDescription map[string]json.Number `json:"PhaseDescription"`
	PhaseID          json.Number            `json:"PhaseID"`
	PhaseName        map[string]json.Number `json:"PhaseName"`
	PhasePic         string                 `json:"PhasePic"`
	SkillList        []json.Number          `json:"SkillList"`
}
type MonsterGuidePhaseAccessor struct {
	_data        []MonsterGuidePhase
	_dataPhaseID map[json.Number]MonsterGuidePhase
}

// LoadData retrieves the data. Must be called before MonsterGuidePhase.GroupData
func (a *MonsterGuidePhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterGuidePhase.json")
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
func (a *MonsterGuidePhaseAccessor) Raw() ([]MonsterGuidePhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterGuidePhase{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonsterGuidePhaseAccessor.LoadData to preload everything
func (a *MonsterGuidePhaseAccessor) GroupData() {
	a._dataPhaseID = map[json.Number]MonsterGuidePhase{}
	for _, d := range a._data {
		a._dataPhaseID[d.PhaseID] = d
	}
}

// ByPhaseID returns the MonsterGuidePhase uniquely identified by PhaseID
//
// Error is only non-nil if the source errors out
func (a *MonsterGuidePhaseAccessor) ByPhaseID(identifier json.Number) (MonsterGuidePhase, error) {
	if a._dataPhaseID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonsterGuidePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhaseID[identifier], nil
}
