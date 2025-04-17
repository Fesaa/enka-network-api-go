package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FinishWayEventMission struct {
	FinishType    string        `json:"FinishType"`
	ID            float64       `json:"ID"`
	ParamInt1     float64       `json:"ParamInt1"`
	ParamInt2     float64       `json:"ParamInt2"`
	ParamInt3     float64       `json:"ParamInt3"`
	ParamIntList  []interface{} `json:"ParamIntList"`
	ParamItemList []interface{} `json:"ParamItemList"`
	ParamStr1     string        `json:"ParamStr1"`
	ParamType     string        `json:"ParamType"`
	Progress      float64       `json:"Progress"`
}
type FinishWayEventMissionAccessor struct {
	_data   []FinishWayEventMission
	_dataID map[float64]FinishWayEventMission
}

// LoadData retrieves the data. Must be called before FinishWayEventMission.GroupData
func (a *FinishWayEventMissionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FinishWayEventMission.json")
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
func (a *FinishWayEventMissionAccessor) Raw() ([]FinishWayEventMission, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FinishWayEventMission{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FinishWayEventMissionAccessor.LoadData to preload everything
func (a *FinishWayEventMissionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the FinishWayEventMission uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *FinishWayEventMissionAccessor) ByID(identifier float64) (FinishWayEventMission, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FinishWayEventMission{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
