package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesCardTheme struct {
	CardIDList []float64 `json:"CardIDList"`
	IconPath   string    `json:"IconPath"`
	Name       hash.Hash `json:"Name"`
	ThemeID    float64   `json:"ThemeID"`
}
type PlanetFesCardThemeAccessor struct {
	_data         []PlanetFesCardTheme
	_dataIconPath map[string]PlanetFesCardTheme
	_dataThemeID  map[float64]PlanetFesCardTheme
}

// LoadData retrieves the data. Must be called before PlanetFesCardTheme.GroupData
func (a *PlanetFesCardThemeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesCardTheme.json")
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
func (a *PlanetFesCardThemeAccessor) Raw() ([]PlanetFesCardTheme, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesCardTheme{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesCardThemeAccessor.LoadData to preload everything
func (a *PlanetFesCardThemeAccessor) GroupData() {
	a._dataIconPath = map[string]PlanetFesCardTheme{}
	a._dataThemeID = map[float64]PlanetFesCardTheme{}
	for _, d := range a._data {
		a._dataIconPath[d.IconPath] = d
		a._dataThemeID[d.ThemeID] = d
	}
}

// ByIconPath returns the PlanetFesCardTheme uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *PlanetFesCardThemeAccessor) ByIconPath(identifier string) (PlanetFesCardTheme, error) {
	if a._dataIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesCardTheme{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByThemeID returns the PlanetFesCardTheme uniquely identified by ThemeID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesCardThemeAccessor) ByThemeID(identifier float64) (PlanetFesCardTheme, error) {
	if a._dataThemeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return PlanetFesCardTheme{}, err
			}
		}
		a.GroupData()
	}
	return a._dataThemeID[identifier], nil
}
