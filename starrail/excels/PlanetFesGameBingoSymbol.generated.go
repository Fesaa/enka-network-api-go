package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesGameBingoSymbol struct {
	ID       float64 `json:"ID"`
	IconPath string  `json:"IconPath"`
}
type PlanetFesGameBingoSymbolAccessor struct {
	_data         []PlanetFesGameBingoSymbol
	_dataID       map[float64]PlanetFesGameBingoSymbol
	_dataIconPath map[string]PlanetFesGameBingoSymbol
}

// LoadData retrieves the data. Must be called before PlanetFesGameBingoSymbol.GroupData
func (a *PlanetFesGameBingoSymbolAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesGameBingoSymbol.json")
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
func (a *PlanetFesGameBingoSymbolAccessor) Raw() ([]PlanetFesGameBingoSymbol, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesGameBingoSymbol{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesGameBingoSymbolAccessor.LoadData to preload everything
func (a *PlanetFesGameBingoSymbolAccessor) GroupData() {
	a._dataID = map[float64]PlanetFesGameBingoSymbol{}
	a._dataIconPath = map[string]PlanetFesGameBingoSymbol{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
	}
}

// ByID returns the PlanetFesGameBingoSymbol uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesGameBingoSymbolAccessor) ByID(identifier float64) (PlanetFesGameBingoSymbol, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesGameBingoSymbol{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the PlanetFesGameBingoSymbol uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *PlanetFesGameBingoSymbolAccessor) ByIconPath(identifier string) (PlanetFesGameBingoSymbol, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesGameBingoSymbol{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
