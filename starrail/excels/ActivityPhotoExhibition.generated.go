package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type ActivityPhotoExhibition struct {
	ActivityModuleID float64   `json:"ActivityModuleID"`
	CommentList      []float64 `json:"CommentList"`
	Daily            hash.Hash `json:"Daily"`
	GroupID          float64   `json:"GroupID"`
	PhotoID          []float64 `json:"PhotoID"`
	QuestID          float64   `json:"QuestID"`
	Tab              hash.Hash `json:"Tab"`
}
type ActivityPhotoExhibitionAccessor struct {
	_data        []ActivityPhotoExhibition
	_dataGroupID map[float64]ActivityPhotoExhibition
}

// LoadData retrieves the data. Must be called before ActivityPhotoExhibition.GroupData
func (a *ActivityPhotoExhibitionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ActivityPhotoExhibition.json")
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
func (a *ActivityPhotoExhibitionAccessor) Raw() ([]ActivityPhotoExhibition, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ActivityPhotoExhibition{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ActivityPhotoExhibitionAccessor.LoadData to preload everything
func (a *ActivityPhotoExhibitionAccessor) GroupData() {
	for _, d := range a._data {
		a._dataGroupID[d.GroupID] = d
	}
}

// ByGroupID returns the ActivityPhotoExhibition uniquely identified by GroupID
//
// Error is only non-nil if the source errors out
func (a *ActivityPhotoExhibitionAccessor) ByGroupID(identifier float64) (ActivityPhotoExhibition, error) {
	if a._dataGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ActivityPhotoExhibition{}, err
			}
		}
		a.GroupData()
	}
	return a._dataGroupID[identifier], nil
}
