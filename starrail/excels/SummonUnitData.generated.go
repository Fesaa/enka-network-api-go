package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type SummonUnitData struct {
	DestroyOnEnterBattle    bool        `json:"DestroyOnEnterBattle"`
	ID                      json.Number `json:"ID"`
	IsClient                bool        `json:"IsClient"`
	IsTeamSummon            bool        `json:"IsTeamSummon"`
	JsonPath                string      `json:"JsonPath"`
	MaxSummonCount          json.Number `json:"MaxSummonCount"`
	RemoveMazeBuffOnDestroy bool        `json:"RemoveMazeBuffOnDestroy"`
	UniqueGroup             string      `json:"UniqueGroup"`
}
type SummonUnitDataAccessor struct {
	_data         []SummonUnitData
	_dataJsonPath map[string]SummonUnitData
}

// LoadData retrieves the data. Must be called before SummonUnitData.GroupData
func (a *SummonUnitDataAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/SummonUnitData.json")
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
func (a *SummonUnitDataAccessor) Raw() ([]SummonUnitData, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []SummonUnitData{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with SummonUnitDataAccessor.LoadData to preload everything
func (a *SummonUnitDataAccessor) GroupData() {
	a._dataJsonPath = map[string]SummonUnitData{}
	for _, d := range a._data {
		a._dataJsonPath[d.JsonPath] = d
	}
}

// ByJsonPath returns the SummonUnitData uniquely identified by JsonPath
//
// Error is only non-nil if the source errors out
func (a *SummonUnitDataAccessor) ByJsonPath(identifier string) (SummonUnitData, error) {
	if a._dataJsonPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return SummonUnitData{}, err
			}
		}
		a.GroupData()
	}
	return a._dataJsonPath[identifier], nil
}
