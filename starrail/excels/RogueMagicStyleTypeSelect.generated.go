package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type RogueMagicStyleTypeSelect struct {
	DisplayID float64   `json:"DisplayID"`
	EnumDesc  hash.Hash `json:"EnumDesc"`
	EnumType  string    `json:"EnumType"`
	IconPath  string    `json:"IconPath"`
	UnlockID  float64   `json:"UnlockID"`
}
type RogueMagicStyleTypeSelectAccessor struct {
	_data          []RogueMagicStyleTypeSelect
	_dataDisplayID map[float64]RogueMagicStyleTypeSelect
	_dataIconPath  map[string]RogueMagicStyleTypeSelect
	_dataEnumType  map[string]RogueMagicStyleTypeSelect
}

// LoadData retrieves the data. Must be called before RogueMagicStyleTypeSelect.GroupData
func (a *RogueMagicStyleTypeSelectAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RogueMagicStyleTypeSelect.json")
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
func (a *RogueMagicStyleTypeSelectAccessor) Raw() ([]RogueMagicStyleTypeSelect, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RogueMagicStyleTypeSelect{}, err
		}
		a.GroupData()
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RogueMagicStyleTypeSelectAccessor.LoadData to preload everything
func (a *RogueMagicStyleTypeSelectAccessor) GroupData() {
	for _, d := range a._data {
		a._dataDisplayID[d.DisplayID] = d
		a._dataIconPath[d.IconPath] = d
		a._dataEnumType[d.EnumType] = d
	}
}

// ByDisplayID returns the RogueMagicStyleTypeSelect uniquely identified by DisplayID
//
// Error is only non-nil if the source errors out
func (a *RogueMagicStyleTypeSelectAccessor) ByDisplayID(identifier float64) (RogueMagicStyleTypeSelect, error) {
	if a._dataDisplayID == nil {
		err := a.LoadData()
		if err != nil {
			return RogueMagicStyleTypeSelect{}, err
		}
		a.GroupData()
	}
	return a._dataDisplayID[identifier], nil
}

// ByIconPath returns the RogueMagicStyleTypeSelect uniquely identified by IconPath
//
// Error is only non-nil if the source errors out
func (a *RogueMagicStyleTypeSelectAccessor) ByIconPath(identifier string) (RogueMagicStyleTypeSelect, error) {
	if a._dataIconPath == nil {
		err := a.LoadData()
		if err != nil {
			return RogueMagicStyleTypeSelect{}, err
		}
		a.GroupData()
	}
	return a._dataIconPath[identifier], nil
}

// ByEnumType returns the RogueMagicStyleTypeSelect uniquely identified by EnumType
//
// Error is only non-nil if the source errors out
func (a *RogueMagicStyleTypeSelectAccessor) ByEnumType(identifier string) (RogueMagicStyleTypeSelect, error) {
	if a._dataEnumType == nil {
		err := a.LoadData()
		if err != nil {
			return RogueMagicStyleTypeSelect{}, err
		}
		a.GroupData()
	}
	return a._dataEnumType[identifier], nil
}
