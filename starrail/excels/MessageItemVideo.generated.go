package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MessageItemVideo struct {
	ID        float64 `json:"ID"`
	ImagePath string  `json:"ImagePath"`
	VideoID   float64 `json:"VideoID"`
}
type MessageItemVideoAccessor struct {
	_data          []MessageItemVideo
	_dataID        map[float64]MessageItemVideo
	_dataImagePath map[string]MessageItemVideo
	_dataVideoID   map[float64]MessageItemVideo
}

// LoadData retrieves the data. Must be called before MessageItemVideo.GroupData
func (a *MessageItemVideoAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageItemVideo.json")
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
func (a *MessageItemVideoAccessor) Raw() ([]MessageItemVideo, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageItemVideo{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MessageItemVideoAccessor.LoadData to preload everything
func (a *MessageItemVideoAccessor) GroupData() {
	a._dataID = map[float64]MessageItemVideo{}
	a._dataImagePath = map[string]MessageItemVideo{}
	a._dataVideoID = map[float64]MessageItemVideo{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataImagePath[d.ImagePath] = d
		a._dataVideoID[d.VideoID] = d
	}
}

// ByID returns the MessageItemVideo uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MessageItemVideoAccessor) ByID(identifier float64) (MessageItemVideo, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageItemVideo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByImagePath returns the MessageItemVideo uniquely identified by ImagePath
//
// Error is only non-nil if the source errors out
func (a *MessageItemVideoAccessor) ByImagePath(identifier string) (MessageItemVideo, error) {
	if a._dataImagePath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageItemVideo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataImagePath[identifier], nil
}

// ByVideoID returns the MessageItemVideo uniquely identified by VideoID
//
// Error is only non-nil if the source errors out
func (a *MessageItemVideoAccessor) ByVideoID(identifier float64) (MessageItemVideo, error) {
	if a._dataVideoID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageItemVideo{}, err
			}
		}
		a.GroupData()
	}
	return a._dataVideoID[identifier], nil
}
