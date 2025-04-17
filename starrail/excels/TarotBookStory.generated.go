package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TarotBookStory struct {
	CardID      float64   `json:"CardID"`
	CharacterID float64   `json:"CharacterID"`
	ClueList    []float64 `json:"ClueList"`
	ID          float64   `json:"ID"`
	PreStoryID  float64   `json:"PreStoryID"`
}
type TarotBookStoryAccessor struct {
	_data   []TarotBookStory
	_dataID map[float64]TarotBookStory
}

// LoadData retrieves the data. Must be called before TarotBookStory.GroupData
func (a *TarotBookStoryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TarotBookStory.json")
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
func (a *TarotBookStoryAccessor) Raw() ([]TarotBookStory, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TarotBookStory{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TarotBookStoryAccessor.LoadData to preload everything
func (a *TarotBookStoryAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TarotBookStory uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TarotBookStoryAccessor) ByID(identifier float64) (TarotBookStory, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return TarotBookStory{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
