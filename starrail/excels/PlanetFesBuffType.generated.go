package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type PlanetFesBuffType struct {
	Decription hash.Hash `json:"Decription"`
	ID         string    `json:"ID"`
	IconPath   string    `json:"IconPath"`
}
type PlanetFesBuffTypeAccessor struct {
	_data         []PlanetFesBuffType
	_dataID       map[string]PlanetFesBuffType
	_dataIconPath map[string]PlanetFesBuffType
}

// LoadData retrieves the data. Must be called before PlanetFesBuffType.GroupData
func (a *PlanetFesBuffTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/PlanetFesBuffType.json")
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
func (a *PlanetFesBuffTypeAccessor) Raw() ([]PlanetFesBuffType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []PlanetFesBuffType{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with PlanetFesBuffTypeAccessor.LoadData to preload everything
func (a *PlanetFesBuffTypeAccessor) GroupData() {
	for _, d := range a._data {
		a._dataID[d.ID] = d
		a._dataIconPath[d.IconPath] = d
	}
}

// ByID returns the PlanetFesBuffType uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *PlanetFesBuffTypeAccessor) ByID(identifier string) (PlanetFesBuffType, error) {
	if a._dataID == nil {
		err := a.LoadData()
		if err != nil {
			return PlanetFesBuffType{}, err
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}

// ByIconPath returns the PlanetFesBuffType uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *PlanetFesBuffTypeAccessor) ByIconPath(identifier string) (PlanetFesBuffType, error) {
	if a._dataIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return PlanetFesBuffType{}, err
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}
