package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type MenuItemName struct {
	ID     float64   `json:"ID"`
	TextID hash.Hash `json:"TextID"`
}
type MenuItemNameAccessor struct {
	_data   []MenuItemName
	_dataID map[float64]MenuItemName
}

// LoadData retrieves the data. Must be called before MenuItemName.GroupData
func (a *MenuItemNameAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MenuItemName.json")
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
func (a *MenuItemNameAccessor) Raw() ([]MenuItemName, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MenuItemName{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MenuItemNameAccessor.LoadData to preload everything
func (a *MenuItemNameAccessor) GroupData() {
	a._dataID = map[float64]MenuItemName{}
	for _, d := range a._data {
		a._dataID[d.ID] = d
	}
}

// ByID returns the MenuItemName uniquely identified by ID
//
// Error is only non-nil if the source errors out
func (a *MenuItemNameAccessor) ByID(identifier float64) (MenuItemName, error) {
	if a._dataID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MenuItemName{}, err
			}
		}
		a.GroupData()
	}
	return a._dataID[identifier], nil
}
