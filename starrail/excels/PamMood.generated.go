package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PamMood struct {
	EmotionClipPath string  `json:"EmotionClipPath"`
	MaxMoodPoint    float64 `json:"MaxMoodPoint"`
	MinMoodPoint    float64 `json:"MinMoodPoint"`
	PamMood         string  `json:"PamMood"`
	PerformanceID   float64 `json:"PerformanceID"`
}
type PamMoodAccessor struct {
	_data              []PamMood
	_dataPamMood       map[string]PamMood
	_dataPerformanceID map[float64]PamMood
}

// LoadData retrieves the data. Must be called before PamMood.GroupData
func (a *PamMoodAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PamMood.json")
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
func (a *PamMoodAccessor) Raw() ([]PamMood, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PamMood{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PamMoodAccessor.LoadData to preload everything
func (a *PamMoodAccessor) GroupData() {
	a._dataPamMood = map[string]PamMood{}
	a._dataPerformanceID = map[float64]PamMood{}
	for _, d := range a._data {
		a._dataPamMood[d.PamMood] = d
		a._dataPerformanceID[d.PerformanceID] = d
	}
}

// ByPamMood returns the PamMood uniquely identified by PamMood
//
// Error is only non-nil if the source errors out
func (a *PamMoodAccessor) ByPamMood(identifier string) (PamMood, error) {
	if a._dataPamMood == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PamMood{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPamMood[identifier], nil
}

// ByPerformanceID returns the PamMood uniquely identified by PerformanceID
//
// Error is only non-nil if the source errors out
func (a *PamMoodAccessor) ByPerformanceID(identifier float64) (PamMood, error) {
	if a._dataPerformanceID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PamMood{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPerformanceID[identifier], nil
}
