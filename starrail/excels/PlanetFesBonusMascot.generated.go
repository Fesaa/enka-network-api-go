package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesBonusMascot struct {
	ID         float64 `json:"ID"`
	SourcePath string  `json:"SourcePath"`
}
type PlanetFesBonusMascotAccessor struct {
	_data   []PlanetFesBonusMascot
	_dataID map[float64]PlanetFesBonusMascot
}

// LoadData retrieves the data. Must be called before PlanetFesBonusMascot.GroupData
func (a *PlanetFesBonusMascotAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesBonusMascot.json")
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
func (a *PlanetFesBonusMascotAccessor) Raw() ([]PlanetFesBonusMascot, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesBonusMascot{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesBonusMascotAccessor.LoadData to preload everything
func (a *PlanetFesBonusMascotAccessor) GroupData() {
	a._dataID = map[float64]PlanetFesBonusMascot{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the PlanetFesBonusMascot uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesBonusMascotAccessor) ByID(identifier float64) (PlanetFesBonusMascot, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesBonusMascot{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
