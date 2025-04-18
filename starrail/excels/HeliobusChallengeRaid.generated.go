package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type HeliobusChallengeRaid struct {
	ChallengeRaidID      float64   `json:"ChallengeRaidID"`
	HeliobusSkillRecList []float64 `json:"HeliobusSkillRecList"`
	RaidID               float64   `json:"RaidID"`
	UnlockQuestID        float64   `json:"UnlockQuestID"`
	UnlockTips           hash.Hash `json:"UnlockTips"`
}
type HeliobusChallengeRaidAccessor struct {
	_data                []HeliobusChallengeRaid
	_dataChallengeRaidID map[float64]HeliobusChallengeRaid
	_dataRaidID          map[float64]HeliobusChallengeRaid
	_dataUnlockQuestID   map[float64]HeliobusChallengeRaid
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
	a._dataChallengeRaidID = map[float64]HeliobusChallengeRaid{}
	a._dataRaidID = map[float64]HeliobusChallengeRaid{}
	a._dataUnlockQuestID = map[float64]HeliobusChallengeRaid{}
	for _, d := range a._data {
		a._dataChallengeRaidID[d.ChallengeRaidID] = d
		a._dataRaidID[d.RaidID] = d
		a._dataUnlockQuestID[d.UnlockQuestID] = d
	}
}

// ByChallengeRaidID returns the HeliobusChallengeRaid uniquely identified by ChallengeRaidID
//
// Error is only non-nil if the source errors out
func (a *HeliobusChallengeRaidAccessor) ByChallengeRaidID(identifier float64) (HeliobusChallengeRaid, error) {
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
func (a *HeliobusChallengeRaidAccessor) ByRaidID(identifier float64) (HeliobusChallengeRaid, error) {
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
func (a *HeliobusChallengeRaidAccessor) ByUnlockQuestID(identifier float64) (HeliobusChallengeRaid, error) {
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
