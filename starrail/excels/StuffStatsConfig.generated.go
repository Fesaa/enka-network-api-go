package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type StuffStatsConfig struct {
	MuseumStatsName hash.Hash `json:"MuseumStatsName"`
	StatsID         string    `json:"StatsID"`
	StatsIconPath   string    `json:"StatsIconPath"`
}
type StuffStatsConfigAccessor struct {
	_data              []StuffStatsConfig
	_dataStatsID       map[string]StuffStatsConfig
	_dataStatsIconPath map[string]StuffStatsConfig
}

// LoadData retrieves the data. Must be called before StuffStatsConfig.GroupData
func (a *StuffStatsConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/StuffStatsConfig.json")
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
func (a *StuffStatsConfigAccessor) Raw() ([]StuffStatsConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []StuffStatsConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with StuffStatsConfigAccessor.LoadData to preload everything
func (a *StuffStatsConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataStatsID[d.StatsID] = d
		a._dataStatsIconPath[d.StatsIconPath] = d
	}
}

// ByStatsID returns the StuffStatsConfig uniquely identified by StatsID
//
// Error is only non-nil if the source errors out
func (a *StuffStatsConfigAccessor) ByStatsID(identifier string) (StuffStatsConfig, error) {
	if a._dataStatsID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StuffStatsConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStatsID[identifier], nil
}

// ByStatsIconPath returns the StuffStatsConfig uniquely identified by StatsIconPath
//
// Error is only non-nil if the source errors out
func (a *StuffStatsConfigAccessor) ByStatsIconPath(identifier string) (StuffStatsConfig, error) {
	if a._dataStatsIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return StuffStatsConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataStatsIconPath[identifier], nil
}
