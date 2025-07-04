package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type StoryCharacter struct {
	ConfigEntityPath string `json:"ConfigEntityPath"`
	JsonPath         string `json:"JsonPath"`
	StoryCharacterID string `json:"StoryCharacterID"`
	SubType          string `json:"SubType"`
}
type StoryCharacterAccessor struct {
	_data                 []StoryCharacter
	_dataStoryCharacterID map[string]StoryCharacter
}

// LoadData retrieves the data. Must be called before StoryCharacter.GroupData
func (a *StoryCharacterAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StoryCharacter.json")
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
func (a *StoryCharacterAccessor) Raw() ([]StoryCharacter, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StoryCharacter{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StoryCharacterAccessor.LoadData to preload everything
func (a *StoryCharacterAccessor) GroupData() {
	a._dataStoryCharacterID = map[string]StoryCharacter{}
	for _, d := range a._data {
		a._dataStoryCharacterID[d.StoryCharacterID] = d
	}
}

// ByStoryCharacterID returns the StoryCharacter uniquely identified by StoryCharacterID
//
// Error is only non-nil if the source errors out
func (a *StoryCharacterAccessor) ByStoryCharacterID(identifier string) (StoryCharacter, error) {
	if a._dataStoryCharacterID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StoryCharacter{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStoryCharacterID[identifier], nil
}
