package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type LoopCGConfig struct {
	CaptionPath      string      `json:"CaptionPath"`
	Encryption       bool        `json:"Encryption"`
	IsPlayerInvolved bool        `json:"IsPlayerInvolved"`
	VideoID          json.Number `json:"VideoID"`
	VideoPath        string      `json:"VideoPath"`
}
type LoopCGConfigAccessor struct {
	_data          []LoopCGConfig
	_dataVideoID   map[json.Number]LoopCGConfig
	_dataVideoPath map[string]LoopCGConfig
}

// LoadData retrieves the data. Must be called before LoopCGConfig.GroupData
func (a *LoopCGConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/LoopCGConfig.json")
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
func (a *LoopCGConfigAccessor) Raw() ([]LoopCGConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []LoopCGConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with LoopCGConfigAccessor.LoadData to preload everything
func (a *LoopCGConfigAccessor) GroupData() {
	a._dataVideoID = map[json.Number]LoopCGConfig{}
	a._dataVideoPath = map[string]LoopCGConfig{}
	for _, d := range a._data {
		a._dataVideoID[d.VideoID] = d
		a._dataVideoPath[d.VideoPath] = d
	}
}

// ByVideoID returns the LoopCGConfig uniquely identified by VideoID
//
// Error is only non-nil if the source errors out
func (a *LoopCGConfigAccessor) ByVideoID(identifier json.Number) (LoopCGConfig, error) {
	if a._dataVideoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return LoopCGConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVideoID[identifier], nil
}

// ByVideoPath returns the LoopCGConfig uniquely identified by VideoPath
//
// Error is only non-nil if the source errors out
func (a *LoopCGConfigAccessor) ByVideoPath(identifier string) (LoopCGConfig, error) {
	if a._dataVideoPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return LoopCGConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVideoPath[identifier], nil
}
