package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type EquipmentAtlas struct {
	DefaultUnlock bool    `json:"DefaultUnlock"`
	EquipmentID   float64 `json:"EquipmentID"`
}
type EquipmentAtlasAccessor struct {
	_data            []EquipmentAtlas
	_dataEquipmentID map[float64]EquipmentAtlas
}

// LoadData retrieves the data. Must be called before EquipmentAtlas.GroupData
func (a *EquipmentAtlasAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/EquipmentAtlas.json")
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
func (a *EquipmentAtlasAccessor) Raw() ([]EquipmentAtlas, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []EquipmentAtlas{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with EquipmentAtlasAccessor.LoadData to preload everything
func (a *EquipmentAtlasAccessor) GroupData() {
	for _, d := range a._data {
		a._dataEquipmentID[d.EquipmentID] = d
	}
}

// ByEquipmentID returns the EquipmentAtlas uniquely identified by EquipmentID
//
// Error is only non-nil if the source errors out
func (a *EquipmentAtlasAccessor) ByEquipmentID(identifier float64) (EquipmentAtlas, error) {
	if a._dataEquipmentID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return EquipmentAtlas{}, err
			}
		}
		a.GroupData()
	}
	return a._dataEquipmentID[identifier], nil
}
