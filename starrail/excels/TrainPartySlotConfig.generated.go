package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartySlotConfig struct {
	CameraStaticID float64   `json:"CameraStaticID"`
	ID             float64   `json:"ID"`
	Name           hash.Hash `json:"Name"`
	SortID         float64   `json:"SortID"`
	TagList        []float64 `json:"TagList"`
}
type TrainPartySlotConfigAccessor struct {
	_data   []TrainPartySlotConfig
	_dataID map[float64]TrainPartySlotConfig
}

// LoadData retrieves the data. Must be called before TrainPartySlotConfig.GroupData
func (a *TrainPartySlotConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartySlotConfig.json")
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
func (a *TrainPartySlotConfigAccessor) Raw() ([]TrainPartySlotConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartySlotConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartySlotConfigAccessor.LoadData to preload everything
func (a *TrainPartySlotConfigAccessor) GroupData() {
	a._dataID = map[float64]TrainPartySlotConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TrainPartySlotConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TrainPartySlotConfigAccessor) ByID(identifier float64) (TrainPartySlotConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartySlotConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
