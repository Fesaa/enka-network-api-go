package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"encoding/json"
	"net/http"
	"io"
)

type ShopItemGroupConfig struct {
	GroupID     json.Number `json:"GroupID"`
	GroupType   string      `json:"GroupType"`
	ItemID      json.Number `json:"ItemID"`
	RotateOrder json.Number `json:"RotateOrder"`
}
type ShopItemGroupConfigAccessor struct {
	_data       []ShopItemGroupConfig
	_dataItemID map[json.Number]ShopItemGroupConfig
}

// LoadData retrieves the data. Must be called before ShopItemGroupConfig.GroupData
func (a *ShopItemGroupConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/ShopItemGroupConfig.json")
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
func (a *ShopItemGroupConfigAccessor) Raw() ([]ShopItemGroupConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []ShopItemGroupConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with ShopItemGroupConfigAccessor.LoadData to preload everything
func (a *ShopItemGroupConfigAccessor) GroupData() {
	a._dataItemID = map[json.Number]ShopItemGroupConfig{}
	for _, d := range a._data {
		a._dataItemID[d.ItemID] = d
	}
}

// ByItemID returns the ShopItemGroupConfig uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *ShopItemGroupConfigAccessor) ByItemID(identifier json.Number) (ShopItemGroupConfig, error) {
	if a._dataItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return ShopItemGroupConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}
