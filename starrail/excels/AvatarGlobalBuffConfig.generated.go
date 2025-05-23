package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AvatarGlobalBuffConfig struct {
	AvatarID                   float64       `json:"AvatarID"`
	Desc                       hash.Hash     `json:"Desc"`
	ExtraEffectIDList          []float64     `json:"ExtraEffectIDList"`
	GameModeBlackList          []float64     `json:"GameModeBlackList"`
	MazeBuffID                 float64       `json:"MazeBuffID"`
	Name                       hash.Hash     `json:"Name"`
	ParamList                  []hash.Value  `json:"ParamList"`
	SimpleDesc                 hash.Hash     `json:"SimpleDesc"`
	SimpleExtraEffectIDList    []interface{} `json:"SimpleExtraEffectIDList"`
	SimpleParamList            []interface{} `json:"SimpleParamList"`
	SkillID                    float64       `json:"SkillID"`
	SkillTag                   hash.Hash     `json:"SkillTag"`
	StageTypeBlackList         []float64     `json:"StageTypeBlackList"`
	TeamBlackList              []float64     `json:"TeamBlackList"`
	TeamStageTypeBlackList     []float64     `json:"TeamStageTypeBlackList"`
	TrialBagStageTypeWhiteList []float64     `json:"TrialBagStageTypeWhiteList"`
}
type AvatarGlobalBuffConfigAccessor struct {
	_data           []AvatarGlobalBuffConfig
	_dataAvatarID   map[float64]AvatarGlobalBuffConfig
	_dataMazeBuffID map[float64]AvatarGlobalBuffConfig
	_dataSkillID    map[float64]AvatarGlobalBuffConfig
}

// LoadData retrieves the data. Must be called before AvatarGlobalBuffConfig.GroupData
func (a *AvatarGlobalBuffConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarGlobalBuffConfig.json")
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
func (a *AvatarGlobalBuffConfigAccessor) Raw() ([]AvatarGlobalBuffConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarGlobalBuffConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarGlobalBuffConfigAccessor.LoadData to preload everything
func (a *AvatarGlobalBuffConfigAccessor) GroupData() {
	a._dataAvatarID = map[float64]AvatarGlobalBuffConfig{}
	a._dataMazeBuffID = map[float64]AvatarGlobalBuffConfig{}
	a._dataSkillID = map[float64]AvatarGlobalBuffConfig{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
		a._dataMazeBuffID[d.MazeBuffID] = d
		a._dataSkillID[d.SkillID] = d
	}
}

// ByAvatarID returns the AvatarGlobalBuffConfig uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *AvatarGlobalBuffConfigAccessor) ByAvatarID(identifier float64) (AvatarGlobalBuffConfig, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarGlobalBuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}

// ByMazeBuffID returns the AvatarGlobalBuffConfig uniquely identified by MazeBuffID
//
// Error is only non-nil if the source errors out
func (a *AvatarGlobalBuffConfigAccessor) ByMazeBuffID(identifier float64) (AvatarGlobalBuffConfig, error) {
	if a._dataMazeBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarGlobalBuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMazeBuffID[identifier], nil
}

// BySkillID returns the AvatarGlobalBuffConfig uniquely identified by SkillID
//
// Error is only non-nil if the source errors out
func (a *AvatarGlobalBuffConfigAccessor) BySkillID(identifier float64) (AvatarGlobalBuffConfig, error) {
	if a._dataSkillID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarGlobalBuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillID[identifier], nil
}
