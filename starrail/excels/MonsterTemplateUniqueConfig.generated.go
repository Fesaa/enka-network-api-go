package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MonsterTemplateUniqueConfig struct {
	AIPath              string                   `json:"AIPath"`
	AISkillSequence     []map[string]json.Number `json:"AISkillSequence"`
	AttackBase          map[string]json.Number   `json:"AttackBase"`
	HPBase              map[string]json.Number   `json:"HPBase"`
	IconPath            string                   `json:"IconPath"`
	ImagePath           string                   `json:"ImagePath"`
	InitialDelayRatio   map[string]json.Number   `json:"InitialDelayRatio"`
	JsonConfig          string                   `json:"JsonConfig"`
	ManikinConfigPath   string                   `json:"ManikinConfigPath"`
	ManikinImagePath    string                   `json:"ManikinImagePath"`
	ManikinPrefabPath   string                   `json:"ManikinPrefabPath"`
	MinimumFatigueRatio map[string]json.Number   `json:"MinimumFatigueRatio"`
	MonsterBaseType     string                   `json:"MonsterBaseType"`
	MonsterName         map[string]json.Number   `json:"MonsterName"`
	MonsterTemplateID   json.Number              `json:"MonsterTemplateID"`
	NPCMonsterList      []interface{}            `json:"NPCMonsterList"`
	NatureID            json.Number              `json:"NatureID"`
	PrefabPath          string                   `json:"PrefabPath"`
	Rank                string                   `json:"Rank"`
	RoundIconPath       string                   `json:"RoundIconPath"`
	SpeedBase           map[string]json.Number   `json:"SpeedBase"`
	StanceBase          map[string]json.Number   `json:"StanceBase"`
	StanceCount         json.Number              `json:"StanceCount"`
}
type MonsterTemplateUniqueConfigAccessor struct {
	_data                  []MonsterTemplateUniqueConfig
	_dataMonsterTemplateID map[json.Number]MonsterTemplateUniqueConfig
}

// LoadData retrieves the data. Must be called before MonsterTemplateUniqueConfig.GroupData
func (a *MonsterTemplateUniqueConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterTemplateUniqueConfig.json")
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
func (a *MonsterTemplateUniqueConfigAccessor) Raw() ([]MonsterTemplateUniqueConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterTemplateUniqueConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonsterTemplateUniqueConfigAccessor.LoadData to preload everything
func (a *MonsterTemplateUniqueConfigAccessor) GroupData() {
	a._dataMonsterTemplateID = map[json.Number]MonsterTemplateUniqueConfig{}
	for _, d := range a._data {
		a._dataMonsterTemplateID[d.MonsterTemplateID] = d
	}
}

// ByMonsterTemplateID returns the MonsterTemplateUniqueConfig uniquely identified by MonsterTemplateID
//
// Error is only non-nil if the source errors out
func (a *MonsterTemplateUniqueConfigAccessor) ByMonsterTemplateID(identifier json.Number) (MonsterTemplateUniqueConfig, error) {
	if a._dataMonsterTemplateID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonsterTemplateUniqueConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMonsterTemplateID[identifier], nil
}
