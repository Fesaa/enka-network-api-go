package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournWeeklyChallenge struct {
	ChallengeID               float64                                            `json:"ChallengeID"`
	DisplayFinalMonsterGroups RogueTournWeeklyChallengeDisplayFinalMonsterGroups `json:"DisplayFinalMonsterGroups"`
	DisplayMonsterGroups1     map[string]float64                                 `json:"DisplayMonsterGroups1"`
	DisplayMonsterGroups2     map[string]float64                                 `json:"DisplayMonsterGroups2"`
	DisplayMonsterGroups3     RogueTournWeeklyChallengeDisplayMonsterGroups3     `json:"DisplayMonsterGroups3"`
	RewardID                  float64                                            `json:"RewardID"`
	WeeklyContentDetailList   []float64                                          `json:"WeeklyContentDetailList"`
	WeeklyContentList         []float64                                          `json:"WeeklyContentList"`
	WeeklyName                hash.Hash                                          `json:"WeeklyName"`
}
type RogueTournWeeklyChallengeDisplayFinalMonsterGroups struct {
	F0 float64 `json:"0"`
}
type RogueTournWeeklyChallengeDisplayMonsterGroups3 struct {
	F0 float64 `json:"0"`
}
type RogueTournWeeklyChallengeAccessor struct {
	_data            []RogueTournWeeklyChallenge
	_dataChallengeID map[float64]RogueTournWeeklyChallenge
}

// LoadData retrieves the data. Must be called before RogueTournWeeklyChallenge.GroupData
func (a *RogueTournWeeklyChallengeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournWeeklyChallenge.json")
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
func (a *RogueTournWeeklyChallengeAccessor) Raw() ([]RogueTournWeeklyChallenge, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournWeeklyChallenge{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournWeeklyChallengeAccessor.LoadData to preload everything
func (a *RogueTournWeeklyChallengeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataChallengeID[d.ChallengeID] = d
	}
}

// ByChallengeID returns the RogueTournWeeklyChallenge uniquely identified by ChallengeID
//
// Error is only non-nil if the source errors out
func (a *RogueTournWeeklyChallengeAccessor) ByChallengeID(identifier float64) (RogueTournWeeklyChallenge, error) {
	if a._dataChallengeID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueTournWeeklyChallenge{}, err
		}
		a.GroupData()
	}
	return a._dataChallengeID[identifier], nil
}
