package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueNousSurfaceTag struct {
	Sort    float64   `json:"Sort"`
	TagID   float64   `json:"TagID"`
	TagName hash.Hash `json:"TagName"`
}
type RogueNousSurfaceTagAccessor struct {
	_data      []RogueNousSurfaceTag
	_dataTagID map[float64]RogueNousSurfaceTag
	_dataSort  map[float64]RogueNousSurfaceTag
}

// LoadData retrieves the data. Must be called before RogueNousSurfaceTag.GroupData
func (a *RogueNousSurfaceTagAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueNousSurfaceTag.json")
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
func (a *RogueNousSurfaceTagAccessor) Raw() ([]RogueNousSurfaceTag, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueNousSurfaceTag{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueNousSurfaceTagAccessor.LoadData to preload everything
func (a *RogueNousSurfaceTagAccessor) GroupData() {
	for _, d := range a._data {
		a._dataTagID[d.TagID] = d
		a._dataSort[d.Sort] = d
	}
}

// ByTagID returns the RogueNousSurfaceTag uniquely identified by TagID
//
// Error is only non-nil if the source errors out
func (a *RogueNousSurfaceTagAccessor) ByTagID(identifier float64) (RogueNousSurfaceTag, error) {
	if a._dataTagID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueNousSurfaceTag{}, err
		}
		a.GroupData()
	}
	return a._dataTagID[identifier], nil
}

// BySort returns the RogueNousSurfaceTag uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *RogueNousSurfaceTagAccessor) BySort(identifier float64) (RogueNousSurfaceTag, error) {
	if a._dataSort == nil {
		err := a.LoadData()
		if err != nil {
			return RogueNousSurfaceTag{}, err
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}
