package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type PlayerIcon struct {
	ID        json.Number `json:"ID"`
	ImagePath string      `json:"ImagePath"`
	IsVisible bool        `json:"IsVisible"`
}
type PlayerIconAccessor struct {
	_data          []PlayerIcon
	_dataID        map[json.Number]PlayerIcon
	_dataImagePath map[string]PlayerIcon
}

// LoadData retrieves the data. Must be called before PlayerIcon.GroupData
func (a *PlayerIconAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerIcon.json")
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
func (a *PlayerIconAccessor) Raw() ([]PlayerIcon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerIcon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerIconAccessor.LoadData to preload everything
func (a *PlayerIconAccessor) GroupData() {
	a._dataID = map[json.Number]PlayerIcon{}
	a._dataImagePath = map[string]PlayerIcon{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataImagePath[d.ImagePath] = d
	}
}

// ByID returns the PlayerIcon uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlayerIconAccessor) ByID(identifier json.Number) (PlayerIcon, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByImagePath returns the PlayerIcon uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *PlayerIconAccessor) ByImagePath(identifier string) (PlayerIcon, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}
