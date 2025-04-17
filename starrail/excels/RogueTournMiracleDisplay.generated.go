package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueTournMiracleDisplay struct {
	MiracleBGDesc         map[string]json.Number `json:"MiracleBGDesc"`
	MiracleDisplayID      json.Number            `json:"MiracleDisplayID"`
	MiracleFigureIconPath string                 `json:"MiracleFigureIconPath"`
	MiracleIconPath       string                 `json:"MiracleIconPath"`
	MiracleName           map[string]json.Number `json:"MiracleName"`
}
type RogueTournMiracleDisplayAccessor struct {
	_data                 []RogueTournMiracleDisplay
	_dataMiracleDisplayID map[json.Number]RogueTournMiracleDisplay
}

// LoadData retrieves the data. Must be called before RogueTournMiracleDisplay.GroupData
func (a *RogueTournMiracleDisplayAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournMiracleDisplay.json")
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
func (a *RogueTournMiracleDisplayAccessor) Raw() ([]RogueTournMiracleDisplay, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournMiracleDisplay{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournMiracleDisplayAccessor.LoadData to preload everything
func (a *RogueTournMiracleDisplayAccessor) GroupData() {
	a._dataMiracleDisplayID = map[json.Number]RogueTournMiracleDisplay{}
	for _, d := range a._data {
		a._dataMiracleDisplayID[d.MiracleDisplayID] = d
	}
}

// ByMiracleDisplayID returns the RogueTournMiracleDisplay uniquely identified by MiracleDisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueTournMiracleDisplayAccessor) ByMiracleDisplayID(identifier json.Number) (RogueTournMiracleDisplay, error) {
	if a._dataMiracleDisplayID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournMiracleDisplay{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiracleDisplayID[identifier], nil
}
