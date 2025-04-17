package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type EvolveBuildTutorial struct {
	ID            json.Number `json:"ID"`
	StageMergedID json.Number `json:"StageMergedID"`
	TutorialID    string      `json:"TutorialID"`
	WeaponLevel   json.Number `json:"WeaponLevel"`
}
type EvolveBuildTutorialAccessor struct {
	_data           []EvolveBuildTutorial
	_dataTutorialID map[string]EvolveBuildTutorial
}

// LoadData retrieves the data. Must be called before EvolveBuildTutorial.GroupData
func (a *EvolveBuildTutorialAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EvolveBuildTutorial.json")
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
func (a *EvolveBuildTutorialAccessor) Raw() ([]EvolveBuildTutorial, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EvolveBuildTutorial{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EvolveBuildTutorialAccessor.LoadData to preload everything
func (a *EvolveBuildTutorialAccessor) GroupData() {
	a._dataTutorialID = map[string]EvolveBuildTutorial{}
	for _, d := range a._data {
		a._dataTutorialID[d.TutorialID] = d
	}
}

// ByTutorialID returns the EvolveBuildTutorial uniquely identified by TutorialID
//
// Error is only non-nil if the source errors out
func (a *EvolveBuildTutorialAccessor) ByTutorialID(identifier string) (EvolveBuildTutorial, error) {
	if a._dataTutorialID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EvolveBuildTutorial{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTutorialID[identifier], nil
}
