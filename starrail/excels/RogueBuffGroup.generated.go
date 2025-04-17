package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueBuffGroup struct {
	GNGDPDOMDFH []float64 `json:"GNGDPDOMDFH"`
	IDLBMIHBAPB float64   `json:"IDLBMIHBAPB"`
}
type RogueBuffGroupAccessor struct {
	_data            []RogueBuffGroup
	_dataIDLBMIHBAPB map[float64]RogueBuffGroup
}

// LoadData retrieves the data. Must be called before RogueBuffGroup.GroupData
func (a *RogueBuffGroupAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueBuffGroup.json")
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
func (a *RogueBuffGroupAccessor) Raw() ([]RogueBuffGroup, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueBuffGroup{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueBuffGroupAccessor.LoadData to preload everything
func (a *RogueBuffGroupAccessor) GroupData() {
	for _, d := range a._data {
		a._dataIDLBMIHBAPB[d.IDLBMIHBAPB] = d
	}
}

// ByIDLBMIHBAPB returns the RogueBuffGroup uniquely identified by IDLBMIHBAPB
//
// Error is only non-nil if the source errors out
func (a *RogueBuffGroupAccessor) ByIDLBMIHBAPB(identifier float64) (RogueBuffGroup, error) {
	if a._dataIDLBMIHBAPB == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueBuffGroup{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIDLBMIHBAPB[identifier], nil
}
