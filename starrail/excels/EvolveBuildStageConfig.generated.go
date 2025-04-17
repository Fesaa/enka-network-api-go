package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EvolveBuildStageConfig struct {
	BuffTextFormat      hash.Hash                             `json:"BuffTextFormat"`
	Difficulty          float64                               `json:"Difficulty"`
	FirstWinReward      []float64                             `json:"FirstWinReward"`
	GearRecommendList   []float64                             `json:"GearRecommendList"`
	InitialWeapon       []float64                             `json:"InitialWeapon"`
	IntroID             float64                               `json:"IntroID"`
	Name                hash.Hash                             `json:"Name"`
	RankList            []EvolveBuildStageConfigRankList      `json:"RankList"`
	RecommendList       []EvolveBuildStageConfigRecommendList `json:"RecommendList"`
	Season              string                                `json:"Season"`
	StageMergedID       float64                               `json:"StageMergedID"`
	StagePeriod1        []float64                             `json:"StagePeriod1"`
	StagePeriod2        []float64                             `json:"StagePeriod2"`
	StagePeriod3        []float64                             `json:"StagePeriod3"`
	StagePeriod4        []float64                             `json:"StagePeriod4"`
	TeamBonusIconPath   string                                `json:"TeamBonusIconPath"`
	TeamBonusMazeBuffID float64                               `json:"TeamBonusMazeBuffID"`
	TeamBonusShortDesc  hash.Hash                             `json:"TeamBonusShortDesc"`
	TrialAvatar         []float64                             `json:"TrialAvatar"`
	UnlockQuest         float64                               `json:"UnlockQuest"`
	WeaponSelectable    bool                                  `json:"WeaponSelectable"`
}
type EvolveBuildStageConfigRankList struct {
	DNPHCJEBIKB float64 `json:"DNPHCJEBIKB"`
	MJJBNDAKOIF string  `json:"MJJBNDAKOIF"`
}
type EvolveBuildStageConfigRecommendList struct {
	BDJECMNLHON float64 `json:"BDJECMNLHON"`
	NOCBBALICHO float64 `json:"NOCBBALICHO"`
}
type EvolveBuildStageConfigAccessor struct {
	_data                    []EvolveBuildStageConfig
	_dataStageMergedID       map[float64]EvolveBuildStageConfig
	_dataTeamBonusMazeBuffID map[float64]EvolveBuildStageConfig
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
	for _, d := range a._data {
		a._dataStageMergedID[d.StageMergedID] = d
		a._dataTeamBonusMazeBuffID[d.TeamBonusMazeBuffID] = d
	}
}

// ByStageMergedID returns the EvolveBuildStageConfig uniquely identified by StageMergedID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildStageConfigAccessor) ByStageMergedID(identifier float64) (EvolveBuildStageConfig, error) {
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
func (a *EvolveBuildStageConfigAccessor) ByTeamBonusMazeBuffID(identifier float64) (EvolveBuildStageConfig, error) {
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
