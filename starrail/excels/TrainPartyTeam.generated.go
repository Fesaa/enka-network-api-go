package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyTeam struct {
	GridNum             float64   `json:"GridNum"`
	InitialMeetingSkill float64   `json:"InitialMeetingSkill"`
	LeaderWorkingBuffID float64   `json:"LeaderWorkingBuffID"`
	PassengerList       []float64 `json:"PassengerList"`
	TeamID              float64   `json:"TeamID"`
	TeamName            hash.Hash `json:"TeamName"`
}
type TrainPartyTeamAccessor struct {
	_data       []TrainPartyTeam
	_dataTeamID map[float64]TrainPartyTeam
}

// LoadData retrieves the data. Must be called before TrainPartyTeam.GroupData
func (a *TrainPartyTeamAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyTeam.json")
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
func (a *TrainPartyTeamAccessor) Raw() ([]TrainPartyTeam, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyTeam{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyTeamAccessor.LoadData to preload everything
func (a *TrainPartyTeamAccessor) GroupData() {
	a._dataTeamID = map[float64]TrainPartyTeam{}
	for _, d := range a._data {
		a._dataTeamID[d.TeamID] = d
	}
}

// ByTeamID returns the TrainPartyTeam uniquely identified by TeamID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyTeamAccessor) ByTeamID(identifier float64) (TrainPartyTeam, error) {
	if a._dataTeamID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyTeam{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTeamID[identifier], nil
}
