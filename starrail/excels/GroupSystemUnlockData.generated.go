package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type GroupSystemUnlockData struct {
	GroupSystemUnlockID float64 `json:"GroupSystemUnlockID"`
	UnlockID            float64 `json:"UnlockID"`
}
type GroupSystemUnlockDataAccessor struct {
	_data                    []GroupSystemUnlockData
	_dataGroupSystemUnlockID map[float64]GroupSystemUnlockData
	_dataUnlockID            map[float64]GroupSystemUnlockData
}

// LoadData retrieves the data. Must be called before GroupSystemUnlockData.GroupData
func (a *GroupSystemUnlockDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/GroupSystemUnlockData.json")
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
func (a *GroupSystemUnlockDataAccessor) Raw() ([]GroupSystemUnlockData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []GroupSystemUnlockData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with GroupSystemUnlockDataAccessor.LoadData to preload everything
func (a *GroupSystemUnlockDataAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGroupSystemUnlockID[d.GroupSystemUnlockID] = d
		a._dataUnlockID[d.UnlockID] = d
	}
}

// ByGroupSystemUnlockID returns the GroupSystemUnlockData uniquely identified by GroupSystemUnlockID
//
// Error is only non-nil if the source errors out
func (a *GroupSystemUnlockDataAccessor) ByGroupSystemUnlockID(identifier float64) (GroupSystemUnlockData, error) {
	if a._dataGroupSystemUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GroupSystemUnlockData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupSystemUnlockID[identifier], nil
}

// ByUnlockID returns the GroupSystemUnlockData uniquely identified by UnlockID
//
// Error is only non-nil if the source errors out
func (a *GroupSystemUnlockDataAccessor) ByUnlockID(identifier float64) (GroupSystemUnlockData, error) {
	if a._dataUnlockID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return GroupSystemUnlockData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUnlockID[identifier], nil
}
