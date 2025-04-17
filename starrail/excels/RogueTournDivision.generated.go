package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueTournDivision struct {
	DivisionHintDesc       map[string]json.Number `json:"DivisionHintDesc"`
	DivisionIconPath       string                 `json:"DivisionIconPath"`
	DivisionIconPrefabPath string                 `json:"DivisionIconPrefabPath"`
	DivisionLevel          json.Number            `json:"DivisionLevel"`
	DivisionName           map[string]json.Number `json:"DivisionName"`
	DivisionProgress       json.Number            `json:"DivisionProgress"`
	DivisionSmallIconPath  string                 `json:"DivisionSmallIconPath"`
}
type RogueTournDivisionAccessor struct {
	_data                       []RogueTournDivision
	_dataDivisionIconPath       map[string]RogueTournDivision
	_dataDivisionIconPrefabPath map[string]RogueTournDivision
	_dataDivisionSmallIconPath  map[string]RogueTournDivision
}

// LoadData retrieves the data. Must be called before RogueTournDivision.GroupData
func (a *RogueTournDivisionAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournDivision.json")
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
func (a *RogueTournDivisionAccessor) Raw() ([]RogueTournDivision, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournDivision{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournDivisionAccessor.LoadData to preload everything
func (a *RogueTournDivisionAccessor) GroupData() {
	a._dataDivisionIconPath = map[string]RogueTournDivision{}
	a._dataDivisionIconPrefabPath = map[string]RogueTournDivision{}
	a._dataDivisionSmallIconPath = map[string]RogueTournDivision{}
	for _, d := range a._data {
		a._dataDivisionIconPath[d.DivisionIconPath] = d
		a._dataDivisionIconPrefabPath[d.DivisionIconPrefabPath] = d
		a._dataDivisionSmallIconPath[d.DivisionSmallIconPath] = d
	}
}

// ByDivisionIconPath returns the RogueTournDivision uniquely identified by DivisionIconPath
//
// Error is only non-nil if the source errors out
func (a *RogueTournDivisionAccessor) ByDivisionIconPath(identifier string) (RogueTournDivision, error) {
	if a._dataDivisionIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournDivision{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDivisionIconPath[identifier], nil
}

// ByDivisionIconPrefabPath returns the RogueTournDivision uniquely identified by DivisionIconPrefabPath
//
// Error is only non-nil if the source errors out
func (a *RogueTournDivisionAccessor) ByDivisionIconPrefabPath(identifier string) (RogueTournDivision, error) {
	if a._dataDivisionIconPrefabPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournDivision{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDivisionIconPrefabPath[identifier], nil
}

// ByDivisionSmallIconPath returns the RogueTournDivision uniquely identified by DivisionSmallIconPath
//
// Error is only non-nil if the source errors out
func (a *RogueTournDivisionAccessor) ByDivisionSmallIconPath(identifier string) (RogueTournDivision, error) {
	if a._dataDivisionSmallIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournDivision{}, err
			}
		}
		a.GroupData()
	}
	return a._dataDivisionSmallIconPath[identifier], nil
}
