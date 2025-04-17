package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type EventMissionChallenge struct {
	ID             json.Number `json:"ID"`
	IsBeginPrepare bool        `json:"IsBeginPrepare"`
	IsCancellable  bool        `json:"IsCancellable"`
	IsResetable    bool        `json:"IsResetable"`
	LimitTime      json.Number `json:"LimitTime"`
}
type EventMissionChallengeAccessor struct {
	_data []EventMissionChallenge
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
