package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type HeliobusChallengeRaid struct {
	ChallengeRaidID      json.Number            `json:"ChallengeRaidID"`
	HeliobusSkillRecList []json.Number          `json:"HeliobusSkillRecList"`
	RaidID               json.Number            `json:"RaidID"`
	UnlockQuestID        json.Number            `json:"UnlockQuestID"`
	UnlockTips           map[string]json.Number `json:"UnlockTips"`
}
type HeliobusChallengeRaidAccessor struct {
	_data                []HeliobusChallengeRaid
	_dataChallengeRaidID map[json.Number]HeliobusChallengeRaid
	_dataRaidID          map[json.Number]HeliobusChallengeRaid
	_dataUnlockQuestID   map[json.Number]HeliobusChallengeRaid
}

// LoadData retrieves the data. Must be called before HeliobusChallengeRaid.GroupData
func (a *HeliobusChallengeRaidAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HeliobusChallengeRaid.json")
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
func (a *HeliobusChallengeRaidAccessor) Raw() ([]HeliobusChallengeRaid, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HeliobusChallengeRaid{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HeliobusChallengeRaidAccessor.LoadData to preload everything
func (a *HeliobusChallengeRaidAccessor) GroupData() {
	a._dataChallengeRaidID = map[json.Number]HeliobusChallengeRaid{}
	a._dataRaidID = map[json.Number]HeliobusChallengeRaid{}
	a._dataUnlockQuestID = map[json.Number]HeliobusChallengeRaid{}
	for _, d := range a._data {
		a._dataChallengeRaidID[d.ChallengeRaidID] = d
		a._dataRaidID[d.RaidID] = d
		a._dataUnlockQuestID[d.UnlockQuestID] = d
	}
}

// ByChallengeRaidID returns the HeliobusChallengeRaid uniquely identified by ChallengeRaidID
//
// Error is only non-nil if the source errors out
func (a *HeliobusChallengeRaidAccessor) ByChallengeRaidID(identifier json.Number) (HeliobusChallengeRaid, error) {
	if a._dataChallengeRaidID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusChallengeRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChallengeRaidID[identifier], nil
}

// ByRaidID returns the HeliobusChallengeRaid uniquely identified by RaidID
//
// Error is only non-nil if the source errors out
func (a *HeliobusChallengeRaidAccessor) ByRaidID(identifier json.Number) (HeliobusChallengeRaid, error) {
	if a._dataRaidID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusChallengeRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRaidID[identifier], nil
}

// ByUnlockQuestID returns the HeliobusChallengeRaid uniquely identified by UnlockQuestID
//
// Error is only non-nil if the source errors out
func (a *HeliobusChallengeRaidAccessor) ByUnlockQuestID(identifier json.Number) (HeliobusChallengeRaid, error) {
	if a._dataUnlockQuestID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HeliobusChallengeRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockQuestID[identifier], nil
}
