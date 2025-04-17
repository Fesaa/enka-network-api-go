package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MissionDisable struct {
	CompensateItemList             []interface{} `json:"CompensateItemList"`
	MainMissionIDList              []float64     `json:"MainMissionIDList"`
	MainMissionIDListClientDisplay []float64     `json:"MainMissionIDListClientDisplay"`
	RecycleItemList                []interface{} `json:"RecycleItemList"`
	SubMissionID                   float64       `json:"SubMissionID"`
}
type MissionDisableAccessor struct {
	_data             []MissionDisable
	_dataSubMissionID map[float64]MissionDisable
}

// LoadData retrieves the data. Must be called before MissionDisable.GroupData
func (a *MissionDisableAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MissionDisable.json")
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
func (a *MissionDisableAccessor) Raw() ([]MissionDisable, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MissionDisable{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MissionDisableAccessor.LoadData to preload everything
func (a *MissionDisableAccessor) GroupData() {
	a._dataSubMissionID = map[float64]MissionDisable{}
	for _, d := range a._data {
		a._dataSubMissionID[d.SubMissionID] = d
	}
}

// BySubMissionID returns the MissionDisable uniquely identified by SubMissionID
//
// Error is only non-nil if the source errors out
func (a *MissionDisableAccessor) BySubMissionID(identifier float64) (MissionDisable, error) {
	if a._dataSubMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MissionDisable{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSubMissionID[identifier], nil
}
