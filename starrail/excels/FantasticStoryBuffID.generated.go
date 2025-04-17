package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type FantasticStoryBuffID struct {
	ActivityModuleID        float64   `json:"ActivityModuleID"`
	AvailableBattleID       []float64 `json:"AvailableBattleID"`
	BuffID                  float64   `json:"BuffID"`
	BuffSlot                float64   `json:"BuffSlot"`
	ClientShowAvailableTips bool      `json:"ClientShowAvailableTips"`
	MazebuffID              float64   `json:"MazebuffID"`
	UnlockChapterID         float64   `json:"UnlockChapterID"`
}
type FantasticStoryBuffIDAccessor struct {
	_data           []FantasticStoryBuffID
	_dataBuffID     map[float64]FantasticStoryBuffID
	_dataMazebuffID map[float64]FantasticStoryBuffID
}

// LoadData retrieves the data. Must be called before FantasticStoryBuffID.GroupData
func (a *FantasticStoryBuffIDAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/FantasticStoryBuffID.json")
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
func (a *FantasticStoryBuffIDAccessor) Raw() ([]FantasticStoryBuffID, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []FantasticStoryBuffID{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with FantasticStoryBuffIDAccessor.LoadData to preload everything
func (a *FantasticStoryBuffIDAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBuffID[d.BuffID] = d
		a._dataMazebuffID[d.MazebuffID] = d
	}
}

// ByBuffID returns the FantasticStoryBuffID uniquely identified by BuffID
//
// Error is only non-nil if the source errors out
func (a *FantasticStoryBuffIDAccessor) ByBuffID(identifier float64) (FantasticStoryBuffID, error) {
	if a._dataBuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FantasticStoryBuffID{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBuffID[identifier], nil
}

// ByMazebuffID returns the FantasticStoryBuffID uniquely identified by MazebuffID
//
// Error is only non-nil if the source errors out
func (a *FantasticStoryBuffIDAccessor) ByMazebuffID(identifier float64) (FantasticStoryBuffID, error) {
	if a._dataMazebuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return FantasticStoryBuffID{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMazebuffID[identifier], nil
}
