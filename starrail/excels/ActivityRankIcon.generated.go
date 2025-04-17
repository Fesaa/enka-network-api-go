package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ActivityRankIcon struct {
	CommonRankIconPath string                 `json:"CommonRankIconPath"`
	ID                 string                 `json:"ID"`
	Text               map[string]json.Number `json:"Text"`
}
type ActivityRankIconAccessor struct {
	_data                   []ActivityRankIcon
	_dataCommonRankIconPath map[string]ActivityRankIcon
	_dataID                 map[string]ActivityRankIcon
}

// LoadData retrieves the data. Must be called before ActivityRankIcon.GroupData
func (a *ActivityRankIconAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityRankIcon.json")
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
func (a *ActivityRankIconAccessor) Raw() ([]ActivityRankIcon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityRankIcon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityRankIconAccessor.LoadData to preload everything
func (a *ActivityRankIconAccessor) GroupData() {
	a._dataCommonRankIconPath = map[string]ActivityRankIcon{}
	a._dataID = map[string]ActivityRankIcon{}
	for _, d := range a._data {
		a._dataCommonRankIconPath[d.CommonRankIconPath] = d
		a._dataID[d.ID] = d
	}
}

// ByCommonRankIconPath returns the ActivityRankIcon uniquely identified by CommonRankIconPath
//
// Error is only non-nil if the source errors out
func (a *ActivityRankIconAccessor) ByCommonRankIconPath(identifier string) (ActivityRankIcon, error) {
	if a._dataCommonRankIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRankIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCommonRankIconPath[identifier], nil
}

// ByID returns the ActivityRankIcon uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ActivityRankIconAccessor) ByID(identifier string) (ActivityRankIcon, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityRankIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
