package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type AreaMapConfig struct {
	Desc               hash.Hash `json:"Desc"`
	ID                 float64   `json:"ID"`
	IsUnlockAfterEnter bool      `json:"IsUnlockAfterEnter"`
	MapSpaceType       string    `json:"MapSpaceType"`
	MenuIconID         float64   `json:"MenuIconID"`
	MenuSortID         float64   `json:"MenuSortID"`
	Name               hash.Hash `json:"Name"`
}
type AreaMapConfigAccessor struct {
	_data   []AreaMapConfig
	_dataID map[float64]AreaMapConfig
}

// LoadData retrieves the data. Must be called before AreaMapConfig.GroupData
func (a *AreaMapConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/AreaMapConfig.json")
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
func (a *AreaMapConfigAccessor) Raw() ([]AreaMapConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []AreaMapConfig{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with AreaMapConfigAccessor.LoadData to preload everything
func (a *AreaMapConfigAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the AreaMapConfig uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *AreaMapConfigAccessor) ByID(identifier float64) (AreaMapConfig, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return AreaMapConfig{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
