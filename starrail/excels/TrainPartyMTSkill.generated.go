package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type TrainPartyMTSkill struct {
	BFNLPBIEHLH json.Number   `json:"BFNLPBIEHLH"`
	NOCBBALICHO json.Number   `json:"NOCBBALICHO"`
	OIEKCNPMDFC []json.Number `json:"OIEKCNPMDFC"`
	POCJCMKKJIK []json.Number `json:"POCJCMKKJIK"`
}
type TrainPartyMTSkillAccessor struct {
	_data []TrainPartyMTSkill
}

// LoadData retrieves the data. Must be called before TrainPartyMTSkill.GroupData
func (a *TrainPartyMTSkillAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyMTSkill.json")
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
func (a *TrainPartyMTSkillAccessor) Raw() ([]TrainPartyMTSkill, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyMTSkill{}, err
		}
	}
	return a._data, nil
}
