package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ResourceDeletionUsmList struct {
	ID   json.Number `json:"ID"`
	Path string      `json:"Path"`
}
type ResourceDeletionUsmListAccessor struct {
	_data     []ResourceDeletionUsmList
	_dataID   map[json.Number]ResourceDeletionUsmList
	_dataPath map[string]ResourceDeletionUsmList
}

// LoadData retrieves the data. Must be called before ResourceDeletionUsmList.GroupData
func (a *ResourceDeletionUsmListAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ResourceDeletionUsmList.json")
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
func (a *ResourceDeletionUsmListAccessor) Raw() ([]ResourceDeletionUsmList, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ResourceDeletionUsmList{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ResourceDeletionUsmListAccessor.LoadData to preload everything
func (a *ResourceDeletionUsmListAccessor) GroupData() {
	a._dataID = map[json.Number]ResourceDeletionUsmList{}
	a._dataPath = map[string]ResourceDeletionUsmList{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataPath[d.Path] = d
	}
}

// ByID returns the ResourceDeletionUsmList uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *ResourceDeletionUsmListAccessor) ByID(identifier json.Number) (ResourceDeletionUsmList, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ResourceDeletionUsmList{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByPath returns the ResourceDeletionUsmList uniquely identified by Path
//
// Error is only non-nil if the source errors out
func (a *ResourceDeletionUsmListAccessor) ByPath(identifier string) (ResourceDeletionUsmList, error) {
	if a._dataPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ResourceDeletionUsmList{}, err
			}
		}
		a.GroupData()
	}
	return a._dataPath[identifier], nil
}
