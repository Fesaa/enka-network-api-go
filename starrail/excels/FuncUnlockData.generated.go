package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type FuncUnlockData struct {
	Conditions    []FuncUnlockDataConditions `json:"Conditions"`
	ShowCondition []string                   `json:"ShowCondition"`
	UnlockID      json.Number                `json:"UnlockID"`
}
type FuncUnlockDataConditions struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type FuncUnlockDataAccessor struct {
	_data         []FuncUnlockData
	_dataUnlockID map[json.Number]FuncUnlockData
}

// LoadData retrieves the data. Must be called before FuncUnlockData.GroupData
func (a *FuncUnlockDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FuncUnlockData.json")
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
func (a *FuncUnlockDataAccessor) Raw() ([]FuncUnlockData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FuncUnlockData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FuncUnlockDataAccessor.LoadData to preload everything
func (a *FuncUnlockDataAccessor) GroupData() {
	a._dataUnlockID = map[json.Number]FuncUnlockData{}
	for _, d := range a._data {
		a._dataUnlockID[d.UnlockID] = d
	}
}

// ByUnlockID returns the FuncUnlockData uniquely identified by UnlockID
//
// Error is only non-nil if the source errors out
func (a *FuncUnlockDataAccessor) ByUnlockID(identifier json.Number) (FuncUnlockData, error) {
	if a._dataUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FuncUnlockData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockID[identifier], nil
}
