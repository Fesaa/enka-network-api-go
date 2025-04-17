package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type MonopolyCellResource struct {
	IconPath   string      `json:"IconPath"`
	ResourceID json.Number `json:"ResourceID"`
	Type       string      `json:"Type"`
}
type MonopolyCellResourceAccessor struct {
	_data         []MonopolyCellResource
	_dataIconPath map[string]MonopolyCellResource
}

// LoadData retrieves the data. Must be called before MonopolyCellResource.GroupData
func (a *MonopolyCellResourceAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyCellResource.json")
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
func (a *MonopolyCellResourceAccessor) Raw() ([]MonopolyCellResource, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyCellResource{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyCellResourceAccessor.LoadData to preload everything
func (a *MonopolyCellResourceAccessor) GroupData() {
	a._dataIconPath = map[string]MonopolyCellResource{}
	for _, d := range a._data {
		a._dataIconPath[d.IconPath] = d
	}
}

// ByIconPath returns the MonopolyCellResource uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *MonopolyCellResourceAccessor) ByIconPath(identifier string) (MonopolyCellResource, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonopolyCellResource{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
