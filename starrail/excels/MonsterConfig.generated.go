package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonsterConfig struct {
	AbilityNameList         []string                               `json:"AbilityNameList"`
	AttackModifyRatio       hash.IntValue                          `json:"AttackModifyRatio"`
	CustomValueTags         []string                               `json:"CustomValueTags"`
	CustomValues            []MonsterConfigCustomValues            `json:"CustomValues"`
	DamageTypeResistance    []MonsterConfigDamageTypeResistance    `json:"DamageTypeResistance"`
	DebuffResist            []MonsterConfigDebuffResist            `json:"DebuffResist"`
	DefenceModifyRatio      hash.IntValue                          `json:"DefenceModifyRatio"`
	DynamicValues           []interface{}                          `json:"DynamicValues"`
	EliteGroup              float64                                `json:"EliteGroup"`
	HPModifyRatio           hash.IntValue                          `json:"HPModifyRatio"`
	HardLevelGroup          float64                                `json:"HardLevelGroup"`
	MonsterID               float64                                `json:"MonsterID"`
	MonsterIntroduction     hash.Hash                              `json:"MonsterIntroduction"`
	MonsterName             hash.Hash                              `json:"MonsterName"`
	MonsterTemplateID       float64                                `json:"MonsterTemplateID"`
	OverrideAIPath          string                                 `json:"OverrideAIPath"`
	OverrideAISkillSequence []MonsterConfigOverrideAISkillSequence `json:"OverrideAISkillSequence"`
	OverrideSkillParams     []interface{}                          `json:"OverrideSkillParams"`
	SkillList               []float64                              `json:"SkillList"`
	SpeedModifyRatio        hash.IntValue                          `json:"SpeedModifyRatio"`
	SpeedModifyValue        hash.IntValue                          `json:"SpeedModifyValue"`
	StanceModifyRatio       hash.IntValue                          `json:"StanceModifyRatio"`
	StanceModifyValue       hash.IntValue                          `json:"StanceModifyValue"`
	StanceWeakList          []string                               `json:"StanceWeakList"`
	SummonIDList            []float64                              `json:"SummonIDList"`
}
type MonsterConfigCustomValues struct {
	HEIKKHLKMOA string  `json:"HEIKKHLKMOA"`
	MBMDOCJIMEJ float64 `json:"MBMDOCJIMEJ"`
}
type MonsterConfigDamageTypeResistance struct {
	DamageType string        `json:"DamageType"`
	Value      hash.IntValue `json:"Value"`
}
type MonsterConfigDebuffResist struct {
	Key   string        `json:"Key"`
	Value hash.IntValue `json:"Value"`
}
type MonsterConfigOverrideAISkillSequence struct {
	BMEJEMLFEIO float64 `json:"BMEJEMLFEIO"`
}
type MonsterConfigAccessor struct {
	_data          []MonsterConfig
	_dataMonsterID map[float64]MonsterConfig
}

// LoadData retrieves the data. Must be called before MonsterConfig.GroupData
func (a *MonsterConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterConfig.json")
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
func (a *MonsterConfigAccessor) Raw() ([]MonsterConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonsterConfigAccessor.LoadData to preload everything
func (a *MonsterConfigAccessor) GroupData() {
	a._dataMonsterID = map[float64]MonsterConfig{}
	for _, d := range a._data {
		a._dataMonsterID[d.MonsterID] = d
	}
}

// ByMonsterID returns the MonsterConfig uniquely identified by MonsterID
//
// Error is only non-nil if the source errors out
func (a *MonsterConfigAccessor) ByMonsterID(identifier float64) (MonsterConfig, error) {
	if a._dataMonsterID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonsterConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterID[identifier], nil
}
