package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ScheduleDataBattlePass struct {
	BeginTime string  `json:"BeginTime"`
	EndTime   string  `json:"EndTime"`
	ID        float64 `json:"ID"`
}
type ScheduleDataBattlePassAccessor struct {
	_data          []ScheduleDataBattlePass
	_dataBeginTime map[string]ScheduleDataBattlePass
	_dataEndTime   map[string]ScheduleDataBattlePass
	_dataID        map[float64]ScheduleDataBattlePass
}

// LoadData retrieves the data. Must be called before ScheduleDataBattlePass.GroupData
func (a *ScheduleDataBattlePassAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScheduleDataBattlePass.json")
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
func (a *ScheduleDataBattlePassAccessor) Raw() ([]ScheduleDataBattlePass, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScheduleDataBattlePass{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScheduleDataBattlePassAccessor.LoadData to preload everything
func (a *ScheduleDataBattlePassAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBeginTime[d.BeginTime] = d
		a._dataEndTime[d.EndTime] = d
		a._dataID[d.ID] = d
	}
}

// ByBeginTime returns the ScheduleDataBattlePass uniquely identified by BeginTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataBattlePassAccessor) ByBeginTime(identifier string) (ScheduleDataBattlePass, error) {
	if a._dataBeginTime == nil {
		err := a.LoadData()
		if err != nil {
			return ScheduleDataBattlePass{}, err
		}
		a.GroupData()
	}
	return a._dataBeginTime[identifier], nil
}

// ByEndTime returns the ScheduleDataBattlePass uniquely identified by EndTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataBattlePassAccessor) ByEndTime(identifier string) (ScheduleDataBattlePass, error) {
	if a._dataEndTime == nil {
		err := a.LoadData()
		if err != nil {
			return ScheduleDataBattlePass{}, err
		}
		a.GroupData()
	}
	return a._dataEndTime[identifier], nil
}

// ByID returns the ScheduleDataBattlePass uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataBattlePassAccessor) ByID(identifier float64) (ScheduleDataBattlePass, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return ScheduleDataBattlePass{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
