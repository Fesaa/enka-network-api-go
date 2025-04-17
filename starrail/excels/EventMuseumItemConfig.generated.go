package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type EventMuseumItemConfig struct {
	EventContentTextID hash.Hash `json:"EventContentTextID"`
	EventMuseumItemID  float64   `json:"EventMuseumItemID"`
	ForceComplete      bool      `json:"ForceComplete"`
	IsTargetReward     bool      `json:"IsTargetReward"`
	MissionID          float64   `json:"MissionID"`
	MissionStartString string    `json:"MissionStartString"`
	MuseumItemID       float64   `json:"MuseumItemID"`
}
type EventMuseumItemConfigAccessor struct {
	_data                   []EventMuseumItemConfig
	_dataEventMuseumItemID  map[float64]EventMuseumItemConfig
	_dataMissionID          map[float64]EventMuseumItemConfig
	_dataMissionStartString map[string]EventMuseumItemConfig
	_dataMuseumItemID       map[float64]EventMuseumItemConfig
}

// LoadData retrieves the data. Must be called before EventMuseumItemConfig.GroupData
func (a *EventMuseumItemConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EventMuseumItemConfig.json")
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
func (a *EventMuseumItemConfigAccessor) Raw() ([]EventMuseumItemConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EventMuseumItemConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EventMuseumItemConfigAccessor.LoadData to preload everything
func (a *EventMuseumItemConfigAccessor) GroupData() {
	a._dataEventMuseumItemID = map[float64]EventMuseumItemConfig{}
	a._dataMissionID = map[float64]EventMuseumItemConfig{}
	a._dataMissionStartString = map[string]EventMuseumItemConfig{}
	a._dataMuseumItemID = map[float64]EventMuseumItemConfig{}
	for _, d := range a._data {
		a._dataEventMuseumItemID[d.EventMuseumItemID] = d
		a._dataMissionID[d.MissionID] = d
		a._dataMissionStartString[d.MissionStartString] = d
		a._dataMuseumItemID[d.MuseumItemID] = d
	}
}

// ByEventMuseumItemID returns the EventMuseumItemConfig uniquely identified by EventMuseumItemID
//
// Error is only non-nil if the source errors out
func (a *EventMuseumItemConfigAccessor) ByEventMuseumItemID(identifier float64) (EventMuseumItemConfig, error) {
	if a._dataEventMuseumItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMuseumItemConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEventMuseumItemID[identifier], nil
}

// ByMissionID returns the EventMuseumItemConfig uniquely identified by MissionID
//
// Error is only non-nil if the source errors out
func (a *EventMuseumItemConfigAccessor) ByMissionID(identifier float64) (EventMuseumItemConfig, error) {
	if a._dataMissionID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMuseumItemConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionID[identifier], nil
}

// ByMissionStartString returns the EventMuseumItemConfig uniquely identified by MissionStartString
//
// Error is only non-nil if the source errors out
func (a *EventMuseumItemConfigAccessor) ByMissionStartString(identifier string) (EventMuseumItemConfig, error) {
	if a._dataMissionStartString == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMuseumItemConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMissionStartString[identifier], nil
}

// ByMuseumItemID returns the EventMuseumItemConfig uniquely identified by MuseumItemID
//
// Error is only non-nil if the source errors out
func (a *EventMuseumItemConfigAccessor) ByMuseumItemID(identifier float64) (EventMuseumItemConfig, error) {
	if a._dataMuseumItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EventMuseumItemConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMuseumItemID[identifier], nil
}
