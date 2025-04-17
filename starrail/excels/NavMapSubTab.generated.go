package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type NavMapSubTab struct {
	FloorID                   float64 `json:"FloorID"`
	MenuSortID                float64 `json:"MenuSortID"`
	NavMapTabID               float64 `json:"NavMapTabID"`
	UnlockConditionExpression string  `json:"UnlockConditionExpression"`
}
type NavMapSubTabAccessor struct {
	_data        []NavMapSubTab
	_dataFloorID map[float64]NavMapSubTab
}

// LoadData retrieves the data. Must be called before NavMapSubTab.GroupData
func (a *NavMapSubTabAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/NavMapSubTab.json")
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
func (a *NavMapSubTabAccessor) Raw() ([]NavMapSubTab, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []NavMapSubTab{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with NavMapSubTabAccessor.LoadData to preload everything
func (a *NavMapSubTabAccessor) GroupData() {
	for _, d := range a._data {
		a._dataFloorID[d.FloorID] = d
	}
}

// ByFloorID returns the NavMapSubTab uniquely identified by FloorID
//
// Error is only non-nil if the source errors out
func (a *NavMapSubTabAccessor) ByFloorID(identifier float64) (NavMapSubTab, error) {
	if a._dataFloorID == nil {
		err := a.LoadData()
		if err != nil {
			return NavMapSubTab{}, err
		}
		a.GroupData()
	}
	return a._dataFloorID[identifier], nil
}
