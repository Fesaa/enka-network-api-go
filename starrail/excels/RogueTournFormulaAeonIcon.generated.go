package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type RogueTournFormulaAeonIcon struct {
	BuffTypeID           float64 `json:"BuffTypeID"`
	FormulaIcon          string  `json:"FormulaIcon"`
	FormulaSubIcon       string  `json:"FormulaSubIcon"`
	UltraFormulaCardIcon string  `json:"UltraFormulaCardIcon"`
	UltraFormulaIcon     string  `json:"UltraFormulaIcon"`
}
type RogueTournFormulaAeonIconAccessor struct {
	_data                     []RogueTournFormulaAeonIcon
	_dataBuffTypeID           map[float64]RogueTournFormulaAeonIcon
	_dataFormulaIcon          map[string]RogueTournFormulaAeonIcon
	_dataFormulaSubIcon       map[string]RogueTournFormulaAeonIcon
	_dataUltraFormulaCardIcon map[string]RogueTournFormulaAeonIcon
	_dataUltraFormulaIcon     map[string]RogueTournFormulaAeonIcon
}

// LoadData retrieves the data. Must be called before RogueTournFormulaAeonIcon.GroupData
func (a *RogueTournFormulaAeonIconAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournFormulaAeonIcon.json")
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
func (a *RogueTournFormulaAeonIconAccessor) Raw() ([]RogueTournFormulaAeonIcon, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournFormulaAeonIcon{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournFormulaAeonIconAccessor.LoadData to preload everything
func (a *RogueTournFormulaAeonIconAccessor) GroupData() {
	a._dataBuffTypeID = map[float64]RogueTournFormulaAeonIcon{}
	a._dataFormulaIcon = map[string]RogueTournFormulaAeonIcon{}
	a._dataFormulaSubIcon = map[string]RogueTournFormulaAeonIcon{}
	a._dataUltraFormulaCardIcon = map[string]RogueTournFormulaAeonIcon{}
	a._dataUltraFormulaIcon = map[string]RogueTournFormulaAeonIcon{}
	for _, d := range a._data {
		a._dataBuffTypeID[d.BuffTypeID] = d
		a._dataFormulaIcon[d.FormulaIcon] = d
		a._dataFormulaSubIcon[d.FormulaSubIcon] = d
		a._dataUltraFormulaCardIcon[d.UltraFormulaCardIcon] = d
		a._dataUltraFormulaIcon[d.UltraFormulaIcon] = d
	}
}

// ByBuffTypeID returns the RogueTournFormulaAeonIcon uniquely identified by BuffTypeID
//
// Error is only non-nil if the source errors out
func (a *RogueTournFormulaAeonIconAccessor) ByBuffTypeID(identifier float64) (RogueTournFormulaAeonIcon, error) {
	if a._dataBuffTypeID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournFormulaAeonIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBuffTypeID[identifier], nil
}

// ByFormulaIcon returns the RogueTournFormulaAeonIcon uniquely identified by FormulaIcon
//
// Error is only non-nil if the source errors out
func (a *RogueTournFormulaAeonIconAccessor) ByFormulaIcon(identifier string) (RogueTournFormulaAeonIcon, error) {
	if a._dataFormulaIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournFormulaAeonIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFormulaIcon[identifier], nil
}

// ByFormulaSubIcon returns the RogueTournFormulaAeonIcon uniquely identified by FormulaSubIcon
//
// Error is only non-nil if the source errors out
func (a *RogueTournFormulaAeonIconAccessor) ByFormulaSubIcon(identifier string) (RogueTournFormulaAeonIcon, error) {
	if a._dataFormulaSubIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournFormulaAeonIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataFormulaSubIcon[identifier], nil
}

// ByUltraFormulaCardIcon returns the RogueTournFormulaAeonIcon uniquely identified by UltraFormulaCardIcon
//
// Error is only non-nil if the source errors out
func (a *RogueTournFormulaAeonIconAccessor) ByUltraFormulaCardIcon(identifier string) (RogueTournFormulaAeonIcon, error) {
	if a._dataUltraFormulaCardIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournFormulaAeonIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUltraFormulaCardIcon[identifier], nil
}

// ByUltraFormulaIcon returns the RogueTournFormulaAeonIcon uniquely identified by UltraFormulaIcon
//
// Error is only non-nil if the source errors out
func (a *RogueTournFormulaAeonIconAccessor) ByUltraFormulaIcon(identifier string) (RogueTournFormulaAeonIcon, error) {
	if a._dataUltraFormulaIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournFormulaAeonIcon{}, err
			}
		}
		a.GroupData()
	}
	return a._dataUltraFormulaIcon[identifier], nil
}
