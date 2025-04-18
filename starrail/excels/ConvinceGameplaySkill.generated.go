package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ConvinceGameplaySkill struct {
	ID                 float64   `json:"ID"`
	SkillDescriptionID hash.Hash `json:"SkillDescriptionID"`
	SkillIconPath      string    `json:"SkillIconPath"`
	SkillNameText      hash.Hash `json:"SkillNameText"`
}
type ConvinceGameplaySkillAccessor struct {
	_data              []ConvinceGameplaySkill
	_dataID            map[float64]ConvinceGameplaySkill
	_dataSkillIconPath map[string]ConvinceGameplaySkill
}

// LoadData retrieves the data. Must be called before ConvinceGameplaySkill.GroupData
func (a *ConvinceGameplaySkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ConvinceGameplaySkill.json")
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
func (a *ConvinceGameplaySkillAccessor) Raw() ([]ConvinceGameplaySkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ConvinceGameplaySkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ConvinceGameplaySkillAccessor.LoadData to preload everything
func (a *ConvinceGameplaySkillAccessor) GroupData() {
	a._dataID = map[float64]ConvinceGameplaySkill{}
	a._dataSkillIconPath = map[string]ConvinceGameplaySkill{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataSkillIconPath[d.SkillIconPath] = d
	}
}

// ByID returns the ConvinceGameplaySkill uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ConvinceGameplaySkillAccessor) ByID(identifier float64) (ConvinceGameplaySkill, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ConvinceGameplaySkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// BySkillIconPath returns the ConvinceGameplaySkill uniquely identified by SkillIconPath
//
// Error is only non-nil if the source errors out
func (a *ConvinceGameplaySkillAccessor) BySkillIconPath(identifier string) (ConvinceGameplaySkill, error) {
	if a._dataSkillIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ConvinceGameplaySkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillIconPath[identifier], nil
}
