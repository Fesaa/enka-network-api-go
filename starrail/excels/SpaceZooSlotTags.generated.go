package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type SpaceZooSlotTags struct {
	Channel   string  `json:"Channel"`
	FeatureID float64 `json:"FeatureID"`
	ImagePath string  `json:"ImagePath"`
}
type SpaceZooSlotTagsAccessor struct {
	_data          []SpaceZooSlotTags
	_dataFeatureID map[float64]SpaceZooSlotTags
	_dataImagePath map[string]SpaceZooSlotTags
}

// LoadData retrieves the data. Must be called before SpaceZooSlotTags.GroupData
func (a *SpaceZooSlotTagsAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SpaceZooSlotTags.json")
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
func (a *SpaceZooSlotTagsAccessor) Raw() ([]SpaceZooSlotTags, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SpaceZooSlotTags{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SpaceZooSlotTagsAccessor.LoadData to preload everything
func (a *SpaceZooSlotTagsAccessor) GroupData() {
	for _, d := range a._data {
		a._dataFeatureID[d.FeatureID] = d
		a._dataImagePath[d.ImagePath] = d
	}
}

// ByFeatureID returns the SpaceZooSlotTags uniquely identified by FeatureID
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSlotTagsAccessor) ByFeatureID(identifier float64) (SpaceZooSlotTags, error) {
	if a._dataFeatureID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSlotTags{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFeatureID[identifier], nil
}

// ByImagePath returns the SpaceZooSlotTags uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *SpaceZooSlotTagsAccessor) ByImagePath(identifier string) (SpaceZooSlotTags, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SpaceZooSlotTags{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}
