package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type SwordTrainingExamComment struct {
	CommentID json.Number            `json:"CommentID"`
	Desc      map[string]json.Number `json:"Desc"`
	ImgPath   string                 `json:"ImgPath"`
}
type SwordTrainingExamCommentAccessor struct {
	_data []SwordTrainingExamComment
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
