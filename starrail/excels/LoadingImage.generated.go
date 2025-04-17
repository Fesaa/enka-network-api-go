package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type LoadingImage struct {
	ID        float64 `json:"ID"`
	ImagePath string  `json:"ImagePath"`
}
type LoadingImageAccessor struct {
	_data          []LoadingImage
	_dataID        map[float64]LoadingImage
	_dataImagePath map[string]LoadingImage
}

// LoadData retrieves the data. Must be called before LoadingImage.GroupData
func (a *LoadingImageAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/LoadingImage.json")
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
func (a *LoadingImageAccessor) Raw() ([]LoadingImage, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []LoadingImage{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with LoadingImageAccessor.LoadData to preload everything
func (a *LoadingImageAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataImagePath[d.ImagePath] = d
	}
}

// ByID returns the LoadingImage uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *LoadingImageAccessor) ByID(identifier float64) (LoadingImage, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return LoadingImage{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByImagePath returns the LoadingImage uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *LoadingImageAccessor) ByImagePath(identifier string) (LoadingImage, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return LoadingImage{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}
