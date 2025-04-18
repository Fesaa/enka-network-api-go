package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ScheduleDataChallengeMaze struct {
	BeginTime string  `json:"BeginTime"`
	EndTime   string  `json:"EndTime"`
	ID        float64 `json:"ID"`
}
type ScheduleDataChallengeMazeAccessor struct {
	_data          []ScheduleDataChallengeMaze
	_dataBeginTime map[string]ScheduleDataChallengeMaze
	_dataEndTime   map[string]ScheduleDataChallengeMaze
	_dataID        map[float64]ScheduleDataChallengeMaze
}

// LoadData retrieves the data. Must be called before ScheduleDataChallengeMaze.GroupData
func (a *ScheduleDataChallengeMazeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScheduleDataChallengeMaze.json")
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
func (a *ScheduleDataChallengeMazeAccessor) Raw() ([]ScheduleDataChallengeMaze, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScheduleDataChallengeMaze{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScheduleDataChallengeMazeAccessor.LoadData to preload everything
func (a *ScheduleDataChallengeMazeAccessor) GroupData() {
	a._dataBeginTime = map[string]ScheduleDataChallengeMaze{}
	a._dataEndTime = map[string]ScheduleDataChallengeMaze{}
	a._dataID = map[float64]ScheduleDataChallengeMaze{}
	for _, d := range a._data {
		a._dataBeginTime[d.BeginTime] = d
		a._dataEndTime[d.EndTime] = d
		a._dataID[d.ID] = d
	}
}

// ByBeginTime returns the ScheduleDataChallengeMaze uniquely identified by BeginTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataChallengeMazeAccessor) ByBeginTime(identifier string) (ScheduleDataChallengeMaze, error) {
	if a._dataBeginTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataChallengeMaze{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBeginTime[identifier], nil
}

// ByEndTime returns the ScheduleDataChallengeMaze uniquely identified by EndTime
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataChallengeMazeAccessor) ByEndTime(identifier string) (ScheduleDataChallengeMaze, error) {
	if a._dataEndTime == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataChallengeMaze{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEndTime[identifier], nil
}

// ByID returns the ScheduleDataChallengeMaze uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataChallengeMazeAccessor) ByID(identifier float64) (ScheduleDataChallengeMaze, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataChallengeMaze{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
