package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"net/http"
	"io"
	"encoding/json"
)

type MonopolyShopConfig struct {
	GoodsIDList []float64 `json:"GoodsIDList"`
	ShopID      float64   `json:"ShopID"`
}
type MonopolyShopConfigAccessor struct {
	_data       []MonopolyShopConfig
	_dataShopID map[float64]MonopolyShopConfig
}

// LoadData retrieves the data. Must be called before MonopolyShopConfig.GroupData
func (a *MonopolyShopConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/MonopolyShopConfig.json")
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
func (a *MonopolyShopConfigAccessor) Raw() ([]MonopolyShopConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []MonopolyShopConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with MonopolyShopConfigAccessor.LoadData to preload everything
func (a *MonopolyShopConfigAccessor) GroupData() {
	a._dataShopID = map[float64]MonopolyShopConfig{}
	for _, d := range a._data {
		a._dataShopID[d.ShopID] = d
	}
}

// ByShopID returns the MonopolyShopConfig uniquely identified by ShopID
//
// Error is only non-nil if the source errors out
func (a *MonopolyShopConfigAccessor) ByShopID(identifier float64) (MonopolyShopConfig, error) {
	if a._dataShopID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return MonopolyShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataShopID[identifier], nil
}
