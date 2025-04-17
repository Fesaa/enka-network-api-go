package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type StoryAtlas struct {
	AvatarID  json.Number            `json:"AvatarID"`
	ReplaceID json.Number            `json:"ReplaceID"`
	SortID    json.Number            `json:"SortID"`
	Story     map[string]json.Number `json:"Story"`
	StoryID   json.Number            `json:"StoryID"`
	Unlock    json.Number            `json:"Unlock"`
}
type StoryAtlasAccessor struct {
	_data []StoryAtlas
}

// LoadData retrieves the data. Must be called before StoryAtlas.GroupData
func (a *StoryAtlasAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StoryAtlas.json")
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
func (a *StoryAtlasAccessor) Raw() ([]StoryAtlas, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StoryAtlas{}, err
		}
	}
	return a._data, nil
}
