package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type PhotoExhibitionComment struct {
	ID          json.Number            `json:"ID"`
	Name        map[string]json.Number `json:"Name"`
	NpcHandIcon string                 `json:"NpcHandIcon"`
	Reply       map[string]json.Number `json:"Reply"`
}
type PhotoExhibitionCommentAccessor struct {
	_data []PhotoExhibitionComment
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
