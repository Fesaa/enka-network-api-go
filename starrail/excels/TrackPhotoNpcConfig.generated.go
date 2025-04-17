package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type TrackPhotoNpcConfig struct {
	CanTypeID string      `json:"CanTypeID"`
	GroupID   json.Number `json:"GroupID"`
	NpcID     json.Number `json:"NpcID"`
	StageID   json.Number `json:"StageID"`
}
type TrackPhotoNpcConfigAccessor struct {
	_data []TrackPhotoNpcConfig
}

// LoadData retrieves the data. Must be called before TrackPhotoNpcConfig.GroupData
func (a *TrackPhotoNpcConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrackPhotoNpcConfig.json")
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
func (a *TrackPhotoNpcConfigAccessor) Raw() ([]TrackPhotoNpcConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrackPhotoNpcConfig{}, err
		}
	}
	return a._data, nil
}
