package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyBuffConfig struct {
	BuffDesc    hash.Hash `json:"BuffDesc"`
	BuffID      float64   `json:"BuffID"`
	BuffName    hash.Hash `json:"BuffName"`
	Duration    float64   `json:"Duration"`
	EffectID    float64   `json:"EffectID"`
	IconPath    string    `json:"IconPath"`
	IsPermanent bool      `json:"IsPermanent"`
	Rank        float64   `json:"Rank"`
}
type MonopolyBuffConfigAccessor struct {
	_data       []MonopolyBuffConfig
	_dataBuffID map[float64]MonopolyBuffConfig
}

// LoadData retrieves the data. Must be called before MonopolyBuffConfig.GroupData
func (a *MonopolyBuffConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyBuffConfig.json")
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
func (a *MonopolyBuffConfigAccessor) Raw() ([]MonopolyBuffConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyBuffConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyBuffConfigAccessor.LoadData to preload everything
func (a *MonopolyBuffConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataBuffID[d.BuffID] = d
	}
}

// ByBuffID returns the MonopolyBuffConfig uniquely identified by BuffID
//
// Error is only non-nil if the source errors out
func (a *MonopolyBuffConfigAccessor) ByBuffID(identifier float64) (MonopolyBuffConfig, error) {
	if a._dataBuffID == nil {
		err := a.LoadData()
		if err != nil {
			return MonopolyBuffConfig{}, err
		}
		a.GroupData()
	}
	return a._dataBuffID[identifier], nil
}
