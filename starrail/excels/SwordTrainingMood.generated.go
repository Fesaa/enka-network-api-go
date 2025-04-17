package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type SwordTrainingMood struct {
	EffectDesc    map[string]json.Number `json:"EffectDesc"`
	EffectIDList  []json.Number          `json:"EffectIDList"`
	EffectNumDesc map[string]json.Number `json:"EffectNumDesc"`
	MaximumValue  json.Number            `json:"MaximumValue"`
	MinimumValue  json.Number            `json:"MinimumValue"`
	MoodIcon      string                 `json:"MoodIcon"`
	MoodLevel     json.Number            `json:"MoodLevel"`
	MoodStatus    string                 `json:"MoodStatus"`
}
type SwordTrainingMoodAccessor struct {
	_data             []SwordTrainingMood
	_dataMaximumValue map[json.Number]SwordTrainingMood
	_dataMoodLevel    map[json.Number]SwordTrainingMood
	_dataMoodStatus   map[string]SwordTrainingMood
}

// LoadData retrieves the data. Must be called before SwordTrainingMood.GroupData
func (a *SwordTrainingMoodAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingMood.json")
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
func (a *SwordTrainingMoodAccessor) Raw() ([]SwordTrainingMood, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingMood{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingMoodAccessor.LoadData to preload everything
func (a *SwordTrainingMoodAccessor) GroupData() {
	a._dataMaximumValue = map[json.Number]SwordTrainingMood{}
	a._dataMoodLevel = map[json.Number]SwordTrainingMood{}
	a._dataMoodStatus = map[string]SwordTrainingMood{}
	for _, d := range a._data {
		a._dataMaximumValue[d.MaximumValue] = d
		a._dataMoodLevel[d.MoodLevel] = d
		a._dataMoodStatus[d.MoodStatus] = d
	}
}

// ByMaximumValue returns the SwordTrainingMood uniquely identified by MaximumValue
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingMoodAccessor) ByMaximumValue(identifier json.Number) (SwordTrainingMood, error) {
	if a._dataMaximumValue == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingMood{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMaximumValue[identifier], nil
}

// ByMoodLevel returns the SwordTrainingMood uniquely identified by MoodLevel
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingMoodAccessor) ByMoodLevel(identifier json.Number) (SwordTrainingMood, error) {
	if a._dataMoodLevel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingMood{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMoodLevel[identifier], nil
}

// ByMoodStatus returns the SwordTrainingMood uniquely identified by MoodStatus
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingMoodAccessor) ByMoodStatus(identifier string) (SwordTrainingMood, error) {
	if a._dataMoodStatus == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingMood{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMoodStatus[identifier], nil
}
