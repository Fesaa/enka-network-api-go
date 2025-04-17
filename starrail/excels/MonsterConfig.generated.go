package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MonsterConfig struct {
	AbilityNameList         []string                            `json:"AbilityNameList"`
	AttackModifyRatio       map[string]json.Number              `json:"AttackModifyRatio"`
	CustomValueTags         []string                            `json:"CustomValueTags"`
	CustomValues            []MonsterConfigCustomValues         `json:"CustomValues"`
	DamageTypeResistance    []MonsterConfigDamageTypeResistance `json:"DamageTypeResistance"`
	DebuffResist            []MonsterConfigDebuffResist         `json:"DebuffResist"`
	DefenceModifyRatio      map[string]json.Number              `json:"DefenceModifyRatio"`
	DynamicValues           []interface{}                       `json:"DynamicValues"`
	EliteGroup              json.Number                         `json:"EliteGroup"`
	HPModifyRatio           map[string]json.Number              `json:"HPModifyRatio"`
	HardLevelGroup          json.Number                         `json:"HardLevelGroup"`
	MonsterID               json.Number                         `json:"MonsterID"`
	MonsterIntroduction     map[string]json.Number              `json:"MonsterIntroduction"`
	MonsterName             map[string]json.Number              `json:"MonsterName"`
	MonsterTemplateID       json.Number                         `json:"MonsterTemplateID"`
	OverrideAIPath          string                              `json:"OverrideAIPath"`
	OverrideAISkillSequence []map[string]json.Number            `json:"OverrideAISkillSequence"`
	OverrideSkillParams     []interface{}                       `json:"OverrideSkillParams"`
	SkillList               []json.Number                       `json:"SkillList"`
	SpeedModifyRatio        map[string]json.Number              `json:"SpeedModifyRatio"`
	SpeedModifyValue        map[string]json.Number              `json:"SpeedModifyValue"`
	StanceModifyRatio       map[string]json.Number              `json:"StanceModifyRatio"`
	StanceModifyValue       map[string]json.Number              `json:"StanceModifyValue"`
	StanceWeakList          []string                            `json:"StanceWeakList"`
	SummonIDList            []json.Number                       `json:"SummonIDList"`
}
type MonsterConfigCustomValues struct {
	HEIKKHLKMOA string      `json:"HEIKKHLKMOA"`
	MBMDOCJIMEJ json.Number `json:"MBMDOCJIMEJ"`
}
type MonsterConfigDamageTypeResistance struct {
	DamageType string                 `json:"DamageType"`
	Value      map[string]json.Number `json:"Value"`
}
type MonsterConfigDebuffResist struct {
	Key   string                 `json:"Key"`
	Value map[string]json.Number `json:"Value"`
}
type MonsterConfigAccessor struct {
	_data          []MonsterConfig
	_dataMonsterID map[json.Number]MonsterConfig
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
	a._dataMonsterID = map[json.Number]MonsterConfig{}
	for _, d := range a._data {
		a._dataMonsterID[d.MonsterID] = d
	}
}

// ByMonsterID returns the MonsterConfig uniquely identified by MonsterID
//
// Error is only non-nil if the source errors out
func (a *MonsterConfigAccessor) ByMonsterID(identifier json.Number) (MonsterConfig, error) {
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
