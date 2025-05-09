package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PhotoExhibitionComment struct {
	ID          float64   `json:"ID"`
	Name        hash.Hash `json:"Name"`
	NpcHandIcon string    `json:"NpcHandIcon"`
	Reply       hash.Hash `json:"Reply"`
}
type PhotoExhibitionCommentAccessor struct {
	_data   []PhotoExhibitionComment
	_dataID map[float64]PhotoExhibitionComment
}

// LoadData retrieves the data. Must be called before PhotoExhibitionComment.GroupData
func (a *PhotoExhibitionCommentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PhotoExhibitionComment.json")
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
func (a *PhotoExhibitionCommentAccessor) Raw() ([]PhotoExhibitionComment, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PhotoExhibitionComment{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PhotoExhibitionCommentAccessor.LoadData to preload everything
func (a *PhotoExhibitionCommentAccessor) GroupData() {
	a._dataID = map[float64]PhotoExhibitionComment{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PhotoExhibitionComment uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PhotoExhibitionCommentAccessor) ByID(identifier float64) (PhotoExhibitionComment, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PhotoExhibitionComment{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
