package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type PlanetFesGameGachaSymbol struct {
	ID       json.Number `json:"ID"`
	IconPath string      `json:"IconPath"`
}
type PlanetFesGameGachaSymbolAccessor struct {
	_data         []PlanetFesGameGachaSymbol
	_dataID       map[json.Number]PlanetFesGameGachaSymbol
	_dataIconPath map[string]PlanetFesGameGachaSymbol
}

// LoadData retrieves the data. Must be called before PlanetFesGameGachaSymbol.GroupData
func (a *PlanetFesGameGachaSymbolAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesGameGachaSymbol.json")
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
func (a *PlanetFesGameGachaSymbolAccessor) Raw() ([]PlanetFesGameGachaSymbol, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesGameGachaSymbol{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesGameGachaSymbolAccessor.LoadData to preload everything
func (a *PlanetFesGameGachaSymbolAccessor) GroupData() {
	a._dataID = map[json.Number]PlanetFesGameGachaSymbol{}
	a._dataIconPath = map[string]PlanetFesGameGachaSymbol{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
	}
}

// ByID returns the PlanetFesGameGachaSymbol uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesGameGachaSymbolAccessor) ByID(identifier json.Number) (PlanetFesGameGachaSymbol, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesGameGachaSymbol{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the PlanetFesGameGachaSymbol uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *PlanetFesGameGachaSymbolAccessor) ByIconPath(identifier string) (PlanetFesGameGachaSymbol, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesGameGachaSymbol{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
