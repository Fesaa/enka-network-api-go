package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type AetherDividePassiveSkill struct {
	AbilityName             string                   `json:"AbilityName"`
	ExtraEffectIDList       []interface{}            `json:"ExtraEffectIDList"`
	ItemDescription         string                   `json:"ItemDescription"`
	ItemID                  json.Number              `json:"ItemID"`
	ParamList               []map[string]json.Number `json:"ParamList"`
	PassiveSkillDescription string                   `json:"PassiveSkillDescription"`
	PassiveSkillName        string                   `json:"PassiveSkillName"`
	PassiveSkillType        string                   `json:"PassiveSkillType"`
	Rarity                  json.Number              `json:"Rarity"`
	SimpleExtraEffectIDList []interface{}            `json:"SimpleExtraEffectIDList"`
	SimpleParamList         []map[string]json.Number `json:"SimpleParamList"`
}
type AetherDividePassiveSkillAccessor struct {
	_data                        []AetherDividePassiveSkill
	_dataAbilityName             map[string]AetherDividePassiveSkill
	_dataItemDescription         map[string]AetherDividePassiveSkill
	_dataPassiveSkillDescription map[string]AetherDividePassiveSkill
	_dataPassiveSkillName        map[string]AetherDividePassiveSkill
}

// LoadData retrieves the data. Must be called before AetherDividePassiveSkill.GroupData
func (a *AetherDividePassiveSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDividePassiveSkill.json")
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
func (a *AetherDividePassiveSkillAccessor) Raw() ([]AetherDividePassiveSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDividePassiveSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDividePassiveSkillAccessor.LoadData to preload everything
func (a *AetherDividePassiveSkillAccessor) GroupData() {
	a._dataAbilityName = map[string]AetherDividePassiveSkill{}
	a._dataItemDescription = map[string]AetherDividePassiveSkill{}
	a._dataPassiveSkillDescription = map[string]AetherDividePassiveSkill{}
	a._dataPassiveSkillName = map[string]AetherDividePassiveSkill{}
	for _, d := range a._data {
		a._dataAbilityName[d.AbilityName] = d
		a._dataItemDescription[d.ItemDescription] = d
		a._dataPassiveSkillDescription[d.PassiveSkillDescription] = d
		a._dataPassiveSkillName[d.PassiveSkillName] = d
	}
}

// ByAbilityName returns the AetherDividePassiveSkill uniquely identified by AbilityName
//
// Error is only non-nil if the source errors out
func (a *AetherDividePassiveSkillAccessor) ByAbilityName(identifier string) (AetherDividePassiveSkill, error) {
	if a._dataAbilityName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDividePassiveSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAbilityName[identifier], nil
}

// ByItemDescription returns the AetherDividePassiveSkill uniquely identified by ItemDescription
//
// Error is only non-nil if the source errors out
func (a *AetherDividePassiveSkillAccessor) ByItemDescription(identifier string) (AetherDividePassiveSkill, error) {
	if a._dataItemDescription == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDividePassiveSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemDescription[identifier], nil
}

// ByPassiveSkillDescription returns the AetherDividePassiveSkill uniquely identified by PassiveSkillDescription
//
// Error is only non-nil if the source errors out
func (a *AetherDividePassiveSkillAccessor) ByPassiveSkillDescription(identifier string) (AetherDividePassiveSkill, error) {
	if a._dataPassiveSkillDescription == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDividePassiveSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPassiveSkillDescription[identifier], nil
}

// ByPassiveSkillName returns the AetherDividePassiveSkill uniquely identified by PassiveSkillName
//
// Error is only non-nil if the source errors out
func (a *AetherDividePassiveSkillAccessor) ByPassiveSkillName(identifier string) (AetherDividePassiveSkill, error) {
	if a._dataPassiveSkillName == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDividePassiveSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPassiveSkillName[identifier], nil
}
