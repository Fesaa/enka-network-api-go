package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AtlasConfig struct {
	ID   float64   `json:"ID"`
	Name hash.Hash `json:"Name"`
}
type AtlasConfigAccessor struct {
	_data   []AtlasConfig
	_dataID map[float64]AtlasConfig
}

// LoadData retrieves the data. Must be called before AtlasConfig.GroupData
func (a *AtlasConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AtlasConfig.json")
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
func (a *AtlasConfigAccessor) Raw() ([]AtlasConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AtlasConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AtlasConfigAccessor.LoadData to preload everything
func (a *AtlasConfigAccessor) GroupData() {
	a._dataID = map[float64]AtlasConfig{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the AtlasConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AtlasConfigAccessor) ByID(identifier float64) (AtlasConfig, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return AtlasConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
