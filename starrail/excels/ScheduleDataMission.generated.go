package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ScheduleDataMission struct {
	BeginTime string      `json:"BeginTime"`
	EndTime   string      `json:"EndTime"`
	ID        json.Number `json:"ID"`
}
type ScheduleDataMissionAccessor struct {
	_data   []ScheduleDataMission
	_dataID map[json.Number]ScheduleDataMission
}

// LoadData retrieves the data. Must be called before ScheduleDataMission.GroupData
func (a *ScheduleDataMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScheduleDataMission.json")
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
func (a *ScheduleDataMissionAccessor) Raw() ([]ScheduleDataMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScheduleDataMission{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScheduleDataMissionAccessor.LoadData to preload everything
func (a *ScheduleDataMissionAccessor) GroupData() {
	a._dataID = map[json.Number]ScheduleDataMission{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ScheduleDataMission uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataMissionAccessor) ByID(identifier json.Number) (ScheduleDataMission, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
