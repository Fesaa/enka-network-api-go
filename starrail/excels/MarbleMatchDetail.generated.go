package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type MarbleMatchDetail struct {
	ID      json.Number   `json:"ID"`
	NpcList []json.Number `json:"NpcList"`
}
type MarbleMatchDetailAccessor struct {
	_data   []MarbleMatchDetail
	_dataID map[json.Number]MarbleMatchDetail
}

// LoadData retrieves the data. Must be called before MarbleMatchDetail.GroupData
func (a *MarbleMatchDetailAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MarbleMatchDetail.json")
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
func (a *MarbleMatchDetailAccessor) Raw() ([]MarbleMatchDetail, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MarbleMatchDetail{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MarbleMatchDetailAccessor.LoadData to preload everything
func (a *MarbleMatchDetailAccessor) GroupData() {
	a._dataID = map[json.Number]MarbleMatchDetail{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MarbleMatchDetail uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MarbleMatchDetailAccessor) ByID(identifier json.Number) (MarbleMatchDetail, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MarbleMatchDetail{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
