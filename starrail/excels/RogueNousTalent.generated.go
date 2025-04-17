package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueNousTalent struct {
	Cost                []map[string]json.Number `json:"Cost"`
	EffectDesc          map[string]json.Number   `json:"EffectDesc"`
	EffectDescParamList []map[string]json.Number `json:"EffectDescParamList"`
	EffectTag           map[string]json.Number   `json:"EffectTag"`
	EffectTitle         map[string]json.Number   `json:"EffectTitle"`
	Icon                string                   `json:"Icon"`
	NextTalentIDList    []json.Number            `json:"NextTalentIDList"`
	TalentID            json.Number              `json:"TalentID"`
	UnlockIDList        []interface{}            `json:"UnlockIDList"`
}
type RogueNousTalentAccessor struct {
	_data []RogueNousTalent
}

// LoadData retrieves the data. Must be called before RogueNousTalent.GroupData
func (a *RogueNousTalentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousTalent.json")
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
func (a *RogueNousTalentAccessor) Raw() ([]RogueNousTalent, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousTalent{}, err
		}
	}
	return a._data, nil
}
