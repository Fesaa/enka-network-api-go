package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AlleyMission struct {
	EventEffect    []float64 `json:"EventEffect"`
	IsMissionTrack bool      `json:"IsMissionTrack"`
	IsUrgent       bool      `json:"IsUrgent"`
	MissionID      float64   `json:"MissionID"`
	NextMission    float64   `json:"NextMission"`
}
type AlleyMissionAccessor struct {
	_data          []AlleyMission
	_dataMissionID map[float64]AlleyMission
}

// LoadData retrieves the data. Must be called before AlleyMission.GroupData
func (a *AlleyMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AlleyMission.json")
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
func (a *AlleyMissionAccessor) Raw() ([]AlleyMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AlleyMission{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AlleyMissionAccessor.LoadData to preload everything
func (a *AlleyMissionAccessor) GroupData() {
	a._dataMissionID = map[float64]AlleyMission{}
	for _, d := range a._data {
		a._dataMissionID[d.MissionID] = d
	}
}

// ByMissionID returns the AlleyMission uniquely identified by MissionID
//
// Error is only non-nil if the source errors out
func (a *AlleyMissionAccessor) ByMissionID(identifier float64) (AlleyMission, error) {
	if a._dataMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AlleyMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionID[identifier], nil
}
