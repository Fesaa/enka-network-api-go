package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicStory struct {
	IsHide         bool      `json:"IsHide"`
	LevelGraphPath string    `json:"LevelGraphPath"`
	StoryCategory  string    `json:"StoryCategory"`
	StoryID        float64   `json:"StoryID"`
	StoryImage     string    `json:"StoryImage"`
	StoryName      hash.Hash `json:"StoryName"`
	UnLockDisplay  float64   `json:"UnLockDisplay"`
}
type RogueMagicStoryAccessor struct {
	_data               []RogueMagicStory
	_dataLevelGraphPath map[string]RogueMagicStory
	_dataStoryID        map[float64]RogueMagicStory
}

// LoadData retrieves the data. Must be called before RogueMagicStory.GroupData
func (a *RogueMagicStoryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicStory.json")
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
func (a *RogueMagicStoryAccessor) Raw() ([]RogueMagicStory, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicStory{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicStoryAccessor.LoadData to preload everything
func (a *RogueMagicStoryAccessor) GroupData() {
	a._dataLevelGraphPath = map[string]RogueMagicStory{}
	a._dataStoryID = map[float64]RogueMagicStory{}
	for _, d := range a._data {
		a._dataLevelGraphPath[d.LevelGraphPath] = d
		a._dataStoryID[d.StoryID] = d
	}
}

// ByLevelGraphPath returns the RogueMagicStory uniquely identified by LevelGraphPath
//
// Error is only non-nil if the source errors out
func (a *RogueMagicStoryAccessor) ByLevelGraphPath(identifier string) (RogueMagicStory, error) {
	if a._dataLevelGraphPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLevelGraphPath[identifier], nil
}

// ByStoryID returns the RogueMagicStory uniquely identified by StoryID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicStoryAccessor) ByStoryID(identifier float64) (RogueMagicStory, error) {
	if a._dataStoryID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueMagicStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStoryID[identifier], nil
}
