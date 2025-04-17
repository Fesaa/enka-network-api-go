package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type VideoConfig struct {
	CaptionPath      string  `json:"CaptionPath"`
	Encryption       bool    `json:"Encryption"`
	IsPlayerInvolved bool    `json:"IsPlayerInvolved"`
	VideoID          float64 `json:"VideoID"`
	VideoPath        string  `json:"VideoPath"`
}
type VideoConfigAccessor struct {
	_data          []VideoConfig
	_dataVideoID   map[float64]VideoConfig
	_dataVideoPath map[string]VideoConfig
}

// LoadData retrieves the data. Must be called before VideoConfig.GroupData
func (a *VideoConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/VideoConfig.json")
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
func (a *VideoConfigAccessor) Raw() ([]VideoConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []VideoConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with VideoConfigAccessor.LoadData to preload everything
func (a *VideoConfigAccessor) GroupData() {
	a._dataVideoID = map[float64]VideoConfig{}
	a._dataVideoPath = map[string]VideoConfig{}
	for _, d := range a._data {
		a._dataVideoID[d.VideoID] = d
		a._dataVideoPath[d.VideoPath] = d
	}
}

// ByVideoID returns the VideoConfig uniquely identified by VideoID
//
// Error is only non-nil if the source errors out
func (a *VideoConfigAccessor) ByVideoID(identifier float64) (VideoConfig, error) {
	if a._dataVideoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return VideoConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVideoID[identifier], nil
}

// ByVideoPath returns the VideoConfig uniquely identified by VideoPath
//
// Error is only non-nil if the source errors out
func (a *VideoConfigAccessor) ByVideoPath(identifier string) (VideoConfig, error) {
	if a._dataVideoPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return VideoConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVideoPath[identifier], nil
}
