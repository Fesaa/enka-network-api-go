package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlayerReturnJourneyItem struct {
	ActivityModuleID float64   `json:"ActivityModuleID"`
	BgPath           string    `json:"BgPath"`
	ExtraDesc        hash.Hash `json:"ExtraDesc"`
	ID               float64   `json:"ID"`
	IsHideInBeta     bool      `json:"IsHideInBeta"`
	Name             hash.Hash `json:"Name"`
	Sort             float64   `json:"Sort"`
	Title            hash.Hash `json:"Title"`
	Type             string    `json:"Type"`
}
type PlayerReturnJourneyItemAccessor struct {
	_data       []PlayerReturnJourneyItem
	_dataBgPath map[string]PlayerReturnJourneyItem
	_dataID     map[float64]PlayerReturnJourneyItem
	_dataSort   map[float64]PlayerReturnJourneyItem
	_dataType   map[string]PlayerReturnJourneyItem
}

// LoadData retrieves the data. Must be called before PlayerReturnJourneyItem.GroupData
func (a *PlayerReturnJourneyItemAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlayerReturnJourneyItem.json")
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
func (a *PlayerReturnJourneyItemAccessor) Raw() ([]PlayerReturnJourneyItem, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlayerReturnJourneyItem{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlayerReturnJourneyItemAccessor.LoadData to preload everything
func (a *PlayerReturnJourneyItemAccessor) GroupData() {
	a._dataBgPath = map[string]PlayerReturnJourneyItem{}
	a._dataID = map[float64]PlayerReturnJourneyItem{}
	a._dataSort = map[float64]PlayerReturnJourneyItem{}
	a._dataType = map[string]PlayerReturnJourneyItem{}
	for _, d := range a._data {
		a._dataBgPath[d.BgPath] = d
		a._dataID[d.ID] = d
		a._dataSort[d.Sort] = d
		a._dataType[d.Type] = d
	}
}

// ByBgPath returns the PlayerReturnJourneyItem uniquely identified by BgPath
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnJourneyItemAccessor) ByBgPath(identifier string) (PlayerReturnJourneyItem, error) {
	if a._dataBgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnJourneyItem{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBgPath[identifier], nil
}

// ByID returns the PlayerReturnJourneyItem uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnJourneyItemAccessor) ByID(identifier float64) (PlayerReturnJourneyItem, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnJourneyItem{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// BySort returns the PlayerReturnJourneyItem uniquely identified by Sort
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnJourneyItemAccessor) BySort(identifier float64) (PlayerReturnJourneyItem, error) {
	if a._dataSort == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnJourneyItem{}, err
			}
		}
		a.GroupData()
	}
	return a._dataSort[identifier], nil
}

// ByType returns the PlayerReturnJourneyItem uniquely identified by Type
//
// Error is only non-nil if the source errors out
func (a *PlayerReturnJourneyItemAccessor) ByType(identifier string) (PlayerReturnJourneyItem, error) {
	if a._dataType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlayerReturnJourneyItem{}, err
			}
		}
		a.GroupData()
	}
	return a._dataType[identifier], nil
}
