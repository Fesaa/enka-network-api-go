package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ScheduleDataGlobal struct {
	BeginTime     string  `json:"BeginTime"`
	EndTime       string  `json:"EndTime"`
	GlobalEndTime string  `json:"GlobalEndTime"`
	ID            float64 `json:"ID"`
}
type ScheduleDataGlobalAccessor struct {
	_data              []ScheduleDataGlobal
	_dataBeginTime     map[string]ScheduleDataGlobal
	_dataEndTime       map[string]ScheduleDataGlobal
	_dataGlobalEndTime map[string]ScheduleDataGlobal
	_dataID            map[float64]ScheduleDataGlobal
}

// LoadData retrieves the data. Must be called before ScheduleDataGlobal.GroupData
func (a *ScheduleDataGlobalAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScheduleDataGlobal.json")
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
func (a *ScheduleDataGlobalAccessor) Raw() ([]ScheduleDataGlobal, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScheduleDataGlobal{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScheduleDataGlobalAccessor.LoadData to preload everything
func (a *ScheduleDataGlobalAccessor) GroupData() {
	a._dataBeginTime = map[string]ScheduleDataGlobal{}
	a._dataEndTime = map[string]ScheduleDataGlobal{}
	a._dataGlobalEndTime = map[string]ScheduleDataGlobal{}
	a._dataID = map[float64]ScheduleDataGlobal{}
	for _, d := range a._data {
		a._dataBeginTime[d.BeginTime] = d
		a._dataEndTime[d.EndTime] = d
		a._dataGlobalEndTime[d.GlobalEndTime] = d
		a._dataID[d.ID] = d
	}
}

// ByBeginTime returns the ScheduleDataGlobal uniquely identified by BeginTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataGlobalAccessor) ByBeginTime(identifier string) (ScheduleDataGlobal, error) {
	if a._dataBeginTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataGlobal{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBeginTime[identifier], nil
}

// ByEndTime returns the ScheduleDataGlobal uniquely identified by EndTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataGlobalAccessor) ByEndTime(identifier string) (ScheduleDataGlobal, error) {
	if a._dataEndTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataGlobal{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEndTime[identifier], nil
}

// ByGlobalEndTime returns the ScheduleDataGlobal uniquely identified by GlobalEndTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataGlobalAccessor) ByGlobalEndTime(identifier string) (ScheduleDataGlobal, error) {
	if a._dataGlobalEndTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataGlobal{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGlobalEndTime[identifier], nil
}

// ByID returns the ScheduleDataGlobal uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataGlobalAccessor) ByID(identifier float64) (ScheduleDataGlobal, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataGlobal{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
