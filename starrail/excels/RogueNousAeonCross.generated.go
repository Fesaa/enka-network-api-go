package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueNousAeonCross struct {
	BuffGroup   float64 `json:"BuffGroup"`
	MainAeonID  float64 `json:"MainAeonID"`
	MainAeonNum float64 `json:"MainAeonNum"`
	SubAeonID   float64 `json:"SubAeonID"`
	SubAeonNum  float64 `json:"SubAeonNum"`
}
type RogueNousAeonCrossAccessor struct {
	_data          []RogueNousAeonCross
	_dataBuffGroup map[float64]RogueNousAeonCross
}

// LoadData retrieves the data. Must be called before RogueNousAeonCross.GroupData
func (a *RogueNousAeonCrossAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousAeonCross.json")
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
func (a *RogueNousAeonCrossAccessor) Raw() ([]RogueNousAeonCross, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousAeonCross{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueNousAeonCrossAccessor.LoadData to preload everything
func (a *RogueNousAeonCrossAccessor) GroupData() {
	a._dataBuffGroup = map[float64]RogueNousAeonCross{}
	for _, d := range a._data {
		a._dataBuffGroup[d.BuffGroup] = d
	}
}

// ByBuffGroup returns the RogueNousAeonCross uniquely identified by BuffGroup
//
// Error is only non-nil if the source errors out
func (a *RogueNousAeonCrossAccessor) ByBuffGroup(identifier float64) (RogueNousAeonCross, error) {
	if a._dataBuffGroup == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueNousAeonCross{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBuffGroup[identifier], nil
}
