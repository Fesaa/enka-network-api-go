package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type AvatarPlayerIconLD struct {
	AvatarID  float64 `json:"AvatarID"`
	ID        float64 `json:"ID"`
	ImagePath string  `json:"ImagePath"`
}
type AvatarPlayerIconLDAccessor struct {
	_data          []AvatarPlayerIconLD
	_dataAvatarID  map[float64]AvatarPlayerIconLD
	_dataID        map[float64]AvatarPlayerIconLD
	_dataImagePath map[string]AvatarPlayerIconLD
}

// LoadData retrieves the data. Must be called before AvatarPlayerIconLD.GroupData
func (a *AvatarPlayerIconLDAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AvatarPlayerIconLD.json")
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
func (a *AvatarPlayerIconLDAccessor) Raw() ([]AvatarPlayerIconLD, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AvatarPlayerIconLD{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AvatarPlayerIconLDAccessor.LoadData to preload everything
func (a *AvatarPlayerIconLDAccessor) GroupData() {
	a._dataAvatarID = map[float64]AvatarPlayerIconLD{}
	a._dataID = map[float64]AvatarPlayerIconLD{}
	a._dataImagePath = map[string]AvatarPlayerIconLD{}
	for _, d := range a._data {
		a._dataAvatarID[d.AvatarID] = d
		a._dataID[d.ID] = d
		a._dataImagePath[d.ImagePath] = d
	}
}

// ByAvatarID returns the AvatarPlayerIconLD uniquely identified by AvatarID
//
// Error is only non-nil if the source errors out
func (a *AvatarPlayerIconLDAccessor) ByAvatarID(identifier float64) (AvatarPlayerIconLD, error) {
	if a._dataAvatarID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarPlayerIconLD{}, err
			}
		}
		a.GroupData()
	}
	return a._dataAvatarID[identifier], nil
}

// ByID returns the AvatarPlayerIconLD uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AvatarPlayerIconLDAccessor) ByID(identifier float64) (AvatarPlayerIconLD, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarPlayerIconLD{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByImagePath returns the AvatarPlayerIconLD uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *AvatarPlayerIconLDAccessor) ByImagePath(identifier string) (AvatarPlayerIconLD, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AvatarPlayerIconLD{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}
