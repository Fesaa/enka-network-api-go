package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type EvolveBuildStageConfig struct {
	BuffTextFormat      map[string]json.Number           `json:"BuffTextFormat"`
	Difficulty          json.Number                      `json:"Difficulty"`
	FirstWinReward      []json.Number                    `json:"FirstWinReward"`
	GearRecommendList   []json.Number                    `json:"GearRecommendList"`
	InitialWeapon       []json.Number                    `json:"InitialWeapon"`
	IntroID             json.Number                      `json:"IntroID"`
	Name                map[string]json.Number           `json:"Name"`
	RankList            []EvolveBuildStageConfigRankList `json:"RankList"`
	RecommendList       []map[string]json.Number         `json:"RecommendList"`
	Season              string                           `json:"Season"`
	StageMergedID       json.Number                      `json:"StageMergedID"`
	StagePeriod1        []json.Number                    `json:"StagePeriod1"`
	StagePeriod2        []json.Number                    `json:"StagePeriod2"`
	StagePeriod3        []json.Number                    `json:"StagePeriod3"`
	StagePeriod4        []json.Number                    `json:"StagePeriod4"`
	TeamBonusIconPath   string                           `json:"TeamBonusIconPath"`
	TeamBonusMazeBuffID json.Number                      `json:"TeamBonusMazeBuffID"`
	TeamBonusShortDesc  map[string]json.Number           `json:"TeamBonusShortDesc"`
	TrialAvatar         []json.Number                    `json:"TrialAvatar"`
	UnlockQuest         json.Number                      `json:"UnlockQuest"`
	WeaponSelectable    bool                             `json:"WeaponSelectable"`
}
type EvolveBuildStageConfigRankList struct {
	DNPHCJEBIKB json.Number `json:"DNPHCJEBIKB"`
	MJJBNDAKOIF string      `json:"MJJBNDAKOIF"`
}
type EvolveBuildStageConfigAccessor struct {
	_data                    []EvolveBuildStageConfig
	_dataStageMergedID       map[json.Number]EvolveBuildStageConfig
	_dataTeamBonusMazeBuffID map[json.Number]EvolveBuildStageConfig
}

// LoadData retrieves the data. Must be called before EvolveBuildStageConfig.GroupData
func (a *EvolveBuildStageConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildStageConfig.json")
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
func (a *EvolveBuildStageConfigAccessor) Raw() ([]EvolveBuildStageConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildStageConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildStageConfigAccessor.LoadData to preload everything
func (a *EvolveBuildStageConfigAccessor) GroupData() {
	a._dataStageMergedID = map[json.Number]EvolveBuildStageConfig{}
	a._dataTeamBonusMazeBuffID = map[json.Number]EvolveBuildStageConfig{}
	for _, d := range a._data {
		a._dataStageMergedID[d.StageMergedID] = d
		a._dataTeamBonusMazeBuffID[d.TeamBonusMazeBuffID] = d
	}
}

// ByStageMergedID returns the EvolveBuildStageConfig uniquely identified by StageMergedID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildStageConfigAccessor) ByStageMergedID(identifier json.Number) (EvolveBuildStageConfig, error) {
	if a._dataStageMergedID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildStageConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStageMergedID[identifier], nil
}

// ByTeamBonusMazeBuffID returns the EvolveBuildStageConfig uniquely identified by TeamBonusMazeBuffID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildStageConfigAccessor) ByTeamBonusMazeBuffID(identifier json.Number) (EvolveBuildStageConfig, error) {
	if a._dataTeamBonusMazeBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildStageConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTeamBonusMazeBuffID[identifier], nil
}
