package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChallengeRaid struct {
	ChallengeID    float64   `json:"ChallengeID"`
	IconPath       string    `json:"IconPath"`
	MonsterList    []float64 `json:"MonsterList"`
	ScoringGroupID float64   `json:"ScoringGroupID"`
}
type ChallengeRaidAccessor struct {
	_data               []ChallengeRaid
	_dataChallengeID    map[float64]ChallengeRaid
	_dataScoringGroupID map[float64]ChallengeRaid
}

// LoadData retrieves the data. Must be called before ChallengeRaid.GroupData
func (a *ChallengeRaidAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChallengeRaid.json")
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
func (a *ChallengeRaidAccessor) Raw() ([]ChallengeRaid, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChallengeRaid{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChallengeRaidAccessor.LoadData to preload everything
func (a *ChallengeRaidAccessor) GroupData() {
	a._dataChallengeID = map[float64]ChallengeRaid{}
	a._dataScoringGroupID = map[float64]ChallengeRaid{}
	for _, d := range a._data {
		a._dataChallengeID[d.ChallengeID] = d
		a._dataScoringGroupID[d.ScoringGroupID] = d
	}
}

// ByChallengeID returns the ChallengeRaid uniquely identified by ChallengeID
//
// Error is only non-nil if the source errors out
func (a *ChallengeRaidAccessor) ByChallengeID(identifier float64) (ChallengeRaid, error) {
	if a._dataChallengeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataChallengeID[identifier], nil
}

// ByScoringGroupID returns the ChallengeRaid uniquely identified by ScoringGroupID
//
// Error is only non-nil if the source errors out
func (a *ChallengeRaidAccessor) ByScoringGroupID(identifier float64) (ChallengeRaid, error) {
	if a._dataScoringGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChallengeRaid{}, err
			}
		}
		a.GroupData()
	}
	return a._dataScoringGroupID[identifier], nil
}
