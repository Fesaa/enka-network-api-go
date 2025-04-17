package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ScheduleDataDropLimit struct {
	BeginTime string  `json:"BeginTime"`
	EndTime   string  `json:"EndTime"`
	ID        float64 `json:"ID"`
}
type ScheduleDataDropLimitAccessor struct {
	_data          []ScheduleDataDropLimit
	_dataBeginTime map[string]ScheduleDataDropLimit
	_dataEndTime   map[string]ScheduleDataDropLimit
	_dataID        map[float64]ScheduleDataDropLimit
}

// LoadData retrieves the data. Must be called before ScheduleDataDropLimit.GroupData
func (a *ScheduleDataDropLimitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScheduleDataDropLimit.json")
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
func (a *ScheduleDataDropLimitAccessor) Raw() ([]ScheduleDataDropLimit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScheduleDataDropLimit{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScheduleDataDropLimitAccessor.LoadData to preload everything
func (a *ScheduleDataDropLimitAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBeginTime[d.BeginTime] = d
		a._dataEndTime[d.EndTime] = d
		a._dataID[d.ID] = d
	}
}

// ByBeginTime returns the ScheduleDataDropLimit uniquely identified by BeginTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataDropLimitAccessor) ByBeginTime(identifier string) (ScheduleDataDropLimit, error) {
	if a._dataBeginTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataDropLimit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBeginTime[identifier], nil
}

// ByEndTime returns the ScheduleDataDropLimit uniquely identified by EndTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataDropLimitAccessor) ByEndTime(identifier string) (ScheduleDataDropLimit, error) {
	if a._dataEndTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataDropLimit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEndTime[identifier], nil
}

// ByID returns the ScheduleDataDropLimit uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataDropLimitAccessor) ByID(identifier float64) (ScheduleDataDropLimit, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataDropLimit{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
