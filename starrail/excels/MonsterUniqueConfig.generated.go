package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonsterUniqueConfig struct {
	AbilityNameList         []interface{}                     `json:"AbilityNameList"`
	AttackModifyRatio       hash.IntValue                     `json:"AttackModifyRatio"`
	CustomValueTags         []string                          `json:"CustomValueTags"`
	CustomValues            []MonsterUniqueConfigCustomValues `json:"CustomValues"`
	DamageTypeResistance    []interface{}                     `json:"DamageTypeResistance"`
	DebuffResist            []MonsterUniqueConfigDebuffResist `json:"DebuffResist"`
	DefenceModifyRatio      hash.IntValue                     `json:"DefenceModifyRatio"`
	DynamicValues           []interface{}                     `json:"DynamicValues"`
	EliteGroup              float64                           `json:"EliteGroup"`
	HPModifyRatio           hash.IntValue                     `json:"HPModifyRatio"`
	HardLevelGroup          float64                           `json:"HardLevelGroup"`
	MonsterID               float64                           `json:"MonsterID"`
	MonsterIntroduction     hash.Hash                         `json:"MonsterIntroduction"`
	MonsterName             hash.Hash                         `json:"MonsterName"`
	MonsterTemplateID       float64                           `json:"MonsterTemplateID"`
	OverrideAIPath          string                            `json:"OverrideAIPath"`
	OverrideAISkillSequence []interface{}                     `json:"OverrideAISkillSequence"`
	OverrideSkillParams     []interface{}                     `json:"OverrideSkillParams"`
	SkillList               []float64                         `json:"SkillList"`
	SpeedModifyRatio        hash.IntValue                     `json:"SpeedModifyRatio"`
	SpeedModifyValue        hash.IntValue                     `json:"SpeedModifyValue"`
	StanceModifyRatio       hash.IntValue                     `json:"StanceModifyRatio"`
	StanceModifyValue       hash.IntValue                     `json:"StanceModifyValue"`
	StanceWeakList          []string                          `json:"StanceWeakList"`
	SummonIDList            []interface{}                     `json:"SummonIDList"`
}
type MonsterUniqueConfigCustomValues struct {
	HEIKKHLKMOA string  `json:"HEIKKHLKMOA"`
	MBMDOCJIMEJ float64 `json:"MBMDOCJIMEJ"`
}
type MonsterUniqueConfigDebuffResist struct {
	Key   string        `json:"Key"`
	Value hash.IntValue `json:"Value"`
}
type MonsterUniqueConfigAccessor struct {
	_data          []MonsterUniqueConfig
	_dataMonsterID map[float64]MonsterUniqueConfig
}

// LoadData retrieves the data. Must be called before MonsterUniqueConfig.GroupData
func (a *MonsterUniqueConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterUniqueConfig.json")
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
func (a *MonsterUniqueConfigAccessor) Raw() ([]MonsterUniqueConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterUniqueConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonsterUniqueConfigAccessor.LoadData to preload everything
func (a *MonsterUniqueConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMonsterID[d.MonsterID] = d
	}
}

// ByMonsterID returns the MonsterUniqueConfig uniquely identified by MonsterID
//
// Error is only non-nil if the source errors out
func (a *MonsterUniqueConfigAccessor) ByMonsterID(identifier float64) (MonsterUniqueConfig, error) {
	if a._dataMonsterID == nil {
		err := a.LoadData()
		if err != nil {
			return MonsterUniqueConfig{}, err
		}
		a.GroupData()
	}
	return a._dataMonsterID[identifier], nil
}
