package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EventStuffConfig struct {
	EventContentTextID hash.Hash `json:"EventContentTextID"`
	EventStuffID       float64   `json:"EventStuffID"`
	MissionID          float64   `json:"MissionID"`
	MissionStartString string    `json:"MissionStartString"`
	StuffID            float64   `json:"StuffID"`
}
type EventStuffConfigAccessor struct {
	_data                   []EventStuffConfig
	_dataEventStuffID       map[float64]EventStuffConfig
	_dataMissionID          map[float64]EventStuffConfig
	_dataMissionStartString map[string]EventStuffConfig
	_dataStuffID            map[float64]EventStuffConfig
}

// LoadData retrieves the data. Must be called before EventStuffConfig.GroupData
func (a *EventStuffConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EventStuffConfig.json")
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
func (a *EventStuffConfigAccessor) Raw() ([]EventStuffConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EventStuffConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EventStuffConfigAccessor.LoadData to preload everything
func (a *EventStuffConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataEventStuffID[d.EventStuffID] = d
		a._dataMissionID[d.MissionID] = d
		a._dataMissionStartString[d.MissionStartString] = d
		a._dataStuffID[d.StuffID] = d
	}
}

// ByEventStuffID returns the EventStuffConfig uniquely identified by EventStuffID
//
// Error is only non-nil if the source errors out
func (a *EventStuffConfigAccessor) ByEventStuffID(identifier float64) (EventStuffConfig, error) {
	if a._dataEventStuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventStuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventStuffID[identifier], nil
}

// ByMissionID returns the EventStuffConfig uniquely identified by MissionID
//
// Error is only non-nil if the source errors out
func (a *EventStuffConfigAccessor) ByMissionID(identifier float64) (EventStuffConfig, error) {
	if a._dataMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventStuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionID[identifier], nil
}

// ByMissionStartString returns the EventStuffConfig uniquely identified by MissionStartString
//
// Error is only non-nil if the source errors out
func (a *EventStuffConfigAccessor) ByMissionStartString(identifier string) (EventStuffConfig, error) {
	if a._dataMissionStartString == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventStuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionStartString[identifier], nil
}

// ByStuffID returns the EventStuffConfig uniquely identified by StuffID
//
// Error is only non-nil if the source errors out
func (a *EventStuffConfigAccessor) ByStuffID(identifier float64) (EventStuffConfig, error) {
	if a._dataStuffID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventStuffConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStuffID[identifier], nil
}
