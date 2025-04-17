package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MainMissionSchedule struct {
	ActivityModuleID float64 `json:"ActivityModuleID"`
	HideRemainTime   bool    `json:"HideRemainTime"`
	IsNotDelete      bool    `json:"IsNotDelete"`
	MainMissionID    float64 `json:"MainMissionID"`
	ScheduleDataID   float64 `json:"ScheduleDataID"`
}
type MainMissionScheduleAccessor struct {
	_data              []MainMissionSchedule
	_dataMainMissionID map[float64]MainMissionSchedule
}

// LoadData retrieves the data. Must be called before MainMissionSchedule.GroupData
func (a *MainMissionScheduleAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MainMissionSchedule.json")
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
func (a *MainMissionScheduleAccessor) Raw() ([]MainMissionSchedule, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MainMissionSchedule{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MainMissionScheduleAccessor.LoadData to preload everything
func (a *MainMissionScheduleAccessor) GroupData() {
	for _, d := range a._data {
		a._dataMainMissionID[d.MainMissionID] = d
	}
}

// ByMainMissionID returns the MainMissionSchedule uniquely identified by MainMissionID
//
// Error is only non-nil if the source errors out
func (a *MainMissionScheduleAccessor) ByMainMissionID(identifier float64) (MainMissionSchedule, error) {
	if a._dataMainMissionID == nil {
		err := a.LoadData()
		if err != nil {
			return MainMissionSchedule{}, err
		}
		a.GroupData()
	}
	return a._dataMainMissionID[identifier], nil
}
