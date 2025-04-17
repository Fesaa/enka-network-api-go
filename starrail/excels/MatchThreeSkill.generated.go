package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MatchThreeSkill struct {
	Desc              map[string]json.Number `json:"Desc"`
	DescFigure        string                 `json:"DescFigure"`
	SkillChargedImg   string                 `json:"SkillChargedImg"`
	SkillID           json.Number            `json:"SkillID"`
	SkillJson         string                 `json:"SkillJson"`
	SkillUnchangedImg string                 `json:"SkillUnchangedImg"`
}
type MatchThreeSkillAccessor struct {
	_data          []MatchThreeSkill
	_dataSkillJson map[string]MatchThreeSkill
}

// LoadData retrieves the data. Must be called before MatchThreeSkill.GroupData
func (a *MatchThreeSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MatchThreeSkill.json")
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
func (a *MatchThreeSkillAccessor) Raw() ([]MatchThreeSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MatchThreeSkill{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MatchThreeSkillAccessor.LoadData to preload everything
func (a *MatchThreeSkillAccessor) GroupData() {
	a._dataSkillJson = map[string]MatchThreeSkill{}
	for _, d := range a._data {
		a._dataSkillJson[d.SkillJson] = d
	}
}

// BySkillJson returns the MatchThreeSkill uniquely identified by SkillJson
//
// Error is only non-nil if the source errors out
func (a *MatchThreeSkillAccessor) BySkillJson(identifier string) (MatchThreeSkill, error) {
	if a._dataSkillJson == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MatchThreeSkill{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSkillJson[identifier], nil
}
