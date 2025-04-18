package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type EventMissionChallenge struct {
	ID             float64 `json:"ID"`
	IsBeginPrepare bool    `json:"IsBeginPrepare"`
	IsCancellable  bool    `json:"IsCancellable"`
	IsResetable    bool    `json:"IsResetable"`
	LimitTime      float64 `json:"LimitTime"`
}
type EventMissionChallengeAccessor struct {
	_data   []EventMissionChallenge
	_dataID map[float64]EventMissionChallenge
}

// LoadData retrieves the data. Must be called before EventMissionChallenge.GroupData
func (a *EventMissionChallengeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EventMissionChallenge.json")
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
func (a *EventMissionChallengeAccessor) Raw() ([]EventMissionChallenge, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EventMissionChallenge{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EventMissionChallengeAccessor.LoadData to preload everything
func (a *EventMissionChallengeAccessor) GroupData() {
	a._dataID = map[float64]EventMissionChallenge{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the EventMissionChallenge uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *EventMissionChallengeAccessor) ByID(identifier float64) (EventMissionChallenge, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMissionChallenge{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
