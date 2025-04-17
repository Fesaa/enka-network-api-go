package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type StoryLine struct {
	BeginCondition         StoryLineBeginCondition `json:"BeginCondition"`
	EarlyAccessContentID   json.Number             `json:"EarlyAccessContentID"`
	EndCondition           StoryLineEndCondition   `json:"EndCondition"`
	InitAnchorID           json.Number             `json:"InitAnchorID"`
	InitEntranceID         json.Number             `json:"InitEntranceID"`
	InitGroupID            json.Number             `json:"InitGroupID"`
	PerformanceStoryAvatar string                  `json:"PerformanceStoryAvatar"`
	ShowCondition          string                  `json:"ShowCondition"`
	StoryLineID            json.Number             `json:"StoryLineID"`
}
type StoryLineBeginCondition struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type StoryLineEndCondition struct {
	Param string `json:"Param"`
	Type  string `json:"Type"`
}
type StoryLineAccessor struct {
	_data              []StoryLine
	_dataShowCondition map[string]StoryLine
	_dataStoryLineID   map[json.Number]StoryLine
}

// LoadData retrieves the data. Must be called before StoryLine.GroupData
func (a *StoryLineAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StoryLine.json")
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
func (a *StoryLineAccessor) Raw() ([]StoryLine, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StoryLine{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StoryLineAccessor.LoadData to preload everything
func (a *StoryLineAccessor) GroupData() {
	a._dataShowCondition = map[string]StoryLine{}
	a._dataStoryLineID = map[json.Number]StoryLine{}
	for _, d := range a._data {
		a._dataShowCondition[d.ShowCondition] = d
		a._dataStoryLineID[d.StoryLineID] = d
	}
}

// ByShowCondition returns the StoryLine uniquely identified by ShowCondition
//
// Error is only non-nil if the source errors out
func (a *StoryLineAccessor) ByShowCondition(identifier string) (StoryLine, error) {
	if a._dataShowCondition == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StoryLine{}, err
			}
		}
		a.GroupData()
	}
	return a._dataShowCondition[identifier], nil
}

// ByStoryLineID returns the StoryLine uniquely identified by StoryLineID
//
// Error is only non-nil if the source errors out
func (a *StoryLineAccessor) ByStoryLineID(identifier json.Number) (StoryLine, error) {
	if a._dataStoryLineID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StoryLine{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStoryLineID[identifier], nil
}
