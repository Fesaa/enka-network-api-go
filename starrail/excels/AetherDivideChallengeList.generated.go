package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AetherDivideChallengeList struct {
	BattleAreaID          float64   `json:"BattleAreaID"`
	ChallengeType         string    `json:"ChallengeType"`
	EventID               float64   `json:"EventID"`
	GroupID               float64   `json:"GroupID"`
	ID                    float64   `json:"ID"`
	MissionID             float64   `json:"MissionID"`
	OpponentImageIconPath string    `json:"OpponentImageIconPath"`
	OpponentImagePath     string    `json:"OpponentImagePath"`
	OpponentName          hash.Hash `json:"OpponentName"`
	OpponentPrefabPath    string    `json:"OpponentPrefabPath"`
	OpponentStrength      float64   `json:"OpponentStrength"`
	Rank                  float64   `json:"Rank"`
	RewardID              float64   `json:"RewardID"`
	VersusImagePath       string    `json:"VersusImagePath"`
}
type AetherDivideChallengeListAccessor struct {
	_data        []AetherDivideChallengeList
	_dataEventID map[float64]AetherDivideChallengeList
	_dataID      map[float64]AetherDivideChallengeList
}

// LoadData retrieves the data. Must be called before AetherDivideChallengeList.GroupData
func (a *AetherDivideChallengeListAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AetherDivideChallengeList.json")
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
func (a *AetherDivideChallengeListAccessor) Raw() ([]AetherDivideChallengeList, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AetherDivideChallengeList{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AetherDivideChallengeListAccessor.LoadData to preload everything
func (a *AetherDivideChallengeListAccessor) GroupData() {
	for _, d := range a._data {
		a._dataEventID[d.EventID] = d
		a._dataID[d.ID] = d
	}
}

// ByEventID returns the AetherDivideChallengeList uniquely identified by EventID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideChallengeListAccessor) ByEventID(identifier float64) (AetherDivideChallengeList, error) {
	if a._dataEventID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideChallengeList{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventID[identifier], nil
}

// ByID returns the AetherDivideChallengeList uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AetherDivideChallengeListAccessor) ByID(identifier float64) (AetherDivideChallengeList, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AetherDivideChallengeList{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
