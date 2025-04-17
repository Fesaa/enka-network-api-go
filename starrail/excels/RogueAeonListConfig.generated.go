package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueAeonListConfig struct {
	ActivityModuleID float64 `json:"ActivityModuleID"`
	DisplayID        float64 `json:"DisplayID"`
	RogueAeonID      float64 `json:"RogueAeonID"`
	Sort             float64 `json:"Sort"`
}
type RogueAeonListConfigAccessor struct {
	_data            []RogueAeonListConfig
	_dataDisplayID   map[float64]RogueAeonListConfig
	_dataSort        map[float64]RogueAeonListConfig
	_dataRogueAeonID map[float64]RogueAeonListConfig
}

// LoadData retrieves the data. Must be called before RogueAeonListConfig.GroupData
func (a *RogueAeonListConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueAeonListConfig.json")
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
func (a *RogueAeonListConfigAccessor) Raw() ([]RogueAeonListConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueAeonListConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueAeonListConfigAccessor.LoadData to preload everything
func (a *RogueAeonListConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataDisplayID[d.DisplayID] = d
		a._dataSort[d.Sort] = d
		a._dataRogueAeonID[d.RogueAeonID] = d
	}
}

// ByDisplayID returns the RogueAeonListConfig uniquely identified by DisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueAeonListConfigAccessor) ByDisplayID(identifier float64) (RogueAeonListConfig, error) {
	if a._dataDisplayID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonListConfig{}, err
		}
		a.GroupData()
	}
	return a._dataDisplayID[identifier], nil
}

// BySort returns the RogueAeonListConfig uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *RogueAeonListConfigAccessor) BySort(identifier float64) (RogueAeonListConfig, error) {
	if a._dataSort == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonListConfig{}, err
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}

// ByRogueAeonID returns the RogueAeonListConfig uniquely identified by RogueAeonID
//
// Error is only non-nil if the source errors out
func (a *RogueAeonListConfigAccessor) ByRogueAeonID(identifier float64) (RogueAeonListConfig, error) {
	if a._dataRogueAeonID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueAeonListConfig{}, err
		}
		a.GroupData()
	}
	return a._dataRogueAeonID[identifier], nil
}
