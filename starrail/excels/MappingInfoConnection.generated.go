package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MappingInfoConnection struct {
	SourceEntranceID    float64 `json:"SourceEntranceID"`
	SourceMappingInfoID float64 `json:"SourceMappingInfoID"`
	TargetEntranceID    float64 `json:"TargetEntranceID"`
	TargetMappingInfoID float64 `json:"TargetMappingInfoID"`
}
type MappingInfoConnectionAccessor struct {
	_data                    []MappingInfoConnection
	_dataSourceMappingInfoID map[float64]MappingInfoConnection
}

// LoadData retrieves the data. Must be called before MappingInfoConnection.GroupData
func (a *MappingInfoConnectionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MappingInfoConnection.json")
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
func (a *MappingInfoConnectionAccessor) Raw() ([]MappingInfoConnection, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MappingInfoConnection{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MappingInfoConnectionAccessor.LoadData to preload everything
func (a *MappingInfoConnectionAccessor) GroupData() {
	a._dataSourceMappingInfoID = map[float64]MappingInfoConnection{}
	for _, d := range a._data {
		a._dataSourceMappingInfoID[d.SourceMappingInfoID] = d
	}
}

// BySourceMappingInfoID returns the MappingInfoConnection uniquely identified by SourceMappingInfoID
//
// Error is only non-nil if the source errors out
func (a *MappingInfoConnectionAccessor) BySourceMappingInfoID(identifier float64) (MappingInfoConnection, error) {
	if a._dataSourceMappingInfoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MappingInfoConnection{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSourceMappingInfoID[identifier], nil
}
