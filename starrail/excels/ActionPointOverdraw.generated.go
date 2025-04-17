package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ActionPointOverdraw struct {
	ActionPoint float64 `json:"ActionPoint"`
	MazeBuff    float64 `json:"MazeBuff"`
}
type ActionPointOverdrawAccessor struct {
	_data            []ActionPointOverdraw
	_dataActionPoint map[float64]ActionPointOverdraw
}

// LoadData retrieves the data. Must be called before ActionPointOverdraw.GroupData
func (a *ActionPointOverdrawAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActionPointOverdraw.json")
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
func (a *ActionPointOverdrawAccessor) Raw() ([]ActionPointOverdraw, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActionPointOverdraw{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActionPointOverdrawAccessor.LoadData to preload everything
func (a *ActionPointOverdrawAccessor) GroupData() {
	a._dataActionPoint = map[float64]ActionPointOverdraw{}
	for _, d := range a._data {
		a._dataActionPoint[d.ActionPoint] = d
	}
}

// ByActionPoint returns the ActionPointOverdraw uniquely identified by ActionPoint
//
// Error is only non-nil if the source errors out
func (a *ActionPointOverdrawAccessor) ByActionPoint(identifier float64) (ActionPointOverdraw, error) {
	if a._dataActionPoint == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActionPointOverdraw{}, err
			}
		}
		a.GroupData()
	}
	return a._dataActionPoint[identifier], nil
}
