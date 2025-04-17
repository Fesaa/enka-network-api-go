package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MonsterStatusResistanceType struct {
	Icon string `json:"Icon"`
	Type string `json:"Type"`
}
type MonsterStatusResistanceTypeAccessor struct {
	_data     []MonsterStatusResistanceType
	_dataIcon map[string]MonsterStatusResistanceType
	_dataType map[string]MonsterStatusResistanceType
}

// LoadData retrieves the data. Must be called before MonsterStatusResistanceType.GroupData
func (a *MonsterStatusResistanceTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonsterStatusResistanceType.json")
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
func (a *MonsterStatusResistanceTypeAccessor) Raw() ([]MonsterStatusResistanceType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonsterStatusResistanceType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonsterStatusResistanceTypeAccessor.LoadData to preload everything
func (a *MonsterStatusResistanceTypeAccessor) GroupData() {
	a._dataIcon = map[string]MonsterStatusResistanceType{}
	a._dataType = map[string]MonsterStatusResistanceType{}
	for _, d := range a._data {
		a._dataIcon[d.Icon] = d
		a._dataType[d.Type] = d
	}
}

// ByIcon returns the MonsterStatusResistanceType uniquely identified by Icon
//
// Error is only non-nil if the source errors out
func (a *MonsterStatusResistanceTypeAccessor) ByIcon(identifier string) (MonsterStatusResistanceType, error) {
	if a._dataIcon == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonsterStatusResistanceType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataIcon[identifier], nil
}

// ByType returns the MonsterStatusResistanceType uniquely identified by Type
//
// Error is only non-nil if the source errors out
func (a *MonsterStatusResistanceTypeAccessor) ByType(identifier string) (MonsterStatusResistanceType, error) {
	if a._dataType == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonsterStatusResistanceType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataType[identifier], nil
}
