package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type FightFestPhase struct {
	Board3DTexture      string    `json:"Board3DTexture"`
	BoardTitle          hash.Hash `json:"BoardTitle"`
	IconPath1           string    `json:"IconPath1"`
	IconPath2           string    `json:"IconPath2"`
	LukaAnimTrigger     string    `json:"LukaAnimTrigger"`
	MiniIconPath1       string    `json:"MiniIconPath1"`
	PhaseID             float64   `json:"PhaseID"`
	PhaseTitle          hash.Hash `json:"PhaseTitle"`
	PhaseTutorialParams []float64 `json:"PhaseTutorialParams"`
	PhaseType           string    `json:"PhaseType"`
	SortWeight          float64   `json:"SortWeight"`
	TargetTips          hash.Hash `json:"TargetTips"`
	UnlockSubMissionID  float64   `json:"UnlockSubMissionID"`
}
type FightFestPhaseAccessor struct {
	_data                   []FightFestPhase
	_dataPhaseID            map[float64]FightFestPhase
	_dataSortWeight         map[float64]FightFestPhase
	_dataUnlockSubMissionID map[float64]FightFestPhase
}

// LoadData retrieves the data. Must be called before FightFestPhase.GroupData
func (a *FightFestPhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FightFestPhase.json")
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
func (a *FightFestPhaseAccessor) Raw() ([]FightFestPhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FightFestPhase{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FightFestPhaseAccessor.LoadData to preload everything
func (a *FightFestPhaseAccessor) GroupData() {
	a._dataPhaseID = map[float64]FightFestPhase{}
	a._dataSortWeight = map[float64]FightFestPhase{}
	a._dataUnlockSubMissionID = map[float64]FightFestPhase{}
	for _, d := range a._data {
		a._dataPhaseID[d.PhaseID] = d
		a._dataSortWeight[d.SortWeight] = d
		a._dataUnlockSubMissionID[d.UnlockSubMissionID] = d
	}
}

// ByPhaseID returns the FightFestPhase uniquely identified by PhaseID
//
// Error is only non-nil if the source errors out
func (a *FightFestPhaseAccessor) ByPhaseID(identifier float64) (FightFestPhase, error) {
	if a._dataPhaseID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhaseID[identifier], nil
}

// BySortWeight returns the FightFestPhase uniquely identified by SortWeight
//
// Error is only non-nil if the source errors out
func (a *FightFestPhaseAccessor) BySortWeight(identifier float64) (FightFestPhase, error) {
	if a._dataSortWeight == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSortWeight[identifier], nil
}

// ByUnlockSubMissionID returns the FightFestPhase uniquely identified by UnlockSubMissionID
//
// Error is only non-nil if the source errors out
func (a *FightFestPhaseAccessor) ByUnlockSubMissionID(identifier float64) (FightFestPhase, error) {
	if a._dataUnlockSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestPhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockSubMissionID[identifier], nil
}
