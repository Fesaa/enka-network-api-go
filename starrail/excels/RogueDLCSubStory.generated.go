package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueDLCSubStory struct {
	ImgPath            string    `json:"ImgPath"`
	Layer              float64   `json:"Layer"`
	LevelGraphPath     string    `json:"LevelGraphPath"`
	OptionPath         string    `json:"OptionPath"`
	RogueDLCSubStoryID float64   `json:"RogueDLCSubStoryID"`
	SubStoryName       hash.Hash `json:"SubStoryName"`
}
type RogueDLCSubStoryAccessor struct {
	_data                   []RogueDLCSubStory
	_dataLevelGraphPath     map[string]RogueDLCSubStory
	_dataOptionPath         map[string]RogueDLCSubStory
	_dataRogueDLCSubStoryID map[float64]RogueDLCSubStory
}

// LoadData retrieves the data. Must be called before RogueDLCSubStory.GroupData
func (a *RogueDLCSubStoryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueDLCSubStory.json")
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
func (a *RogueDLCSubStoryAccessor) Raw() ([]RogueDLCSubStory, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueDLCSubStory{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueDLCSubStoryAccessor.LoadData to preload everything
func (a *RogueDLCSubStoryAccessor) GroupData() {
	a._dataLevelGraphPath = map[string]RogueDLCSubStory{}
	a._dataOptionPath = map[string]RogueDLCSubStory{}
	a._dataRogueDLCSubStoryID = map[float64]RogueDLCSubStory{}
	for _, d := range a._data {
		a._dataLevelGraphPath[d.LevelGraphPath] = d
		a._dataOptionPath[d.OptionPath] = d
		a._dataRogueDLCSubStoryID[d.RogueDLCSubStoryID] = d
	}
}

// ByLevelGraphPath returns the RogueDLCSubStory uniquely identified by LevelGraphPath
//
// Error is only non-nil if the source errors out
func (a *RogueDLCSubStoryAccessor) ByLevelGraphPath(identifier string) (RogueDLCSubStory, error) {
	if a._dataLevelGraphPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCSubStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataLevelGraphPath[identifier], nil
}

// ByOptionPath returns the RogueDLCSubStory uniquely identified by OptionPath
//
// Error is only non-nil if the source errors out
func (a *RogueDLCSubStoryAccessor) ByOptionPath(identifier string) (RogueDLCSubStory, error) {
	if a._dataOptionPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCSubStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataOptionPath[identifier], nil
}

// ByRogueDLCSubStoryID returns the RogueDLCSubStory uniquely identified by RogueDLCSubStoryID
//
// Error is only non-nil if the source errors out
func (a *RogueDLCSubStoryAccessor) ByRogueDLCSubStoryID(identifier float64) (RogueDLCSubStory, error) {
	if a._dataRogueDLCSubStoryID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueDLCSubStory{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRogueDLCSubStoryID[identifier], nil
}
