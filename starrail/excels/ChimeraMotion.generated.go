package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ChimeraMotion struct {
	MotionID  float64 `json:"MotionID"`
	MotionKey string  `json:"MotionKey"`
}
type ChimeraMotionAccessor struct {
	_data          []ChimeraMotion
	_dataMotionID  map[float64]ChimeraMotion
	_dataMotionKey map[string]ChimeraMotion
}

// LoadData retrieves the data. Must be called before ChimeraMotion.GroupData
func (a *ChimeraMotionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ChimeraMotion.json")
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
func (a *ChimeraMotionAccessor) Raw() ([]ChimeraMotion, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ChimeraMotion{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ChimeraMotionAccessor.LoadData to preload everything
func (a *ChimeraMotionAccessor) GroupData() {
	a._dataMotionID = map[float64]ChimeraMotion{}
	a._dataMotionKey = map[string]ChimeraMotion{}
	for _, d := range a._data {
		a._dataMotionID[d.MotionID] = d
		a._dataMotionKey[d.MotionKey] = d
	}
}

// ByMotionID returns the ChimeraMotion uniquely identified by MotionID
//
// Error is only non-nil if the source errors out
func (a *ChimeraMotionAccessor) ByMotionID(identifier float64) (ChimeraMotion, error) {
	if a._dataMotionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraMotion{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMotionID[identifier], nil
}

// ByMotionKey returns the ChimeraMotion uniquely identified by MotionKey
//
// Error is only non-nil if the source errors out
func (a *ChimeraMotionAccessor) ByMotionKey(identifier string) (ChimeraMotion, error) {
	if a._dataMotionKey == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ChimeraMotion{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMotionKey[identifier], nil
}
