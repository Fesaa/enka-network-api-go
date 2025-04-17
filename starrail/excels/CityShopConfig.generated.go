package// Generated by Yoitsu. DO NOT EDIT!
excels

import (
	"github.com/Fesaa/enka-network-api-go/hash"
	"net/http"
	"io"
	"encoding/json"
)

type CityShopConfig struct {
	HintOverNum       float64   `json:"HintOverNum"`
	ItemID            float64   `json:"ItemID"`
	MaxLevel          float64   `json:"MaxLevel"`
	Name              hash.Hash `json:"Name"`
	RewardListGroupID float64   `json:"RewardListGroupID"`
	ShopID            float64   `json:"ShopID"`
	WorldID           float64   `json:"WorldID"`
	WorldImgPath      string    `json:"WorldImgPath"`
}
type CityShopConfigAccessor struct {
	_data                  []CityShopConfig
	_dataItemID            map[float64]CityShopConfig
	_dataMaxLevel          map[float64]CityShopConfig
	_dataRewardListGroupID map[float64]CityShopConfig
	_dataShopID            map[float64]CityShopConfig
	_dataWorldID           map[float64]CityShopConfig
	_dataWorldImgPath      map[string]CityShopConfig
}

// LoadData retrieves the data. Must be called before CityShopConfig.GroupData
func (a *CityShopConfigAccessor) LoadData() error {
	res, err := http.Get("https://gitlab.com/Dimbreath/turnbasedgamedata/-/raw/main/ExcelOutput/CityShopConfig.json")
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
func (a *CityShopConfigAccessor) Raw() ([]CityShopConfig, error) {
	if a._data == nil {
		err := a.LoadData()
		if err != nil {
			return []CityShopConfig{}, err
		}
	}
	return a._data, nil
}

// GroupData groups the data by their unique ids.
// Can be called manually in conjunction with CityShopConfigAccessor.LoadData to preload everything
func (a *CityShopConfigAccessor) GroupData() {
	a._dataItemID = map[float64]CityShopConfig{}
	a._dataMaxLevel = map[float64]CityShopConfig{}
	a._dataRewardListGroupID = map[float64]CityShopConfig{}
	a._dataShopID = map[float64]CityShopConfig{}
	a._dataWorldID = map[float64]CityShopConfig{}
	a._dataWorldImgPath = map[string]CityShopConfig{}
	for _, d := range a._data {
		a._dataItemID[d.ItemID] = d
		a._dataMaxLevel[d.MaxLevel] = d
		a._dataRewardListGroupID[d.RewardListGroupID] = d
		a._dataShopID[d.ShopID] = d
		a._dataWorldID[d.WorldID] = d
		a._dataWorldImgPath[d.WorldImgPath] = d
	}
}

// ByItemID returns the CityShopConfig uniquely identified by ItemID
//
// Error is only non-nil if the source errors out
func (a *CityShopConfigAccessor) ByItemID(identifier float64) (CityShopConfig, error) {
	if a._dataItemID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CityShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataItemID[identifier], nil
}

// ByMaxLevel returns the CityShopConfig uniquely identified by MaxLevel
//
// Error is only non-nil if the source errors out
func (a *CityShopConfigAccessor) ByMaxLevel(identifier float64) (CityShopConfig, error) {
	if a._dataMaxLevel == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CityShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataMaxLevel[identifier], nil
}

// ByRewardListGroupID returns the CityShopConfig uniquely identified by RewardListGroupID
//
// Error is only non-nil if the source errors out
func (a *CityShopConfigAccessor) ByRewardListGroupID(identifier float64) (CityShopConfig, error) {
	if a._dataRewardListGroupID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CityShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataRewardListGroupID[identifier], nil
}

// ByShopID returns the CityShopConfig uniquely identified by ShopID
//
// Error is only non-nil if the source errors out
func (a *CityShopConfigAccessor) ByShopID(identifier float64) (CityShopConfig, error) {
	if a._dataShopID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CityShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataShopID[identifier], nil
}

// ByWorldID returns the CityShopConfig uniquely identified by WorldID
//
// Error is only non-nil if the source errors out
func (a *CityShopConfigAccessor) ByWorldID(identifier float64) (CityShopConfig, error) {
	if a._dataWorldID == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CityShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWorldID[identifier], nil
}

// ByWorldImgPath returns the CityShopConfig uniquely identified by WorldImgPath
//
// Error is only non-nil if the source errors out
func (a *CityShopConfigAccessor) ByWorldImgPath(identifier string) (CityShopConfig, error) {
	if a._dataWorldImgPath == nil {
		if a._data == nil {
			err := a.LoadData()
			if err != nil {
				return CityShopConfig{}, err
			}
		}
		a.GroupData()
	}
	return a._dataWorldImgPath[identifier], nil
}
