package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type TrainPartyStepConfig struct {
	CoinCost         float64   `json:"CoinCost"`
	GroupID          float64   `json:"GroupID"`
	HasCutScene      bool      `json:"HasCutScene"`
	HasPreview       bool      `json:"HasPreview"`
	ID               float64   `json:"ID"`
	ImgPath          string    `json:"ImgPath"`
	Name             hash.Hash `json:"Name"`
	SortID           float64   `json:"SortID"`
	StaticPropIDList []float64 `json:"StaticPropIDList"`
}
type TrainPartyStepConfigAccessor struct {
	_data   []TrainPartyStepConfig
	_dataID map[float64]TrainPartyStepConfig
}

// LoadData retrieves the data. Must be called before TrainPartyStepConfig.GroupData
func (a *TrainPartyStepConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/TrainPartyStepConfig.json")
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
func (a *TrainPartyStepConfigAccessor) Raw() ([]TrainPartyStepConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []TrainPartyStepConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with TrainPartyStepConfigAccessor.LoadData to preload everything
func (a *TrainPartyStepConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the TrainPartyStepConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *TrainPartyStepConfigAccessor) ByID(identifier float64) (TrainPartyStepConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return TrainPartyStepConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
