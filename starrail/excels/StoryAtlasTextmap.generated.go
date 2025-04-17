package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type StoryAtlasTextmap struct {
	StoryID   float64   `json:"StoryID"`
	StoryName hash.Hash `json:"StoryName"`
}
type StoryAtlasTextmapAccessor struct {
	_data        []StoryAtlasTextmap
	_dataStoryID map[float64]StoryAtlasTextmap
}

// LoadData retrieves the data. Must be called before StoryAtlasTextmap.GroupData
func (a *StoryAtlasTextmapAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StoryAtlasTextmap.json")
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
func (a *StoryAtlasTextmapAccessor) Raw() ([]StoryAtlasTextmap, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StoryAtlasTextmap{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StoryAtlasTextmapAccessor.LoadData to preload everything
func (a *StoryAtlasTextmapAccessor) GroupData() {
	for _, d := range a._data {
		a._dataStoryID[d.StoryID] = d
	}
}

// ByStoryID returns the StoryAtlasTextmap uniquely identified by StoryID
//
// Error is only non-nil if the source errors out
func (a *StoryAtlasTextmapAccessor) ByStoryID(identifier float64) (StoryAtlasTextmap, error) {
	if a._dataStoryID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StoryAtlasTextmap{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStoryID[identifier], nil
}
