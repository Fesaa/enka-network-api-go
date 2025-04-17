package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesSummary struct {
	Description hash.Hash `json:"Description"`
	ID          string    `json:"ID"`
	Name        hash.Hash `json:"Name"`
	TargetNum   float64   `json:"TargetNum"`
}
type PlanetFesSummaryAccessor struct {
	_data   []PlanetFesSummary
	_dataID map[string]PlanetFesSummary
}

// LoadData retrieves the data. Must be called before PlanetFesSummary.GroupData
func (a *PlanetFesSummaryAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesSummary.json")
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
func (a *PlanetFesSummaryAccessor) Raw() ([]PlanetFesSummary, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesSummary{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesSummaryAccessor.LoadData to preload everything
func (a *PlanetFesSummaryAccessor) GroupData() {
	a._dataID = map[string]PlanetFesSummary{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PlanetFesSummary uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesSummaryAccessor) ByID(identifier string) (PlanetFesSummary, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesSummary{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
