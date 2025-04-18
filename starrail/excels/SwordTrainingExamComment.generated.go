package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type SwordTrainingExamComment struct {
	CommentID float64   `json:"CommentID"`
	Desc      hash.Hash `json:"Desc"`
	ImgPath   string    `json:"ImgPath"`
}
type SwordTrainingExamCommentAccessor struct {
	_data          []SwordTrainingExamComment
	_dataCommentID map[float64]SwordTrainingExamComment
}

// LoadData retrieves the data. Must be called before SwordTrainingExamComment.GroupData
func (a *SwordTrainingExamCommentAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SwordTrainingExamComment.json")
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
func (a *SwordTrainingExamCommentAccessor) Raw() ([]SwordTrainingExamComment, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SwordTrainingExamComment{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SwordTrainingExamCommentAccessor.LoadData to preload everything
func (a *SwordTrainingExamCommentAccessor) GroupData() {
	a._dataCommentID = map[float64]SwordTrainingExamComment{}
	for _, d := range a._data {
		a._dataCommentID[d.CommentID] = d
	}
}

// ByCommentID returns the SwordTrainingExamComment uniquely identified by CommentID
//
// Error is only non-nil if the source errors out
func (a *SwordTrainingExamCommentAccessor) ByCommentID(identifier float64) (SwordTrainingExamComment, error) {
	if a._dataCommentID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SwordTrainingExamComment{}, err
			}
		}
		a.GroupData()
	}
	return a._dataCommentID[identifier], nil
}
