package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type ChimeraPhase struct {
	LeaderChariotState json.Number   `json:"LeaderChariotState"`
	NextPhaseID        json.Number   `json:"NextPhaseID"`
	PhaseID            json.Number   `json:"PhaseID"`
	RoundList          []json.Number `json:"RoundList"`
	TargetParam        []json.Number `json:"TargetParam"`
	TargetType         string        `json:"TargetType"`
}
type ChimeraPhaseAccessor struct {
	_data []ChimeraPhase
}

// LoadData retrieves the data. Must be called before ChimeraPhase.GroupData
func (a *ChimeraPhaseAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChimeraPhase.json")
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
func (a *ChimeraPhaseAccessor) Raw() ([]ChimeraPhase, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChimeraPhase{}, err
		}
	}
	return a._data, nil
}
