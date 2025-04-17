package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AvatarPlayerIcon struct {
	AvatarID  float64 `json:"AvatarID"`
	ID        float64 `json:"ID"`
	ImagePath string  `json:"ImagePath"`
	IsVisible bool    `json:"IsVisible"`
}
type AvatarPlayerIconAccessor struct {
	_data          []AvatarPlayerIcon
	_dataID        map[float64]AvatarPlayerIcon
	_dataImagePath map[string]AvatarPlayerIcon
	_dataAvatarID  map[float64]AvatarPlayerIcon
}

// LoadData retrieves the data. Must be called before AvatarPlayerIcon.GroupData
func (a *AvatarPlayerIconAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarPlayerIcon.json")
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
func (a *AvatarPlayerIconAccessor) Raw() ([]AvatarPlayerIcon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarPlayerIcon{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarPlayerIconAccessor.LoadData to preload everything
func (a *AvatarPlayerIconAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataImagePath[d.ImagePath] = d
		a._dataAvatarID[d.AvatarID] = d
	}
}

// ByID returns the AvatarPlayerIcon uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AvatarPlayerIconAccessor) ByID(identifier float64) (AvatarPlayerIcon, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return AvatarPlayerIcon{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByImagePath returns the AvatarPlayerIcon uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *AvatarPlayerIconAccessor) ByImagePath(identifier string) (AvatarPlayerIcon, error) {
	if a._dataImagePath == nil {
		err := a.LoadData()
		if err != nil {
			return AvatarPlayerIcon{}, err
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// ByAvatarID returns the AvatarPlayerIcon uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *AvatarPlayerIconAccessor) ByAvatarID(identifier float64) (AvatarPlayerIcon, error) {
	if a._dataAvatarID == nil {
		err := a.LoadData()
		if err != nil {
			return AvatarPlayerIcon{}, err
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}
