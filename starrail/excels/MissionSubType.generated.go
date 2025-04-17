package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MissionSubType struct {
	Type         string  `json:"Type"`
	TypePriority float64 `json:"TypePriority"`
}
type MissionSubTypeAccessor struct {
	_data             []MissionSubType
	_dataTypePriority map[float64]MissionSubType
}

// LoadData retrieves the data. Must be called before MissionSubType.GroupData
func (a *MissionSubTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MissionSubType.json")
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
func (a *MissionSubTypeAccessor) Raw() ([]MissionSubType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MissionSubType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MissionSubTypeAccessor.LoadData to preload everything
func (a *MissionSubTypeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataTypePriority[d.TypePriority] = d
	}
}

// ByTypePriority returns the MissionSubType uniquely identified by TypePriority
//
// Error is only non-nil if the source errors out
func (a *MissionSubTypeAccessor) ByTypePriority(identifier float64) (MissionSubType, error) {
	if a._dataTypePriority == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MissionSubType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataTypePriority[identifier], nil
}
