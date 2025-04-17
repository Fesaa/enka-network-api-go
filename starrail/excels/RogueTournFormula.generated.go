package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueTournFormula struct {
	FormulaCategory  string      `json:"FormulaCategory"`
	FormulaDisplayID json.Number `json:"FormulaDisplayID"`
	FormulaID        json.Number `json:"FormulaID"`
	FormulaStoryJson string      `json:"FormulaStoryJson"`
	IsInHandbook     bool        `json:"IsInHandbook"`
	MainBuffNum      json.Number `json:"MainBuffNum"`
	MainBuffTypeID   json.Number `json:"MainBuffTypeID"`
	MazeBuffID       json.Number `json:"MazeBuffID"`
	SubBuffNum       json.Number `json:"SubBuffNum"`
	SubBuffTypeID    json.Number `json:"SubBuffTypeID"`
	TournMode        string      `json:"TournMode"`
}
type RogueTournFormulaAccessor struct {
	_data          []RogueTournFormula
	_dataFormulaID map[json.Number]RogueTournFormula
}

// LoadData retrieves the data. Must be called before RogueTournFormula.GroupData
func (a *RogueTournFormulaAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournFormula.json")
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
func (a *RogueTournFormulaAccessor) Raw() ([]RogueTournFormula, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournFormula{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournFormulaAccessor.LoadData to preload everything
func (a *RogueTournFormulaAccessor) GroupData() {
	a._dataFormulaID = map[json.Number]RogueTournFormula{}
	for _, d := range a._data {
		a._dataFormulaID[d.FormulaID] = d
	}
}

// ByFormulaID returns the RogueTournFormula uniquely identified by FormulaID
//
// Error is only non-nil if the source errors out
func (a *RogueTournFormulaAccessor) ByFormulaID(identifier json.Number) (RogueTournFormula, error) {
	if a._dataFormulaID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournFormula{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFormulaID[identifier], nil
}
