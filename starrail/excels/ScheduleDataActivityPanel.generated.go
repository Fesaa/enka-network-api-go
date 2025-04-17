package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ScheduleDataActivityPanel struct {
	BeginTime string  `json:"BeginTime"`
	EndTime   string  `json:"EndTime"`
	ID        float64 `json:"ID"`
}
type ScheduleDataActivityPanelAccessor struct {
	_data   []ScheduleDataActivityPanel
	_dataID map[float64]ScheduleDataActivityPanel
}

// LoadData retrieves the data. Must be called before ScheduleDataActivityPanel.GroupData
func (a *ScheduleDataActivityPanelAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ScheduleDataActivityPanel.json")
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
func (a *ScheduleDataActivityPanelAccessor) Raw() ([]ScheduleDataActivityPanel, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ScheduleDataActivityPanel{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ScheduleDataActivityPanelAccessor.LoadData to preload everything
func (a *ScheduleDataActivityPanelAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the ScheduleDataActivityPanel uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ScheduleDataActivityPanelAccessor) ByID(identifier float64) (ScheduleDataActivityPanel, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ScheduleDataActivityPanel{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
