package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type ShopItemGroupConfig struct {
	GroupID     float64 `json:"GroupID"`
	GroupType   string  `json:"GroupType"`
	ItemID      float64 `json:"ItemID"`
	RotateOrder float64 `json:"RotateOrder"`
}
type ShopItemGroupConfigAccessor struct {
	_data       []ShopItemGroupConfig
	_dataItemID map[float64]ShopItemGroupConfig
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
	a._dataItemID = map[float64]ShopItemGroupConfig{}
	for _, d := range a._data {
		a._dataItemID[d.ItemID] = d
	}
}

// ByItemID returns the ShopItemGroupConfig uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *ShopItemGroupConfigAccessor) ByItemID(identifier float64) (ShopItemGroupConfig, error) {
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
