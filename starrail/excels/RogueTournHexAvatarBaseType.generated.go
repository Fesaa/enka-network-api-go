package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RogueTournHexAvatarBaseType struct {
	AvatarDamageType []string    `json:"AvatarDamageType"`
	AvatarType       []string    `json:"AvatarType"`
	MiracleID        json.Number `json:"MiracleID"`
}
type RogueTournHexAvatarBaseTypeAccessor struct {
	_data          []RogueTournHexAvatarBaseType
	_dataMiracleID map[json.Number]RogueTournHexAvatarBaseType
}

// LoadData retrieves the data. Must be called before RogueTournHexAvatarBaseType.GroupData
func (a *RogueTournHexAvatarBaseTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueTournHexAvatarBaseType.json")
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
func (a *RogueTournHexAvatarBaseTypeAccessor) Raw() ([]RogueTournHexAvatarBaseType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueTournHexAvatarBaseType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueTournHexAvatarBaseTypeAccessor.LoadData to preload everything
func (a *RogueTournHexAvatarBaseTypeAccessor) GroupData() {
	a._dataMiracleID = map[json.Number]RogueTournHexAvatarBaseType{}
	for _, d := range a._data {
		a._dataMiracleID[d.MiracleID] = d
	}
}

// ByMiracleID returns the RogueTournHexAvatarBaseType uniquely identified by MiracleID
//
// Error is only non-nil if the source errors out
func (a *RogueTournHexAvatarBaseTypeAccessor) ByMiracleID(identifier json.Number) (RogueTournHexAvatarBaseType, error) {
	if a._dataMiracleID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RogueTournHexAvatarBaseType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMiracleID[identifier], nil
}
