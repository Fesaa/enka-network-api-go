package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueManager struct {
	BeginTime       string        `json:"BeginTime"`
	EndTime         string        `json:"EndTime"`
	RogueAreaIDList []json.Number `json:"RogueAreaIDList"`
	RogueSeason     json.Number   `json:"RogueSeason"`
	RogueVersion    json.Number   `json:"RogueVersion"`
	ScheduleDataID  json.Number   `json:"ScheduleDataID"`
}
type RogueManagerAccessor struct {
	_data               []RogueManager
	_dataBeginTime      map[string]RogueManager
	_dataEndTime        map[string]RogueManager
	_dataRogueSeason    map[json.Number]RogueManager
	_dataScheduleDataID map[json.Number]RogueManager
}

// LoadData retrieves the data. Must be called before RogueManager.GroupData
func (a *RogueManagerAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueManager.json")
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
func (a *RogueManagerAccessor) Raw() ([]RogueManager, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueManager{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueManagerAccessor.LoadData to preload everything
func (a *RogueManagerAccessor) GroupData() {
	a._dataBeginTime = map[string]RogueManager{}
	a._dataEndTime = map[string]RogueManager{}
	a._dataRogueSeason = map[json.Number]RogueManager{}
	a._dataScheduleDataID = map[json.Number]RogueManager{}
	for _, d := range a._data {
		a._dataBeginTime[d.BeginTime] = d
		a._dataEndTime[d.EndTime] = d
		a._dataRogueSeason[d.RogueSeason] = d
		a._dataScheduleDataID[d.ScheduleDataID] = d
	}
}

// ByBeginTime returns the RogueManager uniquely identified by BeginTime
//
// Error is only non-nil if the source errors out
func (a *RogueManagerAccessor) ByBeginTime(identifier string) (RogueManager, error) {
	if a._dataBeginTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueManager{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBeginTime[identifier], nil
}

// ByEndTime returns the RogueManager uniquely identified by EndTime
//
// Error is only non-nil if the source errors out
func (a *RogueManagerAccessor) ByEndTime(identifier string) (RogueManager, error) {
	if a._dataEndTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueManager{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEndTime[identifier], nil
}

// ByRogueSeason returns the RogueManager uniquely identified by RogueSeason
//
// Error is only non-nil if the source errors out
func (a *RogueManagerAccessor) ByRogueSeason(identifier json.Number) (RogueManager, error) {
	if a._dataRogueSeason == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueManager{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueSeason[identifier], nil
}

// ByScheduleDataID returns the RogueManager uniquely identified by ScheduleDataID
//
// Error is only non-nil if the source errors out
func (a *RogueManagerAccessor) ByScheduleDataID(identifier json.Number) (RogueManager, error) {
	if a._dataScheduleDataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueManager{}, err
			}
		}
		a.GroupData()
	}
	return a._dataScheduleDataID[identifier], nil
}
