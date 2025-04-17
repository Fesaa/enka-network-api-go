package // Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"io"
	"net/http"
)

type RogueTournGambleUnit struct {
	GambleUnitID    json.Number `json:"GambleUnitID"`
	GambleUnitIcon  string      `json:"GambleUnitIcon"`
	GambleUnitParam json.Number `json:"GambleUnitParam"`
	GambleUnitType  string      `json:"GambleUnitType"`
}
type RogueTournGambleUnitAccessor struct {
	_data []RogueTournGambleUnit
}

// LoadData retrieves the data. Must be called before RogueTournGambleUnit.GroupData
func (a *RogueTournGambleUnitAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournGambleUnit.json")
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
func (a *RogueTournGambleUnitAccessor) Raw() ([]RogueTournGambleUnit, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournGambleUnit{}, err
		}
	}
	return a._data, nil
}
