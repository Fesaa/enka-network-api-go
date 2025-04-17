package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type RelicBaseType struct {
	BaseTypeIconPath  string                 `json:"BaseTypeIconPath"`
	BaseTypeText      map[string]json.Number `json:"BaseTypeText"`
	Type              string                 `json:"Type"`
	ValidPropertyList []string               `json:"ValidPropertyList"`
}
type RelicBaseTypeAccessor struct {
	_data                 []RelicBaseType
	_dataBaseTypeIconPath map[string]RelicBaseType
}

// LoadData retrieves the data. Must be called before RelicBaseType.GroupData
func (a *RelicBaseTypeAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/RelicBaseType.json")
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
func (a *RelicBaseTypeAccessor) Raw() ([]RelicBaseType, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []RelicBaseType{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with RelicBaseTypeAccessor.LoadData to preload everything
func (a *RelicBaseTypeAccessor) GroupData() {
	a._dataBaseTypeIconPath = map[string]RelicBaseType{}
	for _, d := range a._data {
		a._dataBaseTypeIconPath[d.BaseTypeIconPath] = d
	}
}

// ByBaseTypeIconPath returns the RelicBaseType uniquely identified by BaseTypeIconPath
//
// Error is only non-nil if the source errors out
func (a *RelicBaseTypeAccessor) ByBaseTypeIconPath(identifier string) (RelicBaseType, error) {
	if a._dataBaseTypeIconPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return RelicBaseType{}, err
			}
		}
		a.GroupData()
	}
	return a._dataBaseTypeIconPath[identifier], nil
}
