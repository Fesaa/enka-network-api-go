package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type EventMission struct {
	ClearGroupList       []interface{}          `json:"ClearGroupList"`
	Desc                 map[string]json.Number `json:"Desc"`
	FinishWayID          json.Number            `json:"FinishWayID"`
	ID                   json.Number            `json:"ID"`
	LoadGroupList        []interface{}          `json:"LoadGroupList"`
	MazeFloorID          json.Number            `json:"MazeFloorID"`
	MazePlaneID          json.Number            `json:"MazePlaneID"`
	MissionJsonPath      string                 `json:"MissionJsonPath"`
	NextEventMissionList []json.Number          `json:"NextEventMissionList"`
	RewardID             json.Number            `json:"RewardID"`
	TakeParamIntList     []json.Number          `json:"TakeParamIntList"`
	TakeType             string                 `json:"TakeType"`
	Title                map[string]json.Number `json:"Title"`
	Type                 string                 `json:"Type"`
	UnLoadGroupList      []interface{}          `json:"UnLoadGroupList"`
}
type EventMissionAccessor struct {
	_data            []EventMission
	_dataFinishWayID map[json.Number]EventMission
	_dataID          map[json.Number]EventMission
}

// LoadData retrieves the data. Must be called before EventMission.GroupData
func (a *EventMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EventMission.json")
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
func (a *EventMissionAccessor) Raw() ([]EventMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EventMission{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EventMissionAccessor.LoadData to preload everything
func (a *EventMissionAccessor) GroupData() {
	a._dataFinishWayID = map[json.Number]EventMission{}
	a._dataID = map[json.Number]EventMission{}
	for _, d := range a._data {
		a._dataFinishWayID[d.FinishWayID] = d
		a._dataID[d.ID] = d
	}
}

// ByFinishWayID returns the EventMission uniquely identified by FinishWayID
//
// Error is only non-nil if the source errors out
func (a *EventMissionAccessor) ByFinishWayID(identifier json.Number) (EventMission, error) {
	if a._dataFinishWayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFinishWayID[identifier], nil
}

// ByID returns the EventMission uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *EventMissionAccessor) ByID(identifier json.Number) (EventMission, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
