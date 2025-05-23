package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type StoryAtlas struct {
	AvatarID  float64   `json:"AvatarID"`
	ReplaceID float64   `json:"ReplaceID"`
	SortID    float64   `json:"SortID"`
	Story     hash.Hash `json:"Story"`
	StoryID   float64   `json:"StoryID"`
	Unlock    float64   `json:"Unlock"`
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
