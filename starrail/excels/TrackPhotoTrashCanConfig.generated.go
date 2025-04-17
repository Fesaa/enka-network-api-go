package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type TrackPhotoTrashCanConfig struct {
	CanTypeID     string   `json:"CanTypeID"`
	ExtraAnimList []string `json:"ExtraAnimList"`
	ExtraScore    float64  `json:"ExtraScore"`
	IconPath      string   `json:"IconPath"`
	NpcTemplateID float64  `json:"NpcTemplateID"`
	Score         float64  `json:"Score"`
}
type TrackPhotoTrashCanConfigAccessor struct {
	_data              []TrackPhotoTrashCanConfig
	_dataCanTypeID     map[string]TrackPhotoTrashCanConfig
	_dataNpcTemplateID map[float64]TrackPhotoTrashCanConfig
}

// LoadData retrieves the data. Must be called before TrackPhotoTrashCanConfig.GroupData
func (a *TrackPhotoTrashCanConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrackPhotoTrashCanConfig.json")
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
func (a *TrackPhotoTrashCanConfigAccessor) Raw() ([]TrackPhotoTrashCanConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrackPhotoTrashCanConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrackPhotoTrashCanConfigAccessor.LoadData to preload everything
func (a *TrackPhotoTrashCanConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataCanTypeID[d.CanTypeID] = d
		a._dataNpcTemplateID[d.NpcTemplateID] = d
	}
}

// ByCanTypeID returns the TrackPhotoTrashCanConfig uniquely identified by CanTypeID
//
// Error is only non-nil if the source errors out
func (a *TrackPhotoTrashCanConfigAccessor) ByCanTypeID(identifier string) (TrackPhotoTrashCanConfig, error) {
	if a._dataCanTypeID == nil {
		err := a.LoadData()
		if err != nil {
			return TrackPhotoTrashCanConfig{}, err
		}
		a.GroupData()
	}
	return a._dataCanTypeID[identifier], nil
}

// ByNpcTemplateID returns the TrackPhotoTrashCanConfig uniquely identified by NpcTemplateID
//
// Error is only non-nil if the source errors out
func (a *TrackPhotoTrashCanConfigAccessor) ByNpcTemplateID(identifier float64) (TrackPhotoTrashCanConfig, error) {
	if a._dataNpcTemplateID == nil {
		err := a.LoadData()
		if err != nil {
			return TrackPhotoTrashCanConfig{}, err
		}
		a.GroupData()
	}
	return a._dataNpcTemplateID[identifier], nil
}
