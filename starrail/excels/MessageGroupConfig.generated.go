package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MessageGroupConfig struct {
	ActivityModuleID     float64   `json:"ActivityModuleID"`
	ID                   float64   `json:"ID"`
	MessageContactsID    float64   `json:"MessageContactsID"`
	MessageSectionIDList []float64 `json:"MessageSectionIDList"`
}
type MessageGroupConfigAccessor struct {
	_data   []MessageGroupConfig
	_dataID map[float64]MessageGroupConfig
}

// LoadData retrieves the data. Must be called before MessageGroupConfig.GroupData
func (a *MessageGroupConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MessageGroupConfig.json")
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
func (a *MessageGroupConfigAccessor) Raw() ([]MessageGroupConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MessageGroupConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MessageGroupConfigAccessor.LoadData to preload everything
func (a *MessageGroupConfigAccessor) GroupData() {
	a._dataID = map[float64]MessageGroupConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MessageGroupConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MessageGroupConfigAccessor) ByID(identifier float64) (MessageGroupConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MessageGroupConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
