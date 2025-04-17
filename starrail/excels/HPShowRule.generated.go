package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type HPShowRule struct {
	Color    string  `json:"Color"`
	ID       float64 `json:"ID"`
	IsDanger bool    `json:"IsDanger"`
	Max      float64 `json:"Max"`
}
type HPShowRuleAccessor struct {
	_data    []HPShowRule
	_dataID  map[float64]HPShowRule
	_dataMax map[float64]HPShowRule
}

// LoadData retrieves the data. Must be called before HPShowRule.GroupData
func (a *HPShowRuleAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/HPShowRule.json")
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
func (a *HPShowRuleAccessor) Raw() ([]HPShowRule, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []HPShowRule{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with HPShowRuleAccessor.LoadData to preload everything
func (a *HPShowRuleAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataMax[d.Max] = d
	}
}

// ByID returns the HPShowRule uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *HPShowRuleAccessor) ByID(identifier float64) (HPShowRule, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HPShowRule{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByMax returns the HPShowRule uniquely identified by Max
//
// Error is only non-nil if the source errors out
func (a *HPShowRuleAccessor) ByMax(identifier float64) (HPShowRule, error) {
	if a._dataMax == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return HPShowRule{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMax[identifier], nil
}
