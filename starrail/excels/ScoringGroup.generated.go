package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ScoringGroup struct {
	DisplayType    string                 `json:"DisplayType"`
	ScoreName      map[string]json.Number `json:"ScoreName"`
	ScoringGroupID json.Number            `json:"ScoringGroupID"`
	ScoringIDList  []json.Number          `json:"ScoringIDList"`
}
type ScoringGroupAccessor struct {
	_data               []ScoringGroup
	_dataScoringGroupID map[json.Number]ScoringGroup
}

// LoadData retrieves the data. Must be called before ScoringGroup.GroupData
func (a *ScoringGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScoringGroup.json")
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
func (a *ScoringGroupAccessor) Raw() ([]ScoringGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScoringGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScoringGroupAccessor.LoadData to preload everything
func (a *ScoringGroupAccessor) GroupData() {
	a._dataScoringGroupID = map[json.Number]ScoringGroup{}
	for _, d := range a._data {
		a._dataScoringGroupID[d.ScoringGroupID] = d
	}
}

// ByScoringGroupID returns the ScoringGroup uniquely identified by ScoringGroupID
//
// Error is only non-nil if the source errors out
func (a *ScoringGroupAccessor) ByScoringGroupID(identifier json.Number) (ScoringGroup, error) {
	if a._dataScoringGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScoringGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataScoringGroupID[identifier], nil
}
