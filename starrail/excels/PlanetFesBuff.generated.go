package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesBuff struct {
	Duration  float64   `json:"Duration"`
	ID        float64   `json:"ID"`
	SourceID  float64   `json:"SourceID"`
	Type      string    `json:"Type"`
	TypeParam []float64 `json:"TypeParam"`
}
type PlanetFesBuffAccessor struct {
	_data   []PlanetFesBuff
	_dataID map[float64]PlanetFesBuff
}

// LoadData retrieves the data. Must be called before PlanetFesBuff.GroupData
func (a *PlanetFesBuffAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesBuff.json")
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
func (a *PlanetFesBuffAccessor) Raw() ([]PlanetFesBuff, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesBuff{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesBuffAccessor.LoadData to preload everything
func (a *PlanetFesBuffAccessor) GroupData() {
	a._dataID = map[float64]PlanetFesBuff{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PlanetFesBuff uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesBuffAccessor) ByID(identifier float64) (PlanetFesBuff, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesBuff{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
