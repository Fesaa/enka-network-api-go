package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MessageStateIcon struct {
	ID       string `json:"ID"`
	IconPath string `json:"IconPath"`
}
type MessageStateIconAccessor struct {
	_data         []MessageStateIcon
	_dataID       map[string]MessageStateIcon
	_dataIconPath map[string]MessageStateIcon
}

// LoadData retrieves the data. Must be called before MessageStateIcon.GroupData
func (a *MessageStateIconAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageStateIcon.json")
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
func (a *MessageStateIconAccessor) Raw() ([]MessageStateIcon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageStateIcon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MessageStateIconAccessor.LoadData to preload everything
func (a *MessageStateIconAccessor) GroupData() {
	a._dataID = map[string]MessageStateIcon{}
	a._dataIconPath = map[string]MessageStateIcon{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
	}
}

// ByID returns the MessageStateIcon uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MessageStateIconAccessor) ByID(identifier string) (MessageStateIcon, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageStateIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the MessageStateIcon uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *MessageStateIconAccessor) ByIconPath(identifier string) (MessageStateIcon, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageStateIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
