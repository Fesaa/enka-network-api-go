package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type FightFestScorePhase struct {
	AvatarInfoID         json.Number            `json:"AvatarInfoID"`
	PhaseID              json.Number            `json:"PhaseID"`
	RewardID             json.Number            `json:"RewardID"`
	TargetAvatarIcon     string                 `json:"TargetAvatarIcon"`
	TargetAvatarMiniIcon string                 `json:"TargetAvatarMiniIcon"`
	TargetAvatarName     map[string]json.Number `json:"TargetAvatarName"`
	TargetScore          json.Number            `json:"TargetScore"`
	TargetTip            map[string]json.Number `json:"TargetTip"`
}
type FightFestScorePhaseAccessor struct {
	_data                     []FightFestScorePhase
	_dataAvatarInfoID         map[json.Number]FightFestScorePhase
	_dataPhaseID              map[json.Number]FightFestScorePhase
	_dataRewardID             map[json.Number]FightFestScorePhase
	_dataTargetAvatarIcon     map[string]FightFestScorePhase
	_dataTargetAvatarMiniIcon map[string]FightFestScorePhase
	_dataTargetScore          map[json.Number]FightFestScorePhase
}

// LoadData retrieves the data. Must be called before FightFestScorePhase.GroupData
func (a *FightFestScorePhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FightFestScorePhase.json")
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
func (a *FightFestScorePhaseAccessor) Raw() ([]FightFestScorePhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FightFestScorePhase{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FightFestScorePhaseAccessor.LoadData to preload everything
func (a *FightFestScorePhaseAccessor) GroupData() {
	a._dataAvatarInfoID = map[json.Number]FightFestScorePhase{}
	a._dataPhaseID = map[json.Number]FightFestScorePhase{}
	a._dataRewardID = map[json.Number]FightFestScorePhase{}
	a._dataTargetAvatarIcon = map[string]FightFestScorePhase{}
	a._dataTargetAvatarMiniIcon = map[string]FightFestScorePhase{}
	a._dataTargetScore = map[json.Number]FightFestScorePhase{}
	for _, d := range a._data {
		a._dataAvatarInfoID[d.AvatarInfoID] = d
		a._dataPhaseID[d.PhaseID] = d
		a._dataRewardID[d.RewardID] = d
		a._dataTargetAvatarIcon[d.TargetAvatarIcon] = d
		a._dataTargetAvatarMiniIcon[d.TargetAvatarMiniIcon] = d
		a._dataTargetScore[d.TargetScore] = d
	}
}

// ByAvatarInfoID returns the FightFestScorePhase uniquely identified by AvatarInfoID
//
// Error is only non-nil if the source errors out
func (a *FightFestScorePhaseAccessor) ByAvatarInfoID(identifier json.Number) (FightFestScorePhase, error) {
	if a._dataAvatarInfoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestScorePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarInfoID[identifier], nil
}

// ByPhaseID returns the FightFestScorePhase uniquely identified by PhaseID
//
// Error is only non-nil if the source errors out
func (a *FightFestScorePhaseAccessor) ByPhaseID(identifier json.Number) (FightFestScorePhase, error) {
	if a._dataPhaseID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestScorePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPhaseID[identifier], nil
}

// ByRewardID returns the FightFestScorePhase uniquely identified by RewardID
//
// Error is only non-nil if the source errors out
func (a *FightFestScorePhaseAccessor) ByRewardID(identifier json.Number) (FightFestScorePhase, error) {
	if a._dataRewardID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestScorePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardID[identifier], nil
}

// ByTargetAvatarIcon returns the FightFestScorePhase uniquely identified by TargetAvatarIcon
//
// Error is only non-nil if the source errors out
func (a *FightFestScorePhaseAccessor) ByTargetAvatarIcon(identifier string) (FightFestScorePhase, error) {
	if a._dataTargetAvatarIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestScorePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTargetAvatarIcon[identifier], nil
}

// ByTargetAvatarMiniIcon returns the FightFestScorePhase uniquely identified by TargetAvatarMiniIcon
//
// Error is only non-nil if the source errors out
func (a *FightFestScorePhaseAccessor) ByTargetAvatarMiniIcon(identifier string) (FightFestScorePhase, error) {
	if a._dataTargetAvatarMiniIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestScorePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTargetAvatarMiniIcon[identifier], nil
}

// ByTargetScore returns the FightFestScorePhase uniquely identified by TargetScore
//
// Error is only non-nil if the source errors out
func (a *FightFestScorePhaseAccessor) ByTargetScore(identifier json.Number) (FightFestScorePhase, error) {
	if a._dataTargetScore == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FightFestScorePhase{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTargetScore[identifier], nil
}
